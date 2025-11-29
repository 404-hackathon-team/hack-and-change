export const useApiClient = () => {
    const api = $fetch.create({
        baseURL: '/api',
    })

    return api
}
