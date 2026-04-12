import { useState, type FormEvent } from 'react'
import { useCreateTodo } from '../hooks/useTodos'

export default function AddTodo() {
  const [title, setTitle] = useState('')
  const { mutate: createTodo, isPending } = useCreateTodo()

  function handleSubmit(e: FormEvent) {
    e.preventDefault()
    const trimmed = title.trim()
    if (!trimmed) return

    createTodo(trimmed, {
      onSuccess: () => setTitle(''),
    })
  }

  return (
    <form onSubmit={handleSubmit} className="flex gap-2 mb-6">
      <input
        type="text"
        value={title}
        onChange={e => setTitle(e.target.value)}
        placeholder="What needs to be done?"
        disabled={isPending}
        className="flex-1 px-4 py-2 rounded-lg border border-gray-200 
                   bg-white shadow-sm outline-none
                   focus:ring-2 focus:ring-blue-500 focus:border-transparent
                   disabled:opacity-50 transition"
      />
      <button
        type="submit"
        disabled={isPending || !title.trim()}
        className="px-5 py-2 bg-blue-500 text-white font-medium rounded-lg 
                   shadow-sm hover:bg-blue-600 active:bg-blue-700
                   disabled:opacity-50 disabled:cursor-not-allowed transition"
      >
        {isPending ? 'Adding...' : 'Add'}
      </button>
    </form>
  )
}