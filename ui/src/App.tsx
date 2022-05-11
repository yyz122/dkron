import * as React from "react";
import Dashboard from './dashboard';
import { Admin, CustomRoutes, Resource } from 'react-admin';
import dataProvider from './dataProvider';
import jobs from './jobs';
import { BusyList } from './executions/BusyList';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';
import { Layout } from './layout';
import { createHashHistory } from "history";
import { Route } from 'react-router-dom';
import Configuration from './configuration/Configuration';

declare global {
    interface Window {
        DKRON_API_URL: string;
        DKRON_LEADER: string;
        DKRON_UNTRIGGERED_JOBS: string;
        DKRON_FAILED_JOBS: string;
        DKRON_SUCCESSFUL_JOBS: string;
        DKRON_TOTAL_JOBS: string;
    }
}

const authProvider = () => Promise.resolve();
const history = createHashHistory();

const App = () => (
    <Admin
        dashboard={Dashboard} 
        authProvider={authProvider}
        dataProvider={dataProvider}
        history={history}
        layout={Layout}
    >
        <CustomRoutes>
            <Route path="/configuration" element={<Configuration />}/>
        </CustomRoutes>
        <Resource name="jobs" {...jobs} />
        <Resource name="busy" options={{ label: 'Busy' }} list={BusyList} icon={PlayCircleOutlineIcon} />
        <Resource name="executions" />
        <Resource name="members" />
    </Admin>
);

export default App;
