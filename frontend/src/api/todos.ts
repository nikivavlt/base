import axios, { AxiosError, type InternalAxiosRequestConfig } from 'axios'
import { authService } from './auth'
import type { Todo } from '../types/todo'

// Extended config to track retry attempts
interface RetryConfig extends InternalAxiosRequestConfig {
  _retry?: boolean
}

let accessToken: string | null = null

// Called by AuthContext to keep token in sync
export function setApiToken(token: string | null) {
  accessToken = token
}

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: { 'Content-Type': 'application/json' },
  withCredentials: true,
})

// Attach access token to every request
api.interceptors.request.use(config => {
  if (accessToken) {
    config.headers.Authorization = `Bearer ${accessToken}`
  }
  return config
})

// Auto-refresh on 401
api.interceptors.response.use(
  res => res,
  async (err: AxiosError) => {
    const config = err.config as RetryConfig

    if (err.response?.status === 401 && !config._retry) {
      config._retry = true // prevent infinite retry loop

      try {
        const { access_token } = await authService.refresh()
        accessToken = access_token
        config.headers!.Authorization = `Bearer ${access_token}`
        return api(config) // retry original request
      } catch {
        // Refresh failed — session truly expired
        accessToken = null
        window.dispatchEvent(new Event('auth:logout'))
        return Promise.reject(err)
      }
    }

    const message = (err.response?.data as any)?.error ?? err.message
    return Promise.reject(new Error(message))
  }
)

export const todosApi = {
  getAll:  async (): Promise<Todo[]>  => { const { data } = await api.get('/todos');          return data },
  getOne:  async (id: number)         => { const { data } = await api.get(`/todos/${id}`);    return data },
  create:  async (title: string)      => { const { data } = await api.post('/todos', {title}); return data },
  update:  async (id: number, payload: Partial<Pick<Todo, 'title'|'completed'>>) => {
    const { data } = await api.patch(`/todos/${id}`, payload)
    return data
  },
  delete: async (id: number): Promise<void> => { await api.delete(`/todos/${id}`) },
}