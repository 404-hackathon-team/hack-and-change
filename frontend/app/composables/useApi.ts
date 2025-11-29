import type { UseFetchOptions } from '#app'

export const useAPI = <T>(url: string, options: UseFetchOptions<T> = {}) => {
    const headers = useRequestHeaders(['cookie'])

    return useFetch(url, {
        ...options,
        baseURL: '/api',
        headers: {
            ...headers,
            ...options.headers,
        },

        onResponseError({ response }) {
            if (response.status === 401) {
                const user = useUser()
                user.value = null
                navigateTo('/login')
            }
        }
    })
}