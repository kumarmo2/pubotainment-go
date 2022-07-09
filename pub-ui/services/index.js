const api = {};
// TODO: make this environment variable
const baseUrl = 'http://localhost';

api.get = async (path, ...params) => {
    return fetch(baseUrl + path, params).then(async (res) => {
        if (res.ok) {
            return await res.json();
        }
        if (res.status == 401) {
            throw {
                redirect: {
                    destination: '/login',
                    permanent: false,
                },
            };
        }
    });
};

export const getRandomProps = () => {
    return api.get('/api/props');
};

export default api;
