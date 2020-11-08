import React, { ReactNode } from "react";
import { Route, Redirect } from "react-router-dom";
import { currentUser } from "../../api/entity/user";

export interface Props {
  children: ReactNode;
}

const PrivateRoute: React.FC<Props> = ({ children, ...rest }: Props) => (
  <Route
    {...rest}
    render={(props) => {
      let user;

      currentUser()
        .then((userResponse) => (user = userResponse))
        .catch((error) => console.log(error));

      if (!user) {
        return (
          <Redirect
            to={{ pathname: "/auth/signin", state: { from: props.location } }}
          />
        );
      }

      return children;
    }}
  />
);

export default PrivateRoute;
