/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as ModImport } from './routes/_mod'
import { Route as GuestImport } from './routes/_guest'
import { Route as AuthImport } from './routes/_auth'
import { Route as AdminImport } from './routes/_admin'
import { Route as GuestIndexImport } from './routes/_guest.index'
import { Route as GuestStvImport } from './routes/_guest.stv'
import { Route as AuthPermissionImport } from './routes/_auth.permission'
import { Route as AuthChatlogsImport } from './routes/_auth.chatlogs'
import { Route as GuestLoginIndexImport } from './routes/_guest.login.index'
import { Route as AuthReportIndexImport } from './routes/_auth.report.index'
import { Route as ModAdminVotesImport } from './routes/_mod.admin.votes'
import { Route as ModAdminReportsImport } from './routes/_mod.admin.reports'
import { Route as ModAdminPeopleImport } from './routes/_mod.admin.people'
import { Route as ModAdminFiltersImport } from './routes/_mod.admin.filters'
import { Route as ModAdminContestsImport } from './routes/_mod.admin.contests'
import { Route as ModAdminAppealsImport } from './routes/_mod.admin.appeals'
import { Route as GuestLoginSuccessImport } from './routes/_guest.login.success'
import { Route as ModAdminNetworkIndexImport } from './routes/_mod.admin.network.index'
import { Route as ModAdminNetworkPlayersbyipImport } from './routes/_mod.admin.network.playersbyip'
import { Route as ModAdminNetworkIphistImport } from './routes/_mod.admin.network.iphist'
import { Route as ModAdminNetworkIpInfoImport } from './routes/_mod.admin.network.ipInfo'
import { Route as ModAdminNetworkCidrblocksImport } from './routes/_mod.admin.network.cidrblocks'
import { Route as ModAdminBanSteamImport } from './routes/_mod.admin.ban.steam'
import { Route as ModAdminBanGroupImport } from './routes/_mod.admin.ban.group'
import { Route as ModAdminBanCidrImport } from './routes/_mod.admin.ban.cidr'
import { Route as ModAdminBanAsnImport } from './routes/_mod.admin.ban.asn'

// Create Virtual Routes

const GuestServersLazyImport = createFileRoute('/_guest/servers')()
const GuestPrivacyPolicyLazyImport = createFileRoute('/_guest/privacy-policy')()
const GuestContestsLazyImport = createFileRoute('/_guest/contests')()
const AuthSettingsLazyImport = createFileRoute('/_auth/settings')()
const AuthPatreonLazyImport = createFileRoute('/_auth/patreon')()
const AuthPageNotFoundLazyImport = createFileRoute('/_auth/page-not-found')()
const AuthNotificationsLazyImport = createFileRoute('/_auth/notifications')()
const AuthLogoutLazyImport = createFileRoute('/_auth/logout')()
const GuestWikiIndexLazyImport = createFileRoute('/_guest/wiki/')()
const AuthStatsIndexLazyImport = createFileRoute('/_auth/stats/')()
const AuthForumsIndexLazyImport = createFileRoute('/_auth/forums/')()
const ModAdminNewsLazyImport = createFileRoute('/_mod/admin/news')()
const GuestWikiSlugLazyImport = createFileRoute('/_guest/wiki/$slug')()
const GuestProfileSteamIdLazyImport = createFileRoute(
  '/_guest/profile/$steamId',
)()
const AuthReportReportIdLazyImport = createFileRoute(
  '/_auth/report/$reportId',
)()
const AuthMatchMatchIdLazyImport = createFileRoute('/_auth/match/$matchId')()
const AuthLogsSteamIdLazyImport = createFileRoute('/_auth/logs/$steamId')()
const AuthLoginDiscordLazyImport = createFileRoute('/_auth/login/discord')()
const AuthForumsForumidLazyImport = createFileRoute('/_auth/forums/$forum_id')()
const AuthContestsContestidLazyImport = createFileRoute(
  '/_auth/contests/$contest_id',
)()
const AuthBanBanidLazyImport = createFileRoute('/_auth/ban/$ban_id')()
const AdminAdminServersLazyImport = createFileRoute('/_admin/admin/servers')()
const AuthStatsWeaponWeaponidLazyImport = createFileRoute(
  '/_auth/stats/weapon/$weapon_id',
)()
const AuthStatsPlayerSteamidLazyImport = createFileRoute(
  '/_auth/stats/player/$steam_id',
)()
const AuthForumsThreadForumthreadidLazyImport = createFileRoute(
  '/_auth/forums/thread/$forum_thread_id',
)()

