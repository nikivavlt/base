import axios from 'axios'
import type { User } from '../types/auth'

const authApi = axios.create({
  baseURL: import.meta.env.VITE_AUTH_URL,
  headers: { 'Content-Type': 'application/json' },
  withCredentials: true, // send httpOnly cookie on refresh/logout
})

export const authService = {
  register: async (email: string, password: string): Promise<User> => {
    const { data } = await authApi.post('/auth/register', { email, password })
    return data
  },

  login: async (email: string, password: string): Promise<{ access_token: string; user: User }> => {
    const { data } = await authApi.post('/auth/login', { email, password })
    return data
  },

  refresh: async (): Promise<{ access_token: string }> => {
    const { data } = await authApi.post('/auth/refresh')
    return data
  },

  logout: async (): Promise<void> => {
    await authApi.post('/auth/logout')
  },
}