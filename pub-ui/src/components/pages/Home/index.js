import { useEffect, useReducer } from 'react';
import reducer from './reducer';
import { ADD_EVENT } from './actions';
import Link from 'next/link';

const initialState = {
    events: [],
};

const Home = () => {
    const [state, dispatch] = useReducer(reducer, initialState);

    useEffect(() => {
        const fn = async () => {
            const ws = new WebSocket('ws://localhost/events/');

            setInterval(() => {
                ws.send('sdf klsdflsdfksdf');
            }, 5000);
        };

        fn();
    }, []);

    return (
        <>
            <h1>Home</h1>
            <div>Events</div>
            <Link href="/login/">Signin</Link>
        </>
    );
};

Home.displayName = 'HomeComponent';

export default Home;
