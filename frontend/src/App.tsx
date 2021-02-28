import React from "react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import {Home} from "./page/Home";
import {Settings} from "./page/Settings";
import {Appeal} from "./page/Appeal";
import {Report} from "./page/Report";
import {AdminReports} from "./page/AdminReports";
import {AdminFilters} from "./page/AdminFilters";
import {AdminImport} from "./page/AdminImport";
import {AdminPeople} from "./page/AdminPeople";
import {Bans} from "./page/Bans";
import {Servers} from "./page/Servers";
import {AdminServers} from "./page/AdminServers";
import {Header} from "./Header";
import {PlayerSummary} from "./component/PlayerBanForm";
import {Flash, Flashes} from "./component/Flashes";
import {LoginSuccess} from "./page/LoginSuccess";

// const onRedirectCallback = (appState) => {
//     history.push(
//         appState && appState.returnTo ? appState.returnTo : window.location.pathname
//     );
// };

export const App = () => {
    // @ts-ignore
    const [loggedIn, setLoggedIn] = React.useState<boolean>(false)
    // @ts-ignore
    const [currentProfile, setCurrentProfile] = React.useState<PlayerSummary>({
        personaname: "Guest",
        steam_id: 0,
    })
    // @ts-ignore
    const [flashes, setFlashes] = React.useState<Flash[]>([])
    const handleOnLogin = () => {
        const r = `${window.location.protocol}//${window.location.hostname}/auth/callback?return_url=${window.location.pathname}`
        const oid = "https://steamcommunity.com/openid/login" +
            "?openid.ns=" + encodeURIComponent("http://specs.openid.net/auth/2.0") +
            "&openid.mode=checkid_setup" +
            "&openid.return_to=" + encodeURIComponent(r) +
            `&openid.realm=` + encodeURIComponent(`${window.location.protocol}//${window.location.hostname}`) +
            "&openid.ns.sreg=" + encodeURIComponent("http://openid.net/extensions/sreg/1.1") +
            "&openid.claimed_id=" + encodeURIComponent("http://specs.openid.net/auth/2.0/identifier_select") +
            "&openid.identity=" + encodeURIComponent("http://specs.openid.net/auth/2.0/identifier_select")
        window.open(oid, "_self")
    }

    const handleOnLogout = () => {
        localStorage.clear();
    }

    return (
        <Router>
            <Header loggedIn={currentProfile.steam_id > 0} profile={currentProfile} onLogin={handleOnLogin}
                    onLogout={handleOnLogout}/>

            <Flashes flashes={flashes} />

            <Route path="/" component={Home} exact={true}/>
            <Route path="/servers" component={Servers}/>
            <Route path="/bans" component={Bans}/>
            <Route path="/appeal" component={Appeal}/>
            <Route path="/report" component={Report}/>
            <Route path="/settings" component={Settings}/>
            <Route path="/admin/filters" component={AdminFilters}/>
            <Route path="/admin/reports" component={AdminReports}/>
            <Route path="/admin/import" component={AdminImport}/>
            <Route path="/admin/people" component={AdminPeople}/>
            <Route path="/admin/servers" component={AdminServers}/>
            <Route path="/login/success" component={LoginSuccess}/>
            <footer className="grid-container full">
                <div className="grid-x grid-padding-x" id="footer">
                    <div className="cell">
                        <p>Powered By: <a href="https://github.com/leighmacdonald/gbans">gbans</a></p>
                    </div>
                </div>
            </footer>
        </Router>
    )
}