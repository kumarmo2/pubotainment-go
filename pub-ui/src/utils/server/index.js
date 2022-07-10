export const getRequestHeadersForBackendRequest = (context) => {
    const headers = context.req.headers;
    return {
        headers,
    };
};
