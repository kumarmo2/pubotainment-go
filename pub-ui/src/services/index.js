const api = {};
// TODO: make this environment variable
const baseUrl = 'http://localhost';

const commonApiResponseHandler = async (res) => {
    if (res.ok) {
        return await res.json();
    }
    if (res.status == 401) {
        throw {
            status: res.status,
            redirect: {
                destination: '/login',
                permanent: false,
            },
        };
    }
};
api.get = async (path, params) => {
    console.log('params from get: ', params);
    return fetch(baseUrl + path, params)
        .then(commonApiResponseHandler)
        .catch((err) => {
            throw err;
        });
};

api.post = async (path, { body, ...rest }) => {
    const requestBody = body ? JSON.stringify(body) : null;
    console.log('rest: ', rest);
    return fetch(baseUrl + path, {
        method: 'POST',
        body: requestBody,
        ...rest,
    })
        .then(commonApiResponseHandler)
        .catch((err) => {
            throw err;
        });
};

export const getRandomProps = () => {
    return api.get('/api/props');
};

export const userPing = (rest) => {
    return api.get('/api/user/ping', rest);
};

export const userSignIn = ({ companyId, pass, ...rest }) => {
    const body = { companyId, password: pass };
    return api.post('/api/user/signin', { body, rest });
};

export const events = (rest) => {
    return api.get('/api/user/events/', rest);
};

export default api;
