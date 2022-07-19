import { ADD_EVENT } from './actions';

const reducer = (state, action) => {
    switch (action.type) {
        case ADD_EVENT: {
            return {
                ...state,
                events: [...state.events, action.value],
            };
        }
        default:
            return state;
    }
};

export default reducer;
