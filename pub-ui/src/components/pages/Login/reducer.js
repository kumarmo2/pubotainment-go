import { actionType } from './constants';

const reducer = (state, action) => {
    switch (action.type) {
        case actionType.UPDATE_COMPANY_ID: {
            return {
                ...state,
                companyId: action.value,
            };
        }
        case actionType.UPDATE_PASS: {
            return {
                ...state,
                pass: action.value,
            };
        }
        case actionType.SET_COMPANY_HELPER_TEXT: {
            return {
                ...state,
                companyHelperText: action.value,
            };
        }
        case actionType.SET_PASS_HELPER_TEXT: {
            return {
                ...state,
                passHelperText: action.value,
            };
        }
        default:
            return state;
    }
};
export default reducer;
