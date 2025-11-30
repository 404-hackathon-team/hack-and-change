export const useApiClient = () => {
    const api = $fetch.create({
        baseURL: '/remote/api',
    })

    return api
}
