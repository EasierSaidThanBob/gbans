import { ChangeEvent, useState } from 'react';
import InsightsIcon from '@mui/icons-material/Insights';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import { Link as RouterLink } from '@tanstack/react-router';
import { WeaponsOverallResult } from '../api';
import { useWeaponsOverallStats } from '../hooks/useWeaponsOverallStats';
import { Order, RowsPerPage } from '../util/table.ts';
import { defaultFloatFmtPct, humanCount } from '../util/text.tsx';
import { ContainerWithHeader } from './ContainerWithHeader';
import FmtWhenGt from './FmtWhenGT.tsx';
import { LoadingPlaceholder } from './LoadingPlaceholder';
import { LazyTable } from './table/LazyTable';

export const WeaponsStatListContainer = () => {
    const [page, setPage] = useState(0);
    const [sortOrder, setSortOrder] = useState<Order>('desc');
    const [rows, setRows] = useState<RowsPerPage>(RowsPerPage.TwentyFive);
    const [sortColumn, setSortColumn] =
        useState<keyof WeaponsOverallResult>('kills');

    const { data, loading, count } = useWeaponsOverallStats({
        offset: page * rows,
        limit: rows,
        order_by: sortColumn,
        desc: sortOrder == 'desc'
    });

    return (
        <ContainerWithHeader
            title={'Overall Weapon Stats'}
            iconLeft={<InsightsIcon />}
        >
            {loading ? (
                <LoadingPlaceholder />
            ) : (
                <LazyTable<WeaponsOverallResult>
                    showPager={true}
                    count={count}
                    rows={data}
                    page={Number(page ?? 0)}
                    rowsPerPage={rows}
                    sortOrder={sortOrder}
                    sortColumn={sortColumn}
                    onSortColumnChanged={async (column) => {
                        setSortColumn(column);
                    }}
                    onSortOrderChanged={async (direction) => {
                        setSortOrder(direction);
                    }}
                    onPageChange={(_, newPage: number) => {
                        setPage(newPage);
                    }}
                    onRowsPerPageChange={(
                        event: ChangeEvent<
                            HTMLInputElement | HTMLTextAreaElement
                        >
                    ) => {
                        setRows(Number(event.target.value));
                        setPage(0);
                    }}
                    columns={[
                        {
                            label: 'Weapon Name',
                            sortable: true,
                            sortKey: 'name',
                            tooltip: 'Weapon Name',
                            align: 'left',
                            renderer: (obj) => {
                                return (
                                    <Tooltip title={obj.key}>
                                        <Button
                                            fullWidth
                                            size={'small'}
                                            variant={'text'}
                                            style={{
                                                justifyContent: 'flex-start'
                                            }}
                                            component={RouterLink}
                                            to={`/stats/weapon/${obj.weapon_id}`}
                                            href={`/stats/weapon/${obj.weapon_id}`}
                                        >
                                            {obj.name}
                                        </Button>
                                    </Tooltip>
                                );
                            }
                        },
                        {
                            label: 'Kills',
                            sortable: true,
                            sortKey: 'kills',
                            tooltip: 'Total Kills',
                            renderer: (obj) => FmtWhenGt(obj.kills, humanCount)
                        },
                        {
                            label: 'Kills%',
                            sortable: true,
                            sortKey: 'kills_pct',
                            tooltip: 'Percentage Of Overall Kills',
                            renderer: (obj) =>
                                FmtWhenGt(obj.kills_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Shots',
                            sortable: true,
                            sortKey: 'shots',
                            tooltip: 'Total Shots',
                            renderer: (obj) => FmtWhenGt(obj.shots, humanCount)
                        },
                        {
                            label: 'Shots%',
                            sortable: true,
                            sortKey: 'shots_pct',
                            tooltip: 'Shot Hit Percentage',
                            renderer: (obj) =>
                                FmtWhenGt(obj.shots_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Hits',
                            sortable: true,
                            sortKey: 'hits',
                            tooltip: 'Total Hits',
                            renderer: (obj) => FmtWhenGt(obj.hits, humanCount)
                        },
                        {
                            label: 'Hits%',
                            sortable: true,
                            sortKey: 'hits_pct',
                            tooltip: 'Total Hits',
                            renderer: (obj) =>
                                FmtWhenGt(obj.hits_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Acc%',
                            sortable: false,
                            virtual: true,
                            virtualKey: 'accuracy',
                            tooltip: 'Overall Accuracy',
                            renderer: (obj) =>
                                FmtWhenGt(obj.shots_pct, () =>
                                    defaultFloatFmtPct(
                                        (obj.hits / obj.shots) * 100
                                    )
                                )
                        },
                        {
                            label: 'As',
                            sortable: true,
                            sortKey: 'airshots',
                            tooltip: 'Total Airshots',
                            renderer: (obj) =>
                                FmtWhenGt(obj.airshots, humanCount)
                        },
                        {
                            label: 'As%',
                            sortable: true,
                            sortKey: 'airshots_pct',
                            tooltip: 'Total Airshot Percentage',
                            renderer: (obj) =>
                                FmtWhenGt(obj.airshots_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Bs',
                            sortable: true,
                            sortKey: 'backstabs',
                            tooltip: 'Total Backstabs',
                            renderer: (obj) =>
                                FmtWhenGt(obj.backstabs, humanCount)
                        },
                        {
                            label: 'Bs%',
                            sortable: true,
                            sortKey: 'backstabs_pct',
                            tooltip: 'Total Backstabs Percentage',
                            renderer: (obj) =>
                                FmtWhenGt(obj.backstabs_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Hs',
                            sortable: true,
                            sortKey: 'headshots',
                            tooltip: 'Total Headshots',
                            renderer: (obj) =>
                                FmtWhenGt(obj.headshots, humanCount)
                        },
                        {
                            label: 'Hs%',
                            sortable: true,
                            sortKey: 'headshots_pct',
                            tooltip: 'Total Headshot Percentage',
                            renderer: (obj) =>
                                FmtWhenGt(obj.headshots_pct, defaultFloatFmtPct)
                        },
                        {
                            label: 'Dmg',
                            sortable: true,
                            sortKey: 'damage',
                            tooltip: 'Total Damage',
                            renderer: (obj) => FmtWhenGt(obj.damage, humanCount)
                        },
                        {
                            label: 'Dmg%',
                            sortable: true,
                            sortKey: 'damage_pct',
                            tooltip: 'Total Damage Percentage',
                            renderer: (obj) =>
                                FmtWhenGt(obj.damage_pct, defaultFloatFmtPct)
                        }
                    ]}
                />
            )}
        </ContainerWithHeader>
    );
};