// Create/Update Routes

const ModRoute = ModImport.update({
  id: '/_mod',
  getParentRoute: () => rootRoute,
} as any)

const GuestRoute = GuestImport.update({
  id: '/_guest',
  getParentRoute: () => rootRoute,
} as any)

const AuthRoute = AuthImport.update({
  id: '/_auth',
  getParentRoute: () => rootRoute,
} as any)

const AdminRoute = AdminImport.update({
  id: '/_admin',
  getParentRoute: () => rootRoute,
} as any)

const GuestIndexRoute = GuestIndexImport.update({
  path: '/',
  getParentRoute: () => GuestRoute,
} as any)

const GuestServersLazyRoute = GuestServersLazyImport.update({
  path: '/servers',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.servers.lazy').then((d) => d.Route),
)

const GuestPrivacyPolicyLazyRoute = GuestPrivacyPolicyLazyImport.update({
  path: '/privacy-policy',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.privacy-policy.lazy').then((d) => d.Route),
)

const GuestContestsLazyRoute = GuestContestsLazyImport.update({
  path: '/contests',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.contests.lazy').then((d) => d.Route),
)

const AuthSettingsLazyRoute = AuthSettingsLazyImport.update({
  path: '/settings',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.settings.lazy').then((d) => d.Route),
)

const AuthPatreonLazyRoute = AuthPatreonLazyImport.update({
  path: '/patreon',
  getParentRoute: () => AuthRoute,
} as any).lazy(() => import('./routes/_auth.patreon.lazy').then((d) => d.Route))

const AuthPageNotFoundLazyRoute = AuthPageNotFoundLazyImport.update({
  path: '/page-not-found',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.page-not-found.lazy').then((d) => d.Route),
)

const AuthNotificationsLazyRoute = AuthNotificationsLazyImport.update({
  path: '/notifications',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.notifications.lazy').then((d) => d.Route),
)

const AuthLogoutLazyRoute = AuthLogoutLazyImport.update({
  path: '/logout',
  getParentRoute: () => AuthRoute,
} as any).lazy(() => import('./routes/_auth.logout.lazy').then((d) => d.Route))

const GuestStvRoute = GuestStvImport.update({
  path: '/stv',
  getParentRoute: () => GuestRoute,
} as any)

const AuthPermissionRoute = AuthPermissionImport.update({
  path: '/permission',
  getParentRoute: () => AuthRoute,
} as any)

const AuthChatlogsRoute = AuthChatlogsImport.update({
  path: '/chatlogs',
  getParentRoute: () => AuthRoute,
} as any)

const GuestWikiIndexLazyRoute = GuestWikiIndexLazyImport.update({
  path: '/wiki/',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.wiki.index.lazy').then((d) => d.Route),
)

const AuthStatsIndexLazyRoute = AuthStatsIndexLazyImport.update({
  path: '/stats/',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.stats.index.lazy').then((d) => d.Route),
)

const AuthForumsIndexLazyRoute = AuthForumsIndexLazyImport.update({
  path: '/forums/',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.forums.index.lazy').then((d) => d.Route),
)

const GuestLoginIndexRoute = GuestLoginIndexImport.update({
  path: '/login/',
  getParentRoute: () => GuestRoute,
} as any)

const AuthReportIndexRoute = AuthReportIndexImport.update({
  path: '/report/',
  getParentRoute: () => AuthRoute,
} as any)

const ModAdminNewsLazyRoute = ModAdminNewsLazyImport.update({
  path: '/admin/news',
  getParentRoute: () => ModRoute,
} as any).lazy(() =>
  import('./routes/_mod.admin.news.lazy').then((d) => d.Route),
)

const GuestWikiSlugLazyRoute = GuestWikiSlugLazyImport.update({
  path: '/wiki/$slug',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.wiki.$slug.lazy').then((d) => d.Route),
)

const GuestProfileSteamIdLazyRoute = GuestProfileSteamIdLazyImport.update({
  path: '/profile/$steamId',
  getParentRoute: () => GuestRoute,
} as any).lazy(() =>
  import('./routes/_guest.profile.$steamId.lazy').then((d) => d.Route),
)

const AuthReportReportIdLazyRoute = AuthReportReportIdLazyImport.update({
  path: '/report/$reportId',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.report.$reportId.lazy').then((d) => d.Route),
)

const AuthMatchMatchIdLazyRoute = AuthMatchMatchIdLazyImport.update({
  path: '/match/$matchId',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.match.$matchId.lazy').then((d) => d.Route),
)

const AuthLogsSteamIdLazyRoute = AuthLogsSteamIdLazyImport.update({
  path: '/logs/$steamId',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.logs.$steamId.lazy').then((d) => d.Route),
)

const AuthLoginDiscordLazyRoute = AuthLoginDiscordLazyImport.update({
  path: '/login/discord',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.login.discord.lazy').then((d) => d.Route),
)

const AuthForumsForumidLazyRoute = AuthForumsForumidLazyImport.update({
  path: '/forums/$forum_id',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.forums.$forum_id.lazy').then((d) => d.Route),
)

const AuthContestsContestidLazyRoute = AuthContestsContestidLazyImport.update({
  path: '/contests/$contest_id',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.contests.$contest_id.lazy').then((d) => d.Route),
)

const AuthBanBanidLazyRoute = AuthBanBanidLazyImport.update({
  path: '/ban/$ban_id',
  getParentRoute: () => AuthRoute,
} as any).lazy(() =>
  import('./routes/_auth.ban.$ban_id.lazy').then((d) => d.Route),
)

const AdminAdminServersLazyRoute = AdminAdminServersLazyImport.update({
  path: '/admin/servers',
  getParentRoute: () => AdminRoute,
} as any).lazy(() =>
  import('./routes/_admin.admin.servers.lazy').then((d) => d.Route),
)

const ModAdminVotesRoute = ModAdminVotesImport.update({
  path: '/admin/votes',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminReportsRoute = ModAdminReportsImport.update({
  path: '/admin/reports',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminPeopleRoute = ModAdminPeopleImport.update({
  path: '/admin/people',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminFiltersRoute = ModAdminFiltersImport.update({
  path: '/admin/filters',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminContestsRoute = ModAdminContestsImport.update({
  path: '/admin/contests',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminAppealsRoute = ModAdminAppealsImport.update({
  path: '/admin/appeals',
  getParentRoute: () => ModRoute,
} as any)

const GuestLoginSuccessRoute = GuestLoginSuccessImport.update({
  path: '/login/success',
  getParentRoute: () => GuestRoute,
} as any)

const ModAdminNetworkIndexRoute = ModAdminNetworkIndexImport.update({
  path: '/admin/network/',
  getParentRoute: () => ModRoute,
} as any)

const AuthStatsWeaponWeaponidLazyRoute =
  AuthStatsWeaponWeaponidLazyImport.update({
    path: '/stats/weapon/$weapon_id',
    getParentRoute: () => AuthRoute,
  } as any).lazy(() =>
    import('./routes/_auth.stats.weapon.$weapon_id.lazy').then((d) => d.Route),
  )

const AuthStatsPlayerSteamidLazyRoute = AuthStatsPlayerSteamidLazyImport.update(
  {
    path: '/stats/player/$steam_id',
    getParentRoute: () => AuthRoute,
  } as any,
).lazy(() =>
  import('./routes/_auth.stats.player.$steam_id.lazy').then((d) => d.Route),
)

const AuthForumsThreadForumthreadidLazyRoute =
  AuthForumsThreadForumthreadidLazyImport.update({
    path: '/forums/thread/$forum_thread_id',
    getParentRoute: () => AuthRoute,
  } as any).lazy(() =>
    import('./routes/_auth.forums.thread.$forum_thread_id.lazy').then(
      (d) => d.Route,
    ),
  )

const ModAdminNetworkPlayersbyipRoute = ModAdminNetworkPlayersbyipImport.update(
  {
    path: '/admin/network/playersbyip',
    getParentRoute: () => ModRoute,
  } as any,
)

const ModAdminNetworkIphistRoute = ModAdminNetworkIphistImport.update({
  path: '/admin/network/iphist',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminNetworkIpInfoRoute = ModAdminNetworkIpInfoImport.update({
  path: '/admin/network/ipInfo',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminNetworkCidrblocksRoute = ModAdminNetworkCidrblocksImport.update({
  path: '/admin/network/cidrblocks',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminBanSteamRoute = ModAdminBanSteamImport.update({
  path: '/admin/ban/steam',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminBanGroupRoute = ModAdminBanGroupImport.update({
  path: '/admin/ban/group',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminBanCidrRoute = ModAdminBanCidrImport.update({
  path: '/admin/ban/cidr',
  getParentRoute: () => ModRoute,
} as any)

const ModAdminBanAsnRoute = ModAdminBanAsnImport.update({
  path: '/admin/ban/asn',
  getParentRoute: () => ModRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/_admin': {
      preLoaderRoute: typeof AdminImport
      parentRoute: typeof rootRoute
    }
    '/_auth': {
      preLoaderRoute: typeof AuthImport
      parentRoute: typeof rootRoute
    }
    '/_guest': {
      preLoaderRoute: typeof GuestImport
      parentRoute: typeof rootRoute
    }
    '/_mod': {
      preLoaderRoute: typeof ModImport
      parentRoute: typeof rootRoute
    }
    '/_auth/chatlogs': {
      preLoaderRoute: typeof AuthChatlogsImport
      parentRoute: typeof AuthImport
    }
    '/_auth/permission': {
      preLoaderRoute: typeof AuthPermissionImport
      parentRoute: typeof AuthImport
    }
    '/_guest/stv': {
      preLoaderRoute: typeof GuestStvImport
      parentRoute: typeof GuestImport
    }
    '/_auth/logout': {
      preLoaderRoute: typeof AuthLogoutLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/notifications': {
      preLoaderRoute: typeof AuthNotificationsLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/page-not-found': {
      preLoaderRoute: typeof AuthPageNotFoundLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/patreon': {
      preLoaderRoute: typeof AuthPatreonLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/settings': {
      preLoaderRoute: typeof AuthSettingsLazyImport
      parentRoute: typeof AuthImport
    }
    '/_guest/contests': {
      preLoaderRoute: typeof GuestContestsLazyImport
      parentRoute: typeof GuestImport
    }
    '/_guest/privacy-policy': {
      preLoaderRoute: typeof GuestPrivacyPolicyLazyImport
      parentRoute: typeof GuestImport
    }
    '/_guest/servers': {
      preLoaderRoute: typeof GuestServersLazyImport
      parentRoute: typeof GuestImport
    }
    '/_guest/': {
      preLoaderRoute: typeof GuestIndexImport
      parentRoute: typeof GuestImport
    }
    '/_guest/login/success': {
      preLoaderRoute: typeof GuestLoginSuccessImport
      parentRoute: typeof GuestImport
    }
    '/_mod/admin/appeals': {
      preLoaderRoute: typeof ModAdminAppealsImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/contests': {
      preLoaderRoute: typeof ModAdminContestsImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/filters': {
      preLoaderRoute: typeof ModAdminFiltersImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/people': {
      preLoaderRoute: typeof ModAdminPeopleImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/reports': {
      preLoaderRoute: typeof ModAdminReportsImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/votes': {
      preLoaderRoute: typeof ModAdminVotesImport
      parentRoute: typeof ModImport
    }
    '/_admin/admin/servers': {
      preLoaderRoute: typeof AdminAdminServersLazyImport
      parentRoute: typeof AdminImport
    }
    '/_auth/ban/$ban_id': {
      preLoaderRoute: typeof AuthBanBanidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/contests/$contest_id': {
      preLoaderRoute: typeof AuthContestsContestidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/forums/$forum_id': {
      preLoaderRoute: typeof AuthForumsForumidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/login/discord': {
      preLoaderRoute: typeof AuthLoginDiscordLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/logs/$steamId': {
      preLoaderRoute: typeof AuthLogsSteamIdLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/match/$matchId': {
      preLoaderRoute: typeof AuthMatchMatchIdLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/report/$reportId': {
      preLoaderRoute: typeof AuthReportReportIdLazyImport
      parentRoute: typeof AuthImport
    }
    '/_guest/profile/$steamId': {
      preLoaderRoute: typeof GuestProfileSteamIdLazyImport
      parentRoute: typeof GuestImport
    }
    '/_guest/wiki/$slug': {
      preLoaderRoute: typeof GuestWikiSlugLazyImport
      parentRoute: typeof GuestImport
    }
    '/_mod/admin/news': {
      preLoaderRoute: typeof ModAdminNewsLazyImport
      parentRoute: typeof ModImport
    }
    '/_auth/report/': {
      preLoaderRoute: typeof AuthReportIndexImport
      parentRoute: typeof AuthImport
    }
    '/_guest/login/': {
      preLoaderRoute: typeof GuestLoginIndexImport
      parentRoute: typeof GuestImport
    }
    '/_auth/forums/': {
      preLoaderRoute: typeof AuthForumsIndexLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/stats/': {
      preLoaderRoute: typeof AuthStatsIndexLazyImport
      parentRoute: typeof AuthImport
    }
    '/_guest/wiki/': {
      preLoaderRoute: typeof GuestWikiIndexLazyImport
      parentRoute: typeof GuestImport
    }
    '/_mod/admin/ban/asn': {
      preLoaderRoute: typeof ModAdminBanAsnImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/ban/cidr': {
      preLoaderRoute: typeof ModAdminBanCidrImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/ban/group': {
      preLoaderRoute: typeof ModAdminBanGroupImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/ban/steam': {
      preLoaderRoute: typeof ModAdminBanSteamImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/network/cidrblocks': {
      preLoaderRoute: typeof ModAdminNetworkCidrblocksImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/network/ipInfo': {
      preLoaderRoute: typeof ModAdminNetworkIpInfoImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/network/iphist': {
      preLoaderRoute: typeof ModAdminNetworkIphistImport
      parentRoute: typeof ModImport
    }
    '/_mod/admin/network/playersbyip': {
      preLoaderRoute: typeof ModAdminNetworkPlayersbyipImport
      parentRoute: typeof ModImport
    }
    '/_auth/forums/thread/$forum_thread_id': {
      preLoaderRoute: typeof AuthForumsThreadForumthreadidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/stats/player/$steam_id': {
      preLoaderRoute: typeof AuthStatsPlayerSteamidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_auth/stats/weapon/$weapon_id': {
      preLoaderRoute: typeof AuthStatsWeaponWeaponidLazyImport
      parentRoute: typeof AuthImport
    }
    '/_mod/admin/network/': {
      preLoaderRoute: typeof ModAdminNetworkIndexImport
      parentRoute: typeof ModImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  AdminRoute.addChildren([AdminAdminServersLazyRoute]),
  AuthRoute.addChildren([
    AuthChatlogsRoute,
    AuthPermissionRoute,
    AuthLogoutLazyRoute,
    AuthNotificationsLazyRoute,
    AuthPageNotFoundLazyRoute,
    AuthPatreonLazyRoute,
    AuthSettingsLazyRoute,
    AuthBanBanidLazyRoute,
    AuthContestsContestidLazyRoute,
    AuthForumsForumidLazyRoute,
    AuthLoginDiscordLazyRoute,
    AuthLogsSteamIdLazyRoute,
    AuthMatchMatchIdLazyRoute,
    AuthReportReportIdLazyRoute,
    AuthReportIndexRoute,
    AuthForumsIndexLazyRoute,
    AuthStatsIndexLazyRoute,
    AuthForumsThreadForumthreadidLazyRoute,
    AuthStatsPlayerSteamidLazyRoute,
    AuthStatsWeaponWeaponidLazyRoute,
  ]),
  GuestRoute.addChildren([
    GuestStvRoute,
    GuestContestsLazyRoute,
    GuestPrivacyPolicyLazyRoute,
    GuestServersLazyRoute,
    GuestIndexRoute,
    GuestLoginSuccessRoute,
    GuestProfileSteamIdLazyRoute,
    GuestWikiSlugLazyRoute,
    GuestLoginIndexRoute,
    GuestWikiIndexLazyRoute,
  ]),
  ModRoute.addChildren([
    ModAdminAppealsRoute,
    ModAdminContestsRoute,
    ModAdminFiltersRoute,
    ModAdminPeopleRoute,
    ModAdminReportsRoute,
    ModAdminVotesRoute,
    ModAdminNewsLazyRoute,
    ModAdminBanAsnRoute,
    ModAdminBanCidrRoute,
    ModAdminBanGroupRoute,
    ModAdminBanSteamRoute,
    ModAdminNetworkCidrblocksRoute,
    ModAdminNetworkIpInfoRoute,
    ModAdminNetworkIphistRoute,
    ModAdminNetworkPlayersbyipRoute,
    ModAdminNetworkIndexRoute,
  ]),
])

/* prettier-ignore-end */
