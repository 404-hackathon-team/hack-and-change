
type User = {
    email: string
    lastName?: string,
    firstName?: string,
    id: unknown,
}

const AUTH_USER_KEY = 'authUser'
const AUTH_LOADING_ME_KEY = 'authLoadingMe'

export const useAuth = () => {
    const api = useApiClient()
    
    const user = useState<User | null>(AUTH_USER_KEY, () => null)
    const loadingMe = useState<boolean>(AUTH_LOADING_ME_KEY, () => false)

    const isAuthenticated = computed(() => !!user.value)

    const fetchMe = async () => {
        if (loadingMe.value) return

        loadingMe.value = true
        try {
            user.value = await api<User>('/me', {
                method: 'GET',
            })
        // @eslint-disable-next-line @typescript-eslint/no-explicit-any
        } catch (err: any) {
            const status = err?.status || err?.response?.status

            if (status === 401) {
                user.value = null
            } else {
                console.error('Ошибка при запросе /me', err)
            }
        } finally {
            loadingMe.value = false
        }
    }

    return {
        user,
        isAuthenticated,
        loadingMe,
        fetchMe,
    }
}
