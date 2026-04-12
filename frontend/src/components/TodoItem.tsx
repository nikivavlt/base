import { useUpdateTodo, useDeleteTodo } from '../hooks/useTodos'
import type { Todo } from '../types/todo'

interface Props {
  todo: Todo
}

export default function TodoItem({ todo }: Props) {
  const { mutate: updateTodo, isPending: isUpdating } = useUpdateTodo()
  const { mutate: deleteTodo, isPending: isDeleting } = useDeleteTodo()

  const isPending = isUpdating || isDeleting

  return (
    <li className={`flex items-center gap-3 p-4 bg-white rounded-lg 
                    shadow-sm border border-gray-100 transition
                    ${isPending ? 'opacity-50' : ''}`}>

      {/* Checkbox */}
      <input
        type="checkbox"
        checked={todo.completed}
        disabled={isPending}
        onChange={() => updateTodo({ id: todo.id, completed: !todo.completed })}
        className="w-5 h-5 rounded accent-blue-500 cursor-pointer flex-shrink-0"
      />

      {/* Title */}
      <span className={`flex-1 text-gray-800 text-sm
                        ${todo.completed ? 'line-through text-gray-400' : ''}`}>
        {todo.title}
      </span>

      {/* Delete button */}
      <button
        onClick={() => deleteTodo(todo.id)}
        disabled={isPending}
        className="text-gray-300 hover:text-red-400 disabled:cursor-not-allowed
                   transition text-lg leading-none"
        aria-label="Delete todo"
      >
        ✕
      </button>
    </li>
  )
}