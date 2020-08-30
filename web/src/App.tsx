import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Instances from "./pages/instaces";
import Workers from "./pages/instaces";
import Stats from "./pages/stats";

const App: React.FC = () => {
  return (
    <React.Fragment>
      <Switch>
        <Route exact path="/instances">
          <Instances />
        </Route>
        <Route path="/workers">
          <Workers />
        </Route>
        <Route path="/stats">
          <Stats />
        </Route>
      </Switch>
    </React.Fragment>
  );
};

export default App;
