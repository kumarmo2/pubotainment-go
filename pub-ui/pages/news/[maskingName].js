import { useRouter } from 'next/router'
import { getRandomProps } from '../../services'

const NewsDetails = () => {
    const router = useRouter()
    console.log(router.query)

    return <h1>News Details page</h1>
}

export async function getServerSideProps() {
    console.log('this code will run only on server side')
    const props = await getRandomProps()
    console.log('props: ', props)
    return { props }
}

export default NewsDetails
