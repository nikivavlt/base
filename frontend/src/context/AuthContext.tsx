import {
    createContext,
    useContext,
    useState,
    useCallback,
    useEffect,
    type ReactNode,
  } from 'react'
  import { authService } from '../api/auth'
import type { AuthState, User } from '../types/auth'
import { setApiToken } from '../api/todos'

// Silent refresh on mount — when user reopens the app, 
// we try /auth/refresh. If the cookie is still valid (within 7 days), 
// session is restored transparently. No login needed.
  
  interface AuthContextValue extends AuthState {
    login:  (email: string, password: string) => Promise<void>
    logout: () => Promise<void>
    register: (email: string, password: string) => Promise<void>
    setAccessToken: (token: string) => void
  }
  
  const AuthContext = createContext<AuthContextValue | null>(null)
  
  export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser]               = useState<User | null>(null)
    const [accessToken, setAccessToken] = useState<string | null>(null)
    const [loading, setLoading]         = useState(true) // silent refresh on mount
  
    useEffect(() => {
        const handler = () => {
          setUser(null)
          setAccessToken(null)
        }
        window.addEventListener('auth:logout', handler)
        return () => window.removeEventListener('auth:logout', handler)
    }, [])

    // On app load — try to restore session via refresh cookie
    useEffect(() => {
      authService.refresh()
        .then(({ access_token }) => {
          setAccessToken(access_token)
          setApiToken(access_token) 
          // Decode user from token payload (no extra request needed)
          const payload = JSON.parse(atob(access_token.split('.')[1]))
          setUser({ id: payload.user_id, email: payload.email })
        })
        .catch(() => {
          // No valid session — user needs to log in
        })
        .finally(() => setLoading(false))
    }, [])
  
    // In login:
    const login = useCallback(async (email: string, password: string) => {
        const { access_token, user } = await authService.login(email, password)
        setAccessToken(access_token)
        setApiToken(access_token)    // ← sync to axios
        setUser(user)
    }, [])
    
    // In logout:
    const logout = useCallback(async () => {
        await authService.logout()
        setAccessToken(null)
        setApiToken(null)            // ← sync to axios
        setUser(null)
    }, [])
  
    const register = useCallback(async (email: string, password: string) => {
      await authService.register(email, password)
    }, [])
  
    if (loading) {
      return (
        <div className="min-h-screen bg-gray-50 flex items-center justify-center">
          <div className="w-6 h-6 border-2 border-blue-500 border-t-transparent rounded-full animate-spin" />
        </div>
      )
    }
  
    return (
      <AuthContext.Provider value={{ user, accessToken, login, logout, register, setAccessToken }}>
        {children}
      </AuthContext.Provider>
    )
  }
  
// Typed hook — throws if used outside provider
export function useAuth() {
    const ctx = useContext(AuthContext)
    if (!ctx) throw new Error('useAuth must be used within AuthProvider')
    return ctx
}

