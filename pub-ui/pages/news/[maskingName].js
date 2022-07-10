import { useRouter } from 'next/router';
import { getRandomProps } from '../../src/services';
import { useEffect } from 'react';

const NewsDetails = ({ newsDetails }) => {
    const router = useRouter();
    console.log(router.query);
    console.log('newsDetails:', newsDetails);
    useEffect(() => {
        console.log('NewsDetails: componentDidMount');
    }, []);

    // NOTE: componentDidMount to fetch the inital data is never required for pages
    // that uses getServerSideProps.
    // https://nextjs.org/docs/basic-features/data-fetching/get-server-side-props
    return <h1>News Details page</h1>;
};

export async function getServerSideProps() {
    try {
        console.log('NewsDetails: this code will run only on server side');
        const data = await getRandomProps();
        console.log('props: ', data);
        return {
            props: {
                newsDetails: data,
            },
        };
    } catch (err) {
        console.log('err:', err);
        return {
            redirect: {
                destination: '/login',
                permanent: false,
            },
        };
    }
}

export default NewsDetails;
