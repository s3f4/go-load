import React, { useEffect, useState } from "react";
import { Redirect, Route, Switch } from "react-router-dom";
import { getInstanceInfo, InstanceConfig } from "../../api/entity/instance";
import Create from "./tests/create";
import Show from "./tests/show";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  const [instanceConfig, setInstanceConfig] = useState<InstanceConfig>();

  useEffect(() => {
    listInstances();
    return () => {};
  }, []);

  const listInstances = () => {
    getInstanceInfo()
      .then((response) => {
        setInstanceConfig(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  if (!instanceConfig?.configs) {
    return <Redirect to="/instances" />;
  }

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
