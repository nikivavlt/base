import { useState } from 'react'
import { useAuth } from '../context/AuthContext'
import LoginPage from '../pages/LoginPage'
import RegisterPage from '../pages/RegisterPage'

export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const { user } = useAuth()
  const [mode, setMode] = useState<'login' | 'register'>('login')

  if (!user) {
    return mode === 'login'
      ? <LoginPage    onSwitch={() => setMode('register')} />
      : <RegisterPage onSwitch={() => setMode('login')} />
  }

  return <>{children}</>
}