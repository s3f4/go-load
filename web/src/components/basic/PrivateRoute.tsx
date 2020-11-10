import React, { ReactNode } from "react";
import { Route, Redirect } from "react-router-dom";
import { getUserFromStorage } from "../../api/entity/user";

export interface Props {
  children: ReactNode;
}

const PrivateRoute: React.FC<Props> = ({ children, ...rest }: Props) => (
  <Route
    {...rest}
    render={(props) => {
      const user = getUserFromStorage();

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
