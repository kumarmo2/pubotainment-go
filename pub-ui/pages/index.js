import { userPing } from '../src/services';
import { getRequestHeadersForBackendRequest } from '../src/utils/server';
import Link from 'next/link';

const Home = () => {
    return (
        <>
            <h1>Home</h1>
            <Link href="/login/">Signin</Link>
        </>
    );
};

export async function getServerSideProps(context) {
    const headers = getRequestHeadersForBackendRequest(context);
    try {
        await userPing(headers);
        return {
            props: {},
        };
    } catch (err) {
        console.log('server error in getServerSideProps: ', err);
        if (err && err.redirect) {
            return {
                redirect: err.redirect,
            };
        }
        throw err;
    }
}

export default Home;
