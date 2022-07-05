const api = {}
// TODO: make this environment variable
const baseUrl = 'http://localhost:8000'

api.get = async (path, ...params) => {
    return fetch(baseUrl + path, params).then(async (res) => {
        if (res.ok) {
            return await res.json()
        }
    })
}

export const getRandomProps = () => {
    return api.get('/api/props')
}

export default api
