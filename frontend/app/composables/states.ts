export interface User {
    id: number
    email: string
    name: string
}

export const useUser = () => useState<User | null>('user', () => null)