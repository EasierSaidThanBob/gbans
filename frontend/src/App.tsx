import React, {
    useCallback,
    useMemo,
    useState,
    JSX,
    lazy,
    Suspense
} from 'react';
import {
    BrowserRouter as Router,
    Outlet,
    Route,
    Routes
} from 'react-router-dom';
import NiceModal from '@ebay/nice-modal-react';
import { PaletteMode } from '@mui/material';
import { AlertColor } from '@mui/material/Alert';
import Container from '@mui/material/Container';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import * as Sentry from '@sentry/react';
import { PermissionLevel, UserProfile } from './api';
import { ErrorBoundary } from './component/ErrorBoundary';
import { Flash, Flashes } from './component/Flashes';
import { Footer } from './component/Footer';
import { LogoutHandler } from './component/LogoutHandler';
import { PrivateRoute } from './component/PrivateRoute';
import { TopBar } from './component/TopBar';
import { UserInit } from './component/UserInit';
import { ColourModeContext } from './contexts/ColourModeContext';
import { CurrentUserCtx, GuestProfile } from './contexts/CurrentUserCtx';
import { NotificationsProvider } from './contexts/NotificationsCtx';
import { UserFlashCtx } from './contexts/UserFlashCtx';
import { createThemeByMode } from './theme';

export interface AppProps {
    initialTheme: PaletteMode;
}

const SentryRoutes = Sentry.withSentryReactRouterV6Routing(Routes);
const HomePage = lazy(() => import('./page/HomePage'));
const AdminAppealsPage = lazy(() => import('./page/AdminAppealsPage'));
const AdminBanASNPage = lazy(() => import('./page/AdminBanASNPage'));
const AdminBanCIDRPage = lazy(() => import('./page/AdminBanCIDRPage'));
const AdminBanGroupPage = lazy(() => import('./page/AdminBanGroupPage'));
const AdminBanSteamPage = lazy(() => import('./page/AdminBanSteamPage'));
const AdminContestsPage = lazy(() => import('./page/AdminContestsPage'));
const AdminFiltersPage = lazy(() => import('./page/AdminFiltersPage'));
const AdminNetworkPage = lazy(() => import('./page/AdminNetworkPage'));
const AdminNewsPage = lazy(() => import('./page/AdminNewsPage'));
const AdminPeoplePage = lazy(() => import('./page/AdminPeoplePage'));
const AdminReportsPage = lazy(() => import('./page/AdminReportsPage'));
const AdminServersPage = lazy(() => import('./page/AdminServersPage'));
const BanPage = lazy(() => import('./page/BanPage'));
const ChatLogPage = lazy(() => import('./page/ChatLogPage'));
const ContestListPage = lazy(() => import('./page/ContestListPage'));
const ContestPage = lazy(() => import('./page/ContestPage'));
const ForumOverviewPage = lazy(() => import('./page/ForumOverviewPage'));
const ForumPage = lazy(() => import('./page/ForumPage'));
const ForumThreadPage = lazy(() => import('./page/ForumThreadPage'));
const LoginDiscordSuccessPage = lazy(
    () => import('./page/LoginDiscordSuccessPage')
);
const LoginPage = lazy(() => import('./page/LoginPage'));
const LoginSteamSuccessPage = lazy(
    () => import('./page/LoginSteamSuccessPage')
);
const LogoutPage = lazy(() => import('./page/LogoutPage'));
const MatchListPage = lazy(() => import('./page/MatchListPage'));
const MatchPage = lazy(() => import('./page/MatchPage'));
const NotificationsPage = lazy(() => import('./page/NotificationsPage'));
const PageNotFoundPage = lazy(() => import('./page/PageNotFoundPage'));
const PlayerStatsPage = lazy(() => import('./page/PlayerStatsPage'));
const PrivacyPolicyPage = lazy(() => import('./page/PrivacyPolicyPage'));
const ProfilePage = lazy(() => import('./page/ProfilePage'));
const ProfileSettingsPage = lazy(() => import('./page/ProfileSettingsPage'));
const ReportCreatePage = lazy(() => import('./page/ReportCreatePage'));
const ReportViewPage = lazy(() => import('./page/ReportViewPage'));
const STVPage = lazy(() => import('./page/STVPage'));
const ServersPage = lazy(() => import('./page/ServersPage'));
const StatsPage = lazy(() => import('./page/StatsPage'));
const StatsWeaponOverallPage = lazy(
    () => import('./page/StatsWeaponOverallPage')
);
const WikiPage = lazy(() => import('./page/WikiPage'));

export const App = ({ initialTheme }: AppProps): JSX.Element => {
    const [currentUser, setCurrentUser] =
        useState<NonNullable<UserProfile>>(GuestProfile);
    const [flashes, setFlashes] = useState<Flash[]>([]);

    const saveUser = (profile: UserProfile) => {
        setCurrentUser(profile);
    };

    const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
    const [mode, setMode] = useState<'light' | 'dark'>(
        initialTheme ? initialTheme : prefersDarkMode ? 'dark' : 'light'
    );

    const updateMode = (prevMode: PaletteMode): PaletteMode => {
        const m = prevMode === 'light' ? 'dark' : ('light' as PaletteMode);
        localStorage.setItem('theme', m);
        return m;
    };

    const colorMode = useMemo(
        () => ({
            toggleColorMode: () => {
                setMode(updateMode);
            }
        }),
        []
    );

    const theme = useMemo(() => createThemeByMode(mode), [mode]);

    const sendFlash = useCallback(
        (
            level: AlertColor,
            message: string,
            heading = 'header',
            closable = true
        ) => {
            if (
                flashes.length &&
                flashes[flashes.length - 1]?.message == message
            ) {
                // Skip duplicates
                return;
            }
            setFlashes([
                ...flashes,
                {
                    closable: closable ?? false,
                    heading: heading,
                    level: level,
                    message: message
                }
            ]);
        },
        [flashes, setFlashes]
    );

    return (
        <CurrentUserCtx.Provider
            value={{
                currentUser,
                setCurrentUser: saveUser
            }}
        >
            <UserFlashCtx.Provider value={{ flashes, setFlashes, sendFlash }}>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                    <Router>
                        <React.Fragment>
                            <ColourModeContext.Provider value={colorMode}>
                                <ThemeProvider theme={theme}>
                                    <NotificationsProvider>
                                        <React.StrictMode>
                                            <NiceModal.Provider>
                                                <UserInit />
                                                <LogoutHandler />
                                                <CssBaseline />
                                                <Container maxWidth={'lg'}>
                                                    <TopBar />
                                                    <div
                                                        style={{
                                                            marginTop: 24
                                                        }}
                                                    >
                                                        <Sentry.ErrorBoundary>
                                                            <SentryRoutes>
                                                                <Route
                                                                    path={'/'}
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <HomePage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/servers'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ServersPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/stv'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <STVPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/login/success'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <LoginSteamSuccessPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/privacy-policy'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivacyPolicyPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/contests'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ContestListPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={'/'}
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <HomePage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/servers'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ServersPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/stv'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <STVPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/login/success'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <LoginSteamSuccessPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/contests'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ContestListPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/contests/:contest_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ContestPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/stats'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <StatsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/stats/player/:steam_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <PlayerStatsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/stats/weapon/:weapon_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <StatsWeaponOverallPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/forums'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ForumOverviewPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/forums/:forum_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ForumPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/forums/thread/:forum_thread_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ForumThreadPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/wiki'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <WikiPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/wiki/:slug'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <WikiPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/ban/:ban_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <BanPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/report/:report_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ReportViewPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/log/:match_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <MatchPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/logs/:steam_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <MatchListPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/report'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <ReportCreatePage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/notifications'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <NotificationsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/settings'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <ProfileSettingsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/profile/:steam_id'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <ProfilePage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/report'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <Route
                                                                                    path={
                                                                                        '/ban/:ban_id'
                                                                                    }
                                                                                    element={
                                                                                        <BanPage />
                                                                                    }
                                                                                />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/ban/steam'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminBanSteamPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/ban/cidr'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminBanCIDRPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/ban/asn'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminBanASNPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/ban/group'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminBanGroupPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/filters'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Editor
                                                                                }
                                                                            >
                                                                                <AdminFiltersPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/network'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Editor
                                                                                }
                                                                            >
                                                                                <AdminNetworkPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/reports'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminReportsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/contests'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminContestsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/appeals'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminAppealsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/admin/news'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Editor
                                                                                }
                                                                            >
                                                                                <AdminNewsPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/chatlogs'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <ChatLogPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/admin/people'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Moderator
                                                                                }
                                                                            >
                                                                                <AdminPeoplePage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />

                                                                <Route
                                                                    path={
                                                                        '/admin/servers'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.Admin
                                                                                }
                                                                            >
                                                                                <AdminServersPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/login'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <LoginPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/login/discord'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PrivateRoute
                                                                                permission={
                                                                                    PermissionLevel.User
                                                                                }
                                                                            >
                                                                                <LoginDiscordSuccessPage />
                                                                            </PrivateRoute>
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path={
                                                                        '/logout'
                                                                    }
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <LogoutPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                                <Route
                                                                    path="/404"
                                                                    element={
                                                                        <ErrorBoundary>
                                                                            <PageNotFoundPage />
                                                                        </ErrorBoundary>
                                                                    }
                                                                />
                                                            </SentryRoutes>
                                                        </Sentry.ErrorBoundary>
                                                        <Suspense
                                                            fallback={
                                                                <h2>Loading</h2>
                                                            }
                                                        >
                                                            <Outlet />
                                                        </Suspense>
                                                    </div>
                                                    <Footer />
                                                </Container>
                                                <Flashes />
                                            </NiceModal.Provider>
                                        </React.StrictMode>
                                    </NotificationsProvider>
                                </ThemeProvider>
                            </ColourModeContext.Provider>
                        </React.Fragment>
                    </Router>
                </LocalizationProvider>
            </UserFlashCtx.Provider>
        </CurrentUserCtx.Provider>
    );
};
