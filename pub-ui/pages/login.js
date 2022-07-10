import Login from '../src/components/pages/Login';
import { userPing } from '../src/services';
import { getRequestHeadersForBackendRequest } from '../src/utils/server';

const LoginPage = () => {
    return (
        <>
            <Login />
        </>
    );
};

export async function getServerSideProps(context) {
    try {
        const headers = getRequestHeadersForBackendRequest(context);
        await userPing(headers);
        console.log('redirecting to home');
        return {
            redirect: {
                destination: '/',
                permanent: false,
            },
        };
    } catch (err) {
        console.log('Login: catch block');
        return { prop: {} };
    }
}
export default LoginPage;
