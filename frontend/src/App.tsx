import { useAuth } from './context/AuthContext'
import AddTodo from './components/AddTodo'
import TodoList from './components/TodoList'
import ProtectedRoute from './pages/ProtectedRoute'

export default function App() {
  const { user, logout } = useAuth()

  return (
    <ProtectedRoute>
      <main className="min-h-screen bg-gray-50 py-12 px-4">
        <div className="max-w-md mx-auto">

          {/* Header */}
          <div className="mb-8 flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">My Todos</h1>
              <p className="text-gray-400 text-sm mt-1">{user?.email}</p>
            </div>
            <button
              onClick={logout}
              className="text-sm text-gray-400 hover:text-red-400 transition"
            >
              Sign out
            </button>
          </div>

          <AddTodo />
          <TodoList />

        </div>
      </main>
    </ProtectedRoute>
  )
}