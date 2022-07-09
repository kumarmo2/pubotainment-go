import '../styles/globals.css';
import theme from '../ui-common-lib/theme';
import Head from 'next/head';
import createEmotionCache from '../ui-common-lib/createEmotionCache';
import { CacheProvider } from '@emotion/react';

import { themeProvider as ThemeProvider } from '../ui-common-lib/theme';

const clientSideEmotionCache = createEmotionCache();

function MyApp(props) {
    const {
        Component,
        emotionCache = clientSideEmotionCache,
        pageProps,
    } = props;

    return (
        <CacheProvider value={emotionCache}>
            <Head>
                <meta
                    name="viewport"
                    content="initial-scale=1, width=device-width"
                />
            </Head>
            <ThemeProvider theme={theme}>
                <Component {...pageProps} />;
            </ThemeProvider>
        </CacheProvider>
    );
}

export default MyApp;
