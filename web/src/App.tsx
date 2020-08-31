import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Instances from "./pages/instaces";
import Workers from "./pages/instaces";
import Stats from "./pages/stats";
import Index from "./pages";

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Index />
        </Route>
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
    </Router>
  );
};

export default App;
