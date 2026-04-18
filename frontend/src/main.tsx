import { createRoot } from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import App from './App.tsx'
import './index.css'
import ErrorBoundary from './components/ErrorBoundary.tsx'
import { AuthProvider } from './context/AuthContext.tsx'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 30,   // data stays fresh for 30 seconds
      retry: 1,               // retry failed requests once
    },
  },
})

createRoot(document.getElementById('root')!).render(
  <ErrorBoundary>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <App />
      </AuthProvider>
    </QueryClientProvider>
  </ErrorBoundary>
)