// import Link from 'next/link';
import Box from '@mui/material/Box';
import Input from '../../../../ui-common-lib/components/Input';
import Button from '../../../../ui-common-lib/components/Button';
// import { useTheme } from '@mui/material';
import { useRouter } from 'next/router';
import { useReducer, useCallback } from 'react';
import reducer from './reducer';
import { actionType } from './constants';
import { userSignIn } from '../../../services';

const boxProps = {
    textAlign: 'center',
    // border: '5px solid red',
    height: '100%',
    paddingLeft: 4,
    paddingRight: 4,
    paddingTop: 20,
};

const initialState = {
    companyId: '',
    pass: '',
    companyHelperText: null,
    passHelperText: null,
};

const Login = () => {
    // const theme = useTheme();
    const [state, dispatch] = useReducer(reducer, initialState);
    const router = useRouter();

    const handleCompanyIdChange = useCallback((event) => {
        if (!event || !event.target) {
            return;
        }
        const value = event.target.value;
        dispatch({ type: actionType.UPDATE_COMPANY_ID, value: value || '' });
        console.log('comId value: ', value);
        if (!value) {
            dispatch({
                type: actionType.SET_COMPANY_HELPER_TEXT,
                value: 'Required',
            });
            return;
        }
        const numberValue = +value;
        if (!numberValue) {
            dispatch({
                type: actionType.SET_COMPANY_HELPER_TEXT,
                value: 'Enter valid companyId(Numeric)',
            });
            return;
        }
        dispatch({ type: actionType.UPDATE_COMPANY_ID, value: numberValue });
        dispatch({ type: actionType.SET_COMPANY_HELPER_TEXT, value: null });
    }, []);

    const handlePassChange = useCallback((event) => {
        if (!event || !event.target) {
            return;
        }
        const value = event.target.value;
        dispatch({ type: actionType.UPDATE_PASS, value });
        if (!value) {
            dispatch({
                type: actionType.SET_PASS_HELPER_TEXT,
                value: 'Required',
            });
            return;
        }
        dispatch({ type: actionType.SET_PASS_HELPER_TEXT, value: null });

        console.log('event: ', event.target.value);
    }, []);

    const isLoginButtonDisabled = useCallback(() => {
        return !state.pass || isNaN(state.companyId) || +state.companyId <= 0;
    }, [state.companyId, state.pass]);

    const handleLoginButtonClick = useCallback(async () => {
        console.log('login clicked');
        try {
            await userSignIn({
                companyId: state.companyId,
                pass: state.pass,
            });
            router.replace('/');
        } catch (err) {
            if (err && err.status === 401) {
                dispatch({
                    actionType: actionType.SET_PASS_HELPER_TEXT,
                    value: 'Invalid Credentials',
                });
                return;
            }
        }
    }, [state.companyId, state.pass]);

    return (
        <Box sx={boxProps}>
            <h1>Login</h1>
            <Input
                required
                value={state.companyId}
                onChange={handleCompanyIdChange}
                error={!!state.companyHelperText}
                helperText={state.companyHelperText}
                label="CompanyId"
                fullWidth
                variant="standard"
            />
            <Input
                required
                onChange={handlePassChange}
                helperText={state.passHelperText}
                error={!!state.passHelperText}
                type="password"
                value={state.pass}
                label="Enter Password"
                fullWidth
            />
            <Button
                onClick={handleLoginButtonClick}
                disabled={isLoginButtonDisabled()}
                fullWidth
            >
                Login
            </Button>
        </Box>
    );
};

export default Login;
