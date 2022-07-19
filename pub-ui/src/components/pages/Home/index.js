import { events } from '../../../services/';
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
            console.log('getting events...');
            const result = await events();
            result.result &&
                dispatch({ type: ADD_EVENT, value: result.result[0] });
            console.log('result: ', result);
        };

        fn();
    }, []);

    return (
        <>
            <h1>Home</h1>
            <div>Events</div>
            {state.events.map((e) => {
                return <div> {e.event}</div>;
            })}
            <Link href="/login/">Signin</Link>
        </>
    );
};

Home.displayName = 'HomeComponent';

export default Home;
