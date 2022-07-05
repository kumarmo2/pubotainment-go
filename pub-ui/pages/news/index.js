import Link from 'next/link'

const News = () => {
    return (
        <>
            <h1>News Page</h1>
            <ul>
                <li>
                    <Link href="/news/some-news-details">
                        Path to news details page
                    </Link>
                </li>
            </ul>
        </>
    )
}

export default News
