import React from "react";
import { Route, Switch } from "react-router-dom";
import CreateTest from "./CreateTest";
import ShowTests from "./ShowTests";

const TestContent: React.FC = () => {
  return (
    <Switch>
      <Route exact path={"/tests/create"}>
        <CreateTest />
      </Route>
      <Route exact path={"/tests"}>
        <ShowTests />
      </Route>
    </Switch>
  );
};

export default TestContent;
