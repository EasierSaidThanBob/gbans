import AddIcon from '@mui/icons-material/Add';
import FilterListIcon from '@mui/icons-material/FilterList';
import GavelIcon from '@mui/icons-material/Gavel';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import { useForm } from '@tanstack/react-form';
import { useQuery } from '@tanstack/react-query';
import { createFileRoute, useNavigate } from '@tanstack/react-router';
import { createColumnHelper, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { apiGetBansASN, ASNBanRecord, BanReasons } from '../api';
import { ContainerWithHeader } from '../component/ContainerWithHeader.tsx';
import { ContainerWithHeaderAndButtons } from '../component/ContainerWithHeaderAndButtons.tsx';
import { DataTable } from '../component/DataTable.tsx';
import { PersonCell } from '../component/PersonCell.tsx';
import RouterLink from '../component/RouterLink.tsx';
import { TableCellRelativeDateField } from '../component/TableCellRelativeDateField.tsx';
import { TableHeadingCell } from '../component/TableHeadingCell.tsx';
import { Buttons } from '../component/field/Buttons.tsx';
import { CheckboxSimple } from '../component/field/CheckboxSimple.tsx';
import { TextFieldSimple } from '../component/field/TextFieldSimple.tsx';
import { commonTableSearchSchema, isPermanentBan, LazyResult, RowsPerPage } from '../util/table.ts';
import { renderDate } from '../util/text.tsx';
import { makeSteamidValidatorsOptional } from '../util/validator/makeSteamidValidatorsOptional.ts';

const banASNSearchSchema = z.object({
    ...commonTableSearchSchema,
    sortColumn: z
        .enum(['ban_asn_id', 'source_id', 'target_id', 'deleted', 'reason', 'as_num', 'valid_until'])
        .optional(),
    source_id: z.string().optional(),
    target_id: z.string().optional(),
    as_num: z.string().optional(),
    deleted: z.boolean().optional()
});

export const Route = createFileRoute('/_mod/admin/ban/asn')({
    component: AdminBanASN,
    validateSearch: (search) => banASNSearchSchema.parse(search)
});

function AdminBanASN() {
    const defaultRows = RowsPerPage.TwentyFive;

    const navigate = useNavigate({ from: Route.fullPath });
    const { page, rows, deleted, as_num, sortOrder, sortColumn, target_id, source_id } = Route.useSearch();
    const { data: bans, isLoading } = useQuery({
        queryKey: ['steamBans', { page, rows, sortOrder, sortColumn, target_id, source_id }],
        queryFn: async () => {
            return await apiGetBansASN({
                limit: rows ?? defaultRows,
                offset: (page ?? 0) * (rows ?? defaultRows),
                order_by: sortColumn ?? 'ban_asn_id',
                desc: sortOrder == 'desc',
                source_id: source_id,
                target_id: target_id,
                as_num: as_num ? Number(as_num) : undefined,
                deleted: deleted ?? false
            });
        }
    });

    const { Field, Subscribe, handleSubmit, reset } = useForm({
        onSubmit: async ({ value }) => {
            await navigate({ to: '/admin/ban/asn', search: (prev) => ({ ...prev, ...value }) });
        },
        validatorAdapter: zodValidator,
        validators: {
            onChange: banASNSearchSchema
        },
        defaultValues: {
            source_id: source_id ?? '',
            target_id: target_id ?? '',
            as_num: as_num ?? '',
            deleted: deleted ?? false
        }
    });

    const clear = async () => {
        await navigate({
            to: '/admin/ban/asn',
            search: (prev) => ({
                ...prev,
                source_id: undefined,
                target_id: undefined,
                as_num: undefined,
                deleted: undefined
            })
        });
    };

    // const [newASNBans, setNewASNBans] = useState<ASNBanRecord[]>([]);

    // const onNewBanASN = useCallback(async () => {
    //     try {
    //         const ban = await NiceModal.show<ASNBanRecord>(ModalBanASN, {});
    //         setNewASNBans((prevState) => {
    //             return [ban, ...prevState];
    //         });
    //         sendFlash('success', `Created ASN ban successfully #${ban.ban_asn_id}`);
    //     } catch (e) {
    //         logErr(e);
    //     }
    // }, [sendFlash]);

    return (
        <Grid container spacing={2}>
            <Grid xs={12}>
                <ContainerWithHeader title={'Filters'} iconLeft={<FilterListIcon />} marginTop={2}>
                    <form
                        onSubmit={async (e) => {
                            e.preventDefault();
                            e.stopPropagation();
                            await handleSubmit();
                        }}
                    >
                        <Grid container spacing={2}>
                            <Grid xs={4}>
                                <Grid xs={6} md={3}>
                                    <Field
                                        name={'source_id'}
                                        validators={makeSteamidValidatorsOptional()}
                                        children={(props) => {
                                            return (
                                                <TextFieldSimple
                                                    {...props}
                                                    label={'Author Steam ID'}
                                                    fullwidth={true}
                                                />
                                            );
                                        }}
                                    />
                                </Grid>
                            </Grid>
                            <Grid xs={6} md={3}>
                                <Field
                                    name={'target_id'}
                                    validators={makeSteamidValidatorsOptional()}
                                    children={(props) => {
                                        return (
                                            <TextFieldSimple {...props} label={'Subject Steam ID'} fullwidth={true} />
                                        );
                                    }}
                                />
                            </Grid>

                            <Grid xs={6} md={3}>
                                <Field
                                    name={'as_num'}
                                    children={(props) => {
                                        return <TextFieldSimple {...props} label={'AS Number'} fullwidth={true} />;
                                    }}
                                />
                            </Grid>

                            <Grid xs="auto">
                                <Field
                                    name={'deleted'}
                                    children={(props) => {
                                        return <CheckboxSimple {...props} label={'Show Deleted'} />;
                                    }}
                                />
                            </Grid>

                            <Grid xs={12} mdOffset="auto">
                                <Subscribe
                                    selector={(state) => [state.canSubmit, state.isSubmitting]}
                                    children={([canSubmit, isSubmitting]) => (
                                        <Buttons
                                            reset={reset}
                                            canSubmit={canSubmit}
                                            isSubmitting={isSubmitting}
                                            onClear={clear}
                                        />
                                    )}
                                />
                            </Grid>
                        </Grid>
                    </form>
                </ContainerWithHeader>
            </Grid>

            <Grid xs={12}>
                <ContainerWithHeaderAndButtons
                    title={'ASN Ban History'}
                    marginTop={0}
                    iconLeft={<GavelIcon />}
                    buttons={[
                        <Button
                            key={'btn-asn'}
                            variant={'contained'}
                            color={'success'}
                            startIcon={<AddIcon />}
                            sx={{ marginRight: 2 }}
                            // onClick={onNewBanASN}
                        >
                            Create
                        </Button>
                    ]}
                >
                    {/*<Formik*/}
                    {/*    initialValues={{*/}
                    {/*        as_num: Number(state.asNum),*/}
                    {/*        source_id: state.source,*/}
                    {/*        target_id: state.target,*/}
                    {/*        deleted: Boolean(state.deleted)*/}
                    {/*    }}*/}
                    {/*    onReset={onReset}*/}
                    {/*    onSubmit={onSubmit}*/}
                    {/*    validationSchema={validationSchema}*/}
                    {/*    validateOnChange={true}*/}
                    {/*>*/}
                    {/*    <Grid container spacing={3}>*/}
                    {/*        <Grid xs={12}>*/}
                    {/*            <Grid container spacing={2}>*/}
                    {/*                <Grid xs={4} sm={3} md={2}>*/}
                    {/*                    <ASNumberField />*/}
                    {/*                </Grid>*/}
                    {/*                <Grid xs={4} sm={3} md={2}>*/}
                    {/*                    <SourceIDField />*/}
                    {/*                </Grid>*/}
                    {/*                <Grid xs={4} sm={3} md={2}>*/}
                    {/*                    <TargetIDField />*/}
                    {/*                </Grid>*/}
                    {/*                <Grid xs={4} sm={3} md={2}>*/}
                    {/*                    <DeletedField />*/}
                    {/*                </Grid>*/}
                    {/*                <Grid xs={4} sm={3} md={2}>*/}
                    {/*                    <FilterButtons />*/}
                    {/*                </Grid>*/}
                    {/*            </Grid>*/}
                    {/*        </Grid>*/}
                    {/*        <Grid xs={12}>*/}
                    {/*<LazyTable<ASNBanRecord>*/}
                    {/*    showPager={true}*/}
                    {/*    count={count}*/}
                    {/*    rows={allBans}*/}
                    {/*    page={Number(state.page ?? 0)}*/}
                    {/*    rowsPerPage={Number(state.rows ?? RowsPerPage.Ten)}*/}
                    {/*    sortOrder={state.sortOrder}*/}
                    {/*    sortColumn={state.sortColumn}*/}
                    {/*    onSortColumnChanged={async (column) => {*/}
                    {/*        setState({ sortColumn: column });*/}
                    {/*    }}*/}
                    {/*    onSortOrderChanged={async (direction) => {*/}
                    {/*        setState({ sortOrder: direction });*/}
                    {/*    }}*/}
                    {/*    onPageChange={(_, newPage: number) => {*/}
                    {/*        setState({ page: newPage });*/}
                    {/*    }}*/}
                    {/*    onRowsPerPageChange={(*/}
                    {/*        event: ChangeEvent<*/}
                    {/*            HTMLInputElement | HTMLTextAreaElement*/}
                    {/*        >*/}
                    {/*    ) => {*/}
                    {/*        setState({*/}
                    {/*            rows: Number(event.target.value),*/}
                    {/*            page: 0*/}
                    {/*        });*/}
                    {/*    }}*/}
                    {/*    columns={[*/}
                    {/*        {*/}
                    {/*            label: '#',*/}
                    {/*            tooltip: 'Ban ID',*/}
                    {/*            sortKey: 'ban_asn_id',*/}
                    {/*            sortable: true,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (obj) => (*/}
                    {/*                <Typography variant={'body1'}>*/}
                    {/*                    #{obj.ban_asn_id.toString()}*/}
                    {/*                </Typography>*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'A',*/}
                    {/*            tooltip: 'Ban Author Name',*/}
                    {/*            sortKey: 'source_personaname',*/}
                    {/*            sortable: true,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <SteamIDSelectField*/}
                    {/*                    steam_id={row.source_id}*/}
                    {/*                    personaname={row.source_personaname}*/}
                    {/*                    avatarhash={row.source_avatarhash}*/}
                    {/*                    field_name={'source_id'}*/}
                    {/*                />*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Name',*/}
                    {/*            tooltip: 'Persona Name',*/}
                    {/*            sortKey: 'target_personaname',*/}
                    {/*            sortable: true,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <SteamIDSelectField*/}
                    {/*                    steam_id={row.target_id}*/}
                    {/*                    personaname={row.target_personaname}*/}
                    {/*                    avatarhash={row.target_avatarhash}*/}
                    {/*                    field_name={'target_id'}*/}
                    {/*                />*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'ASN',*/}
                    {/*            tooltip: 'Autonomous System Numbers',*/}
                    {/*            sortKey: 'as_num',*/}
                    {/*            sortable: true,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <Typography variant={'body1'}>*/}
                    {/*                    {row.as_num}*/}
                    {/*                </Typography>*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Reason',*/}
                    {/*            tooltip: 'Reason',*/}
                    {/*            sortKey: 'reason',*/}
                    {/*            sortable: true,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <Tooltip*/}
                    {/*                    title={*/}
                    {/*                        row.reason == BanReason.Custom*/}
                    {/*                            ? row.reason_text*/}
                    {/*                            : BanReason[row.reason]*/}
                    {/*                    }*/}
                    {/*                >*/}
                    {/*                    <Typography variant={'body1'}>*/}
                    {/*                        {BanReason[row.reason]}*/}
                    {/*                    </Typography>*/}
                    {/*                </Tooltip>*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Created',*/}
                    {/*            tooltip: 'Created On',*/}
                    {/*            sortType: 'date',*/}
                    {/*            align: 'left',*/}
                    {/*            width: '150px',*/}
                    {/*            virtual: true,*/}
                    {/*            virtualKey: 'created_on',*/}
                    {/*            renderer: (obj) => {*/}
                    {/*                return (*/}
                    {/*                    <Typography variant={'body1'}>*/}
                    {/*                        {renderDate(obj.created_on)}*/}
                    {/*                    </Typography>*/}
                    {/*                );*/}
                    {/*            }*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Expires',*/}
                    {/*            tooltip: 'Valid Until',*/}
                    {/*            sortType: 'date',*/}
                    {/*            align: 'left',*/}
                    {/*            width: '150px',*/}
                    {/*            virtual: true,*/}
                    {/*            virtualKey: 'valid_until',*/}
                    {/*            sortable: true,*/}
                    {/*            renderer: (obj) => {*/}
                    {/*                return (*/}
                    {/*                    <Typography variant={'body1'}>*/}
                    {/*                        {renderDate(obj.valid_until)}*/}
                    {/*                    </Typography>*/}
                    {/*                );*/}
                    {/*            }*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Duration',*/}
                    {/*            tooltip: 'Total Ban Duration',*/}
                    {/*            sortType: 'number',*/}
                    {/*            align: 'left',*/}
                    {/*            width: '150px',*/}
                    {/*            virtual: true,*/}
                    {/*            virtualKey: 'duration',*/}
                    {/*            renderer: (row) => {*/}
                    {/*                const dur = intervalToDuration({*/}
                    {/*                    start: row.created_on,*/}
                    {/*                    end: row.valid_until*/}
                    {/*                });*/}
                    {/*                const durationText =*/}
                    {/*                    dur.years && dur.years > 5*/}
                    {/*                        ? 'Permanent'*/}
                    {/*                        : formatDuration(dur);*/}
                    {/*                return (*/}
                    {/*                    <Typography*/}
                    {/*                        variant={'body1'}*/}
                    {/*                        overflow={'hidden'}*/}
                    {/*                    >*/}
                    {/*                        {durationText}*/}
                    {/*                    </Typography>*/}
                    {/*                );*/}
                    {/*            }*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'A',*/}
                    {/*            tooltip:*/}
                    {/*                'Is this ban active (not deleted/inactive/unbanned)',*/}
                    {/*            align: 'center',*/}
                    {/*            width: '50px',*/}
                    {/*            sortKey: 'deleted',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <TableCellBool enabled={!row.deleted} />*/}
                    {/*            )*/}
                    {/*        },*/}
                    {/*        {*/}
                    {/*            label: 'Act.',*/}
                    {/*            tooltip: 'Actions',*/}
                    {/*            sortKey: 'reason',*/}
                    {/*            sortable: false,*/}
                    {/*            align: 'left',*/}
                    {/*            renderer: (row) => (*/}
                    {/*                <ButtonGroup fullWidth>*/}
                    {/*                    <IconButton*/}
                    {/*                        color={'warning'}*/}
                    {/*                        onClick={async () =>*/}
                    {/*                            await onEditASN(row)*/}
                    {/*                        }*/}
                    {/*                    >*/}
                    {/*                        <Tooltip title={'Edit ASN Ban'}>*/}
                    {/*                            <EditIcon />*/}
                    {/*                        </Tooltip>*/}
                    {/*                    </IconButton>*/}
                    {/*                    <IconButton*/}
                    {/*                        color={'success'}*/}
                    {/*                        onClick={async () =>*/}
                    {/*                            await onUnbanASN(row.as_num)*/}
                    {/*                        }*/}
                    {/*                    >*/}
                    {/*                        <Tooltip title={'Remove CIDR Ban'}>*/}
                    {/*                            <UndoIcon />*/}
                    {/*                        </Tooltip>*/}
                    {/*                    </IconButton>*/}
                    {/*                </ButtonGroup>*/}
                    {/*            )*/}
                    {/*        }*/}
                    {/*    ]}*/}
                    {/*/>*/}
                    <BanASMTable bans={bans ?? { data: [], count: 0 }} isLoading={isLoading} />
                </ContainerWithHeaderAndButtons>
            </Grid>
        </Grid>
    );
}

const columnHelper = createColumnHelper<ASNBanRecord>();

const BanASMTable = ({ bans, isLoading }: { bans: LazyResult<ASNBanRecord>; isLoading: boolean }) => {
    const columns = [
        columnHelper.accessor('ban_asn_id', {
            header: () => <TableHeadingCell name={'Ban ID'} />,
            cell: (info) => (
                <Link component={RouterLink} to={`/ban/$ban_id`} params={{ ban_id: info.getValue() }}>
                    {`#${info.getValue()}`}
                </Link>
            )
        }),
        columnHelper.accessor('source_id', {
            header: () => <TableHeadingCell name={'Author'} />,
            cell: (info) => (
                <PersonCell
                    steam_id={bans.data[info.row.index].source_id}
                    personaname={bans.data[info.row.index].source_personaname}
                    avatar_hash={bans.data[info.row.index].source_avatarhash}
                />
            )
        }),
        columnHelper.accessor('target_id', {
            header: () => <TableHeadingCell name={'Subject'} />,
            cell: (info) => (
                <PersonCell
                    steam_id={bans.data[info.row.index].target_id}
                    personaname={bans.data[info.row.index].target_personaname}
                    avatar_hash={bans.data[info.row.index].target_avatarhash}
                />
            )
        }),
        columnHelper.accessor('as_num', {
            header: () => <TableHeadingCell name={'ASN'} />,
            cell: (info) => <Typography>{info.getValue()}</Typography>
        }),
        columnHelper.accessor('reason', {
            header: () => <TableHeadingCell name={'Reason'} />,
            cell: (info) => <Typography>{BanReasons[info.getValue()]}</Typography>
        }),
        columnHelper.accessor('created_on', {
            header: () => <TableHeadingCell name={'Created'} />,
            cell: (info) => <Typography>{renderDate(info.getValue())}</Typography>
        }),
        columnHelper.accessor('valid_until', {
            header: () => <TableHeadingCell name={'Expires'} />,
            cell: (info) =>
                isPermanentBan(bans.data[info.row.index].created_on, bans.data[info.row.index].valid_until) ? (
                    'Permanent'
                ) : (
                    <TableCellRelativeDateField
                        date={bans.data[info.row.index].created_on}
                        compareDate={bans.data[info.row.index].valid_until}
                    />
                )
        })
    ];

    const table = useReactTable({
        data: bans.data,
        columns: columns,
        getCoreRowModel: getCoreRowModel(),
        manualPagination: true,
        autoResetPageIndex: true
    });

    return <DataTable table={table} isLoading={isLoading} />;
};
