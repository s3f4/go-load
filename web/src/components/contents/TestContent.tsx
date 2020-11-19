import React, { useEffect } from "react";
import { Route, Switch, useHistory } from "react-router-dom";
import { getInstanceInfo } from "../../api/entity/instance";
import Create from "./tests/create";
import Show from "./tests/show";

const TestContent: React.FC = () => {
  const history = useHistory();

  useEffect(() => {
    getInstanceInfo()
      .then((response) => {
        const conf = response.data;
        if (!conf && !conf.configs.length) {
          history.push("/instances");
        }
      })
      .catch(() => {
        history.push("/instances");
      });
    return () => {};
  }, []);

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
