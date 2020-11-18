import React, { useEffect, useState } from "react";
import { Redirect, Route, Switch, useHistory } from "react-router-dom";
import { getInstanceInfo, InstanceConfig } from "../../api/entity/instance";
import Create from "./tests/create";
import Show from "./tests/show";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  const history = useHistory();

  useEffect(() => {
    getInstanceInfo()
      .then((response) => {
        const conf = response.data;
        if (!conf.data && !conf.data.configs.length) {
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
