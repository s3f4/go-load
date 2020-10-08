import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Instances from "./pages/instaces";
import Workers from "./pages/workers";
import Stats from "./pages/stats";
import Swarm from "./pages/swarm";
import Index from "./pages";
import Tests from "./pages/tests";

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/instances">
          <Instances />
        </Route>
        <Route path="/tests">
          <Tests />
        </Route>
        <Route exact path="/workers">
          <Workers />
        </Route>
        <Route exact path="/stats">
          <Stats />
        </Route>
        <Route exact path="/swarm">
          <Swarm />
        </Route>
        <Route exact path="/">
          <Index />
        </Route>
      </Switch>
    </Router>
  );
};

export default App;
