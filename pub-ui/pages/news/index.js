import Link from 'next/link';
import { useEffect } from 'react';

const News = () => {
    useEffect(() => {
        console.log('News: componenetDidMount');
    }, []);
    return (
        <>
            <h1>News Page</h1>
            <ul>
                <li>
                    <Link href="/news/some-news-details">
                        <a>Path to news details page</a>
                    </Link>
                </li>
            </ul>
        </>
    );
};

export default News;
