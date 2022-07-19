import { userPing } from '../src/services';
import { getRequestHeadersForBackendRequest } from '../src/utils/server';
import HomeComponent from '../src/components/pages/Home';

const Home = () => {
    return <HomeComponent />;
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

Home.displayName = 'Home';

export default Home;
