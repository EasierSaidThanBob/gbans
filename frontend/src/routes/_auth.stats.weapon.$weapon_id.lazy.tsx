import { ChangeEvent, JSX, useState } from 'react';
import { useParams } from 'react-router';
import InsightsIcon from '@mui/icons-material/Insights';
import Grid from '@mui/material/Unstable_Grid2';
import { createLazyFileRoute } from '@tanstack/react-router';
import { PlayerWeaponStats } from '../api';
import { ContainerWithHeader } from '../component/ContainerWithHeader';
import FmtWhenGt from '../component/FmtWhenGT.tsx';
import { LoadingPlaceholder } from '../component/LoadingPlaceholder';
import { PersonCell } from '../component/PersonCell';
import { LazyTable } from '../component/table/LazyTable';
import { useWeaponsStats } from '../hooks/useWeaponsStats';
import { Order, RowsPerPage } from '../util/table.ts';
import { defaultFloatFmtPct, humanCount } from '../util/text.tsx';

export const Route = createLazyFileRoute('/_auth/stats/weapon/$weapon_id')({
    component: StatsWeapon
});

interface WeaponStatsContainerProps {
    weapon_id: number;
}

const WeaponStatsContainer = ({ weapon_id }: WeaponStatsContainerProps) => {
    const [page, setPage] = useState(0);
    const [sortOrder, setSortOrder] = useState<Order>('asc');
    const [rows, setRows] = useState<RowsPerPage>(RowsPerPage.TwentyFive);
    const [sortColumn, setSortColumn] = useState<keyof PlayerWeaponStats>('rank');

    const { data, weapon, loading, count } = useWeaponsStats(weapon_id, {
        offset: page * rows,
        limit: rows,
        order_by: sortColumn,
        desc: sortOrder == 'desc'
    });

    return (
        <ContainerWithHeader title={`Top 250 Weapon Users: ${weapon?.name}`} iconLeft={<InsightsIcon />}>
            {loading ? (
                <LoadingPlaceholder />
            ) : (
                <LazyTable<PlayerWeaponStats>
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
                    onRowsPerPageChange={(event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                        setRows(Number(event.target.value));
                        setPage(0);
                    }}
                    columns={[
                        {
                            label: '#',
                            sortable: true,
                            sortKey: 'rank',
                            tooltip: 'Overall Rank',
                            align: 'center',
                            renderer: (obj) => obj.rank
                        },
                        {
                            label: 'Player Name',
                            sortable: true,
                            sortKey: 'personaname',
                            tooltip: 'Player Name',
                            align: 'left',
                            renderer: (obj) => (
                                <PersonCell avatar_hash={obj.avatar_hash} personaname={obj.personaname} steam_id={obj.steam_id} />
                            )
                        },
                        {
                            label: 'Kills',
                            sortable: true,
                            sortKey: 'kills',
                            tooltip: 'Total Kills',
                            renderer: (obj) => FmtWhenGt(obj.kills, humanCount)
                        },
                        {
                            label: 'Dmg',
                            sortable: true,
                            sortKey: 'damage',
                            tooltip: 'Total Damage',
                            renderer: (obj) => FmtWhenGt(obj.damage, humanCount)
                        },
                        {
                            label: 'Shots',
                            sortable: true,
                            sortKey: 'shots',
                            tooltip: 'Total Shots',
                            renderer: (obj) => FmtWhenGt(obj.shots, humanCount)
                        },
                        {
                            label: 'Hits',
                            sortable: true,
                            sortKey: 'hits',
                            tooltip: 'Total Shots Landed',
                            renderer: (obj) => FmtWhenGt(obj.hits, humanCount)
                        },
                        {
                            label: 'Acc%',
                            sortable: false,
                            virtual: true,
                            virtualKey: 'accuracy',
                            tooltip: 'Overall Accuracy',
                            renderer: (obj) => FmtWhenGt(obj.shots, () => defaultFloatFmtPct(obj.accuracy))
                        },
                        {
                            label: 'As',
                            sortable: true,
                            sortKey: 'airshots',
                            tooltip: 'Total Airshots',
                            renderer: (obj) => FmtWhenGt(obj.airshots, humanCount)
                        },

                        {
                            label: 'Bs',
                            sortable: true,
                            sortKey: 'backstabs',
                            tooltip: 'Total Backstabs',
                            renderer: (obj) => FmtWhenGt(obj.backstabs, humanCount)
                        },

                        {
                            label: 'Hs',
                            sortable: true,
                            sortKey: 'headshots',
                            tooltip: 'Total Headshots',
                            renderer: (obj) => FmtWhenGt(obj.headshots, humanCount)
                        }
                    ]}
                />
            )}
        </ContainerWithHeader>
    );
};

function StatsWeapon() {
    const { weapon_id } = useParams();

    return (
        <Grid container spacing={2}>
            <Grid xs={12}>
                <WeaponStatsContainer weapon_id={Number(weapon_id ?? '0')} />
            </Grid>
        </Grid>
    );
}
