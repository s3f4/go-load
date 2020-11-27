import React from "react";
import StatsContent from "../components/contents/stats/StatsContent";
import MainLayout from "../components/layouts/MainLayout";
import { Switch, Route, useRouteMatch } from "react-router-dom";
import ListStats from "../components/contents/stats/ListStats";

const Stats: React.FC = () => {
  let { path } = useRouteMatch();

  return (
    <Switch>
      <Route exact path={`${path}`}>
        <MainLayout content={<ListStats />} />
      </Route>
      <Route exact path={`${path}/:id`}>
        <MainLayout content={<StatsContent />} />
      </Route>
    </Switch>
  );
};

export default Stats;
