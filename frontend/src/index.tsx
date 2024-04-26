import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import '@fontsource/roboto/latin-300.css';
import '@fontsource/roboto/latin-400.css';
import '@fontsource/roboto/latin-500.css';
import '@fontsource/roboto/latin-700.css';
import * as Sentry from '@sentry/react';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import './fonts/tf2build.css';
import { AuthProvider, useCurrentUserCtx } from './hooks/useCurrentUserCtx.tsx';
import { routeTree } from './routeTree.gen';

// Create a new router instance
const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
    context: {
        auth: undefined!
    }
});

// Register the router instance for type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router;
    }
}

// extend window with our own items that we inject
declare global {
    interface Window {
        gbans: {
            site_name: string;
            discord_client_id: string;
            discord_link_id: string;
            asset_url: string;
            bucket_media: string;
            bucket_demo: string;
            build_version: string;
            build_date: string;
            build_commit: string;
            sentry_dsn_web: string;
        };
    }
}

window.gbans = window.gbans || [];

if (window.gbans.sentry_dsn_web != '') {
    // TODO instrumentation for tanstack router, not currently officially supported
    Sentry.init({
        dsn: window.gbans.sentry_dsn_web,
        release: window.gbans.build_version,
        // Performance Monitoring
        tracesSampleRate: 1.0, //  Capture 100% of the transactions
        // Session Replay
        replaysSessionSampleRate: 0.1, // This sets the sample rate at 10%. You may want to change it to 100% while in development and then sample at a lower rate in production.
        replaysOnErrorSampleRate: 1.0 // If you're not already sampling the entire session, change the sample rate to 100% when sampling sessions where errors occur.
    });
}

function InnerApp() {
    const auth = useCurrentUserCtx();
    return <RouterProvider router={router} context={{ auth }} />;
}

function App() {
    return (
        <AuthProvider>
            <StrictMode>
                <InnerApp />
            </StrictMode>
        </AuthProvider>
    );
}

const container = document.getElementById('root');
if (container) {
    createRoot(container).render(<App />);
}
