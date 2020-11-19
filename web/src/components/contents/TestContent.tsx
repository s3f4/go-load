import React from "react";
import { Route, Switch } from "react-router-dom";
import Create from "./tests/create";
import Show from "./tests/show";

const TestContent: React.FC = () => {
  return (
    <Switch>
      <Route exact path={"/tests/create"}>
        <Create />
      </Route>
      <Route exact path={"/tests"}>
        <Show />
      </Route>
    </Switch>
  );
};

export default TestContent;
