import CellTowerIcon from '@mui/icons-material/CellTower';
import Grid from '@mui/material/Unstable_Grid2';
import { createFileRoute } from '@tanstack/react-router';
import { ContainerWithHeader } from '../component/ContainerWithHeader.tsx';

export const Route = createFileRoute('/_mod/admin/network/ipInfo')({
    component: AdminNetworkInfo
});

// const InfoRow = ({ label, children }: { label: string; children: ReactNode }) => {
//     return (
//         <TableRow hover>
//             <TableCell>
//                 <Typography fontWeight={700}> {label}</Typography>
//             </TableCell>
//             <TableCell>{children}</TableCell>
//         </TableRow>
//     );
// };

function AdminNetworkInfo() {
    // const [state, setState] = useUrlState({
    //     ip: undefined
    // });
    //
    // const { data, loading } = useNetworkQuery({ ip: state.ip });
    //
    // const onSubmit = (values: IPFieldProps) => {
    //     setState((prevState) => {
    //         return { ...prevState, ip: values.ip };
    //     });
    // };
    //
    // const pos = useMemo(() => {
    //     if (!data || data?.location.lat_long.latitude == 0) {
    //         return { lat: 50, lng: 50 };
    //     }
    //     return {
    //         lat: data?.location.lat_long.latitude,
    //         lng: data?.location.lat_long.longitude
    //     };
    // }, [data]);

    return (
        <ContainerWithHeader title="Network Info" iconLeft={<CellTowerIcon />}>
            <>
                <Grid container spacing={2}>
                    <Grid xs={12}>
                        {/*<Formik onSubmit={onSubmit} initialValues={{ ip: '' }}>*/}
                        <Grid container direction="row" alignItems="top" justifyContent="center" spacing={2}>
                            {/*<Grid xs>*/}
                            {/*    <IPField />*/}
                            {/*</Grid>*/}
                            {/*<Grid xs={2}>*/}
                            {/*    <SubmitButton*/}
                            {/*        label={'Submit'}*/}
                            {/*        fullWidth*/}
                            {/*        disabled={loading}*/}
                            {/*        startIcon={<SearchIcon />}*/}
                            {/*    />*/}
                            {/*</Grid>*/}
                        </Grid>
                        {/*</Formik>*/}
                    </Grid>
                </Grid>
                <Grid container spacing={2}>
                    <Grid xs={12}>
                        {/*{loading ? (*/}
                        {/*    <LoadingPlaceholder />*/}
                        {/*) : (*/}
                        {/*    <div>*/}
                        {/*        <Grid container spacing={2}>*/}
                        {/*            <Grid xs={12} md={6}>*/}
                        {/*                <Typography variant={'h4'} padding={2}>*/}
                        {/*                    Location*/}
                        {/*                </Typography>*/}
                        {/*                <TableContainer>*/}
                        {/*                    <Table>*/}
                        {/*                        <TableBody>*/}
                        {/*                            <InfoRow label={'Country'}>*/}
                        {/*                                {data && (*/}
                        {/*                                    <>*/}
                        {/*                                        {data?.location*/}
                        {/*                                                .country_code &&*/}
                        {/*                                            getFlagEmoji(*/}
                        {/*                                                data*/}
                        {/*                                                    ?.location*/}
                        {/*                                                    .country_code*/}
                        {/*                                            )}{' '}*/}
                        {/*                                        {*/}
                        {/*                                            data?.location*/}
                        {/*                                                .country_code*/}
                        {/*                                        }{' '}*/}
                        {/*                                        (*/}
                        {/*                                        {*/}
                        {/*                                            data?.location*/}
                        {/*                                                .country_code*/}
                        {/*                                        }*/}
                        {/*                                        )*/}
                        {/*                                    </>*/}
                        {/*                                )}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Region'}>*/}
                        {/*                                {data?.location.region_name}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'City'}>*/}
                        {/*                                {data?.location.city_name}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Latitude'}>*/}
                        {/*                                {*/}
                        {/*                                    data?.location.lat_long*/}
                        {/*                                        .latitude*/}
                        {/*                                }*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Longitude'}>*/}
                        {/*                                {*/}
                        {/*                                    data?.location.lat_long*/}
                        {/*                                        .longitude*/}
                        {/*                                }*/}
                        {/*                            </InfoRow>*/}
                        {/*                        </TableBody>*/}
                        {/*                    </Table>*/}
                        {/*                </TableContainer>*/}
                        {/*            </Grid>*/}
                        {/*            <Grid xs={12} md={6} padding={2}>*/}
                        {/*                <MapContainer*/}
                        {/*                    zoom={3}*/}
                        {/*                    scrollWheelZoom={true}*/}
                        {/*                    id={'map'}*/}
                        {/*                    style={{*/}
                        {/*                        height: '400px',*/}
                        {/*                        width: '100%'*/}
                        {/*                    }}*/}
                        {/*                    attributionControl={true}*/}
                        {/*                    minZoom={3}*/}
                        {/*                    worldCopyJump={true}*/}
                        {/*                    center={pos}*/}
                        {/*                >*/}
                        {/*                    <TileLayer*/}
                        {/*                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"*/}
                        {/*                        attribution={*/}
                        {/*                            '© OpenStreetMap contributors '*/}
                        {/*                        }*/}
                        {/*                    />*/}
                        {/*                    {(data?.location.lat_long.latitude ||*/}
                        {/*                        0) > 0 && (*/}
                        {/*                        <Marker*/}
                        {/*                            autoPan={true}*/}
                        {/*                            title={'Location'}*/}
                        {/*                            position={pos}*/}
                        {/*                        />*/}
                        {/*                    )}*/}
                        {/*                </MapContainer>*/}
                        {/*            </Grid>*/}
                        {/*            <Grid xs={12} md={6}>*/}
                        {/*                <Typography variant={'h4'} padding={2}>*/}
                        {/*                    ASN*/}
                        {/*                </Typography>*/}
                        {/*                <TableContainer>*/}
                        {/*                    <Table>*/}
                        {/*                        <TableBody>*/}
                        {/*                            <InfoRow label={'AS Name'}>*/}
                        {/*                                {data?.asn.as_name}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'AS Number'}>*/}
                        {/*                                <Link*/}
                        {/*                                    href={`https://bgpview.io/asn/${data?.asn.as_num}`}*/}
                        {/*                                >*/}
                        {/*                                    {data?.asn.as_num}*/}
                        {/*                                </Link>*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'CIDR Block'}>*/}
                        {/*                                {data?.asn.cidr}*/}
                        {/*                            </InfoRow>*/}
                        {/*                        </TableBody>*/}
                        {/*                    </Table>*/}
                        {/*                </TableContainer>*/}
                        {/*            </Grid>*/}
                        {/*            <Grid xs={12} md={6}>*/}
                        {/*                <Typography variant={'h4'} padding={2}>*/}
                        {/*                    Proxy Info*/}
                        {/*                </Typography>*/}
                        {/*                <TableContainer>*/}
                        {/*                    <Table>*/}
                        {/*                        <TableBody>*/}
                        {/*                            <InfoRow label={'Proxy Type'}>*/}
                        {/*                                {data?.proxy.proxy_type}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'ISP'}>*/}
                        {/*                                {data?.proxy.isp}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Domain'}>*/}
                        {/*                                {data?.proxy.domain}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Usage Type'}>*/}
                        {/*                                {data?.proxy.usage_type}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Last Seen'}>*/}
                        {/*                                {data?.proxy.last_seen}*/}
                        {/*                            </InfoRow>*/}
                        {/*                            <InfoRow label={'Threat'}>*/}
                        {/*                                {data?.proxy.threat}*/}
                        {/*                            </InfoRow>*/}
                        {/*                        </TableBody>*/}
                        {/*                    </Table>*/}
                        {/*                </TableContainer>*/}
                        {/*            </Grid>*/}
                        {/*        </Grid>*/}
                        {/*    </div>*/}
                        {/*)}*/}
                    </Grid>
                </Grid>
            </>
        </ContainerWithHeader>
    );
}
