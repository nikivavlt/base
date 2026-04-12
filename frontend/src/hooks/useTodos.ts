import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { todosApi } from '../api/todo'
import type { Todo } from '../types/todo'

const QUERY_KEY = ['todos']

export function useTodos() {
  return useQuery({
    queryKey: QUERY_KEY,
    queryFn: todosApi.getAll,
  })
}

export function useCreateTodo() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (title: string) => todosApi.create(title),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: QUERY_KEY }),
  })
}

export function useUpdateTodo() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, ...payload }: { id: number; title?: string; completed?: boolean }) =>
      todosApi.update(id, payload),

    // Fire before the request — update UI instantly
    onMutate: async ({ id, completed }) => {
      await queryClient.cancelQueries({ queryKey: QUERY_KEY })
      const previous = queryClient.getQueryData<Todo[]>(QUERY_KEY)

      if (completed !== undefined) {
        queryClient.setQueryData<Todo[]>(QUERY_KEY, old =>
          old?.map(t => t.id === id ? { ...t, completed } : t) ?? []
        )
      }

      return { previous } // snapshot for rollback
    },

    // Rollback on failure
    onError: (_err, _vars, context) => {
      if (context?.previous) {
        queryClient.setQueryData(QUERY_KEY, context.previous)
      }
    },

    onSettled: () => queryClient.invalidateQueries({ queryKey: QUERY_KEY }),
  })
}

export function useDeleteTodo() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: number) => todosApi.delete(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: QUERY_KEY }),
  })
}