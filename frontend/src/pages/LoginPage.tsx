import { useState, type FormEvent } from 'react'
import { useAuth } from '../context/AuthContext'

interface Props {
  onSwitch: () => void // switch to register
}

export default function LoginPage({ onSwitch }: Props) {
  const { login } = useAuth()
  const [email, setEmail]       = useState('')
  const [password, setPassword] = useState('')
  const [error, setError]       = useState('')
  const [loading, setLoading]   = useState(false)

  async function handleSubmit(e: FormEvent) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      await login(email, password)
    } catch (err: any) {
      setError(err.message ?? 'Login failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center px-4">
      <div className="w-full max-w-sm bg-white rounded-2xl shadow-sm border border-gray-100 p-8">
        <h1 className="text-2xl font-bold text-gray-900 mb-1">Welcome back</h1>
        <p className="text-sm text-gray-400 mb-6">Sign in to your account</p>

        {error && (
          <div className="mb-4 px-4 py-3 rounded-lg bg-red-50 border border-red-100
                          text-sm text-red-600">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Email
            </label>
            <input
              type="email"
              value={email}
              onChange={e => setEmail(e.target.value)}
              required
              disabled={loading}
              placeholder="you@example.com"
              className="w-full px-4 py-2 rounded-lg border border-gray-200
                         outline-none focus:ring-2 focus:ring-blue-500
                         focus:border-transparent text-sm disabled:opacity-50 transition"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Password
            </label>
            <input
              type="password"
              value={password}
              onChange={e => setPassword(e.target.value)}
              required
              disabled={loading}
              placeholder="••••••••"
              className="w-full px-4 py-2 rounded-lg border border-gray-200
                         outline-none focus:ring-2 focus:ring-blue-500
                         focus:border-transparent text-sm disabled:opacity-50 transition"
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-2 bg-blue-500 text-white font-medium rounded-lg
                       hover:bg-blue-600 active:bg-blue-700 transition
                       disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? 'Signing in...' : 'Sign in'}
          </button>
        </form>

        <p className="mt-6 text-center text-sm text-gray-400">
          No account?{' '}
          <button
            onClick={onSwitch}
            className="text-blue-500 hover:underline font-medium"
          >
            Register
          </button>
        </p>
      </div>
    </div>
  )
}