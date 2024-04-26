import { useMemo } from 'react';
import DoDisturbIcon from '@mui/icons-material/DoDisturb';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import { createLazyFileRoute } from '@tanstack/react-router';
import SteamID from 'steamid';
import { generateOIDCLink } from '../api';
import { ContainerWithHeader } from '../component/ContainerWithHeader';
import { useCurrentUserCtx } from '../hooks/useCurrentUserCtx.tsx';
import steamLogo from '../icons/steam_login_lg.png';

export const Route = createLazyFileRoute('/login')({
    component: LoginPage
});

export interface LoginFormProps {
    message?: string;
    title?: string;
}

export function LoginPage() {
    const { currentUser } = useCurrentUserCtx();
    const message =
        'To access this page, please login using your steam account below.';
    const title = 'Permission Denied';
    const loggedInUser = useMemo(() => {
        const sid = new SteamID(currentUser.steam_id);
        return sid.isValidIndividual();
    }, [currentUser.steam_id]);

    return (
        <Grid container justifyContent={'center'} alignItems={'center'}>
            <Grid xs={12}>
                <ContainerWithHeader title={title} iconLeft={<DoDisturbIcon />}>
                    <>
                        {loggedInUser && (
                            <Typography variant={'body1'} padding={2}>
                                Insufficient permission to access this page.
                            </Typography>
                        )}
                        {!loggedInUser && (
                            <>
                                <Typography
                                    variant={'body1'}
                                    padding={2}
                                    paddingBottom={0}
                                >
                                    {message}
                                </Typography>
                                <Stack
                                    justifyContent="center"
                                    gap={2}
                                    flexDirection="row"
                                    width={1.0}
                                    flexWrap="wrap"
                                    padding={2}
                                >
                                    <Button
                                        sx={{ alignSelf: 'center' }}
                                        component={Link}
                                        href={generateOIDCLink(
                                            window.location.pathname
                                        )}
                                    >
                                        <img
                                            src={steamLogo}
                                            alt={'Steam Login'}
                                        />
                                    </Button>
                                </Stack>
                            </>
                        )}
                    </>
                </ContainerWithHeader>
            </Grid>
        </Grid>
    );
}
