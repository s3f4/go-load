import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Instances from "./pages/instaces";
import Workers from "./pages/workers";
import Stats from "./pages/stats";
import Swarm from "./pages/swarm";
import Tests from "./pages/tests";
import Login from "./pages/signin";
import NotFound from "./pages/not_found";

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
        <Route exact path="/stats/:id">
          <Stats />
        </Route>
        <Route exact path="/swarm">
          <Swarm />
        </Route>
        <Route exact path="/login">
          <Login />
        </Route>
        <Route exact path="/">
          <Instances />
        </Route>
        <Route component={NotFound} />
      </Switch>
    </Router>
  );
};

export default App;
