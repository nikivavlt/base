import AddTodo from './components/AddTodo'
import TodoList from './components/TodoList'

export default function App() {
  return (
    <main className="min-h-screen bg-gray-50 py-12 px-4">
      <div className="max-w-md mx-auto">

        {/* Header */}
        <div className="mb-8 text-center">
          <h1 className="text-3xl font-bold text-gray-900">My Todos</h1>
          <p className="text-gray-400 text-sm mt-1">Stay on top of things</p>
        </div>

        {/* Add form */}
        <AddTodo />

        {/* List */}
        <TodoList />

      </div>
    </main>
  )
}