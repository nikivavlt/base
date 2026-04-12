import axios from 'axios'
import type { Todo } from '../types/todo'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: { 'Content-Type': 'application/json' },
})

api.interceptors.response.use(
  res => res,
  err => {
    const message = err.response?.data?.error ?? err.message ?? 'Unknown error'
    return Promise.reject(new Error(message))
  }
)

export const todosApi = {
  getAll: async (): Promise<Todo[]> => {
    const { data } = await api.get('/todos')
    return data
  },

  getOne: async (id: number): Promise<Todo> => {
    const { data } = await api.get(`/todos/${id}`)
    return data
  },

  create: async (title: string): Promise<Todo> => {
    const { data } = await api.post('/todos', { title })
    return data
  },

  update: async (id: number, payload: Partial<Pick<Todo, 'title' | 'completed'>>): Promise<Todo> => {
    const { data } = await api.patch(`/todos/${id}`, payload)
    return data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/todos/${id}`)
  },
}