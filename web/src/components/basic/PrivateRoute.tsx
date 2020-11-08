import React, { ReactNode } from "react";
import { Route, Redirect } from "react-router-dom";

export interface Props {
  children: ReactNode;
}

const PrivateRoute: React.FC<Props> = ({ children, ...rest }: Props) => (
  <Route
    {...rest}
    render={(props) => {
      const user = "";

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
