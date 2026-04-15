export interface User {
    id: number
    email: string
}
  
  export interface AuthState {
    user: User | null
    accessToken: string | null
}