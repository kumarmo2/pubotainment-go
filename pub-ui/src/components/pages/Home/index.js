import { useEffect, useReducer } from 'react';
import reducer from './reducer';
import { ADD_EVENT } from './actions';
import Link from 'next/link';


const Home = () => {
    const [state, dispatch] = useReducer(reducer, { events: []});

    useEffect(() => {
            const messageHandler = event => {
                console.log('Message from server ', event.data);
                dispatch({ type: ADD_EVENT, value: JSON.parse(event.data)})
            }
            const ws = new WebSocket('ws://localhost/events/');
            ws.addEventListener('message', messageHandler);

        return () => {
            console.log("removing listener");
            ws.close()
            ws.removeEventListener('message', messageHandler, true);
        }

    }, []);
    console.log("events.length: ", state.events.length);

    return (
        <>
            <h1>Home</h1>
            <div>Events-1</div>
            {
                state.events.map((event, index) => {
                        console.log("rendering: ", index, event) 
                    return <div key={index}> {event.event.name} </div>;
            })
            }
            <Link href="/login/">Signin</Link>
        </>
    );
};

Home.displayName = 'HomeComponent';

export default Home;
