import { useTodos } from '../hooks/useTodos'
import TodoItem from './TodoItem'
import TodoSkeleton from './TodoSkeleton'

export default function TodoList() {
  const { data: todos, isLoading, isError } = useTodos()

  if (isLoading) return <TodoSkeleton />

  if (isError) {
    return (
      <div className="text-center py-10 text-red-400">
        Failed to load todos. Is the backend running?
      </div>
    )
  }

  if (!todos?.length) {
    return (
      <div className="text-center py-10 text-gray-400">
        No todos yet. Add one above ↑
      </div>
    )
  }

  const remaining = todos.filter(t => !t.completed).length

  return (
    <div>
      {/* Stats bar */}
      <p className="text-xs text-gray-400 mb-3">
        {remaining} of {todos.length} remaining
      </p>

      {/* List */}
      <ul className="flex flex-col gap-2">
        {todos.map(todo => (
          <TodoItem key={todo.id} todo={todo} />
        ))}
      </ul>
    </div>
  )
}