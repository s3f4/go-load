import React, { useEffect } from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Instances from "./pages/instaces";
import Workers from "./pages/workers";
import Stats from "./pages/stats";
import Swarm from "./pages/swarm";
import Tests from "./pages/tests";
import Auth from "./pages/auth";
import NotFound from "./pages/not_found";
import PrivateRoute from "./components/basic/PrivateRoute";
import { currentUser } from "./api/entity/user";

const App: React.FC = () => {
  currentUser()
    .then((response) => console.log(response.data))
    .catch((error) => console.log(error));
  
  return (
    <Router>
      <Switch>
        <Route exact path="/auth/:type">
          <Auth />
        </Route>
        <PrivateRoute>
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
          <Route exact path="/">
            <Instances />
          </Route>
        </PrivateRoute>
        <Route component={NotFound} />
      </Switch>
    </Router>
  );
};

export default App;
