import { ADD_EVENT } from './actions';

const reducer = (state, action) => {
    switch (action.type) {
        case ADD_EVENT: {
            return {
                ...state,
                events: [action.value, ...state.events],
            };
        }
        default:
            return state;
    }
};

export default reducer;
