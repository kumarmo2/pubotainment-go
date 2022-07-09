import Link from 'next/link';
import Box from '@mui/material/Box';
import Input from '../../../ui-common-lib/components/Input';
import Button from '../../../ui-common-lib/components/Button';
import { useTheme } from '@mui/material';

const boxProps = {
    textAlign: 'center',
    border: '5px solid red',
    height: '100%',
    paddingLeft: 4,
    paddingRight: 4,
};

const Login = () => {
    const theme = useTheme();
    console.log('theme from LoginPage:', theme);

    return (
        <Box sx={boxProps}>
            <h1>Login Page</h1>
            <Link href="/news/sdzfvklds">
                <a>News Details Page</a>
            </Link>
            <Link href="/news/">
                <a>News Page</a>
            </Link>
            <Input label="CompanyId" fullWidth variant="standard" />
            <Input
                type="password"
                label="Enter Password"
                fullWidth
                variant="standard"
            />
            <Button>Login</Button>
        </Box>
    );
};

export default Login;
