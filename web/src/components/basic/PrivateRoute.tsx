import React, { ReactNode, useEffect, useState } from "react";
import { Route, Redirect } from "react-router-dom";
import { getToken, token } from "../../api/entity/jwt";
import { getUserFromStorage } from "../../api/entity/user";
import Loader from "./Loader";
export interface Props {
  children: ReactNode;
}

const PrivateRoute: React.FC<Props> = ({ children, ...rest }: Props) => {
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    loginCheck();
  }, []);

  const loginCheck = () => {
    getToken()
      .then(() => {
        setLoading(false);
      })
      .catch(() => setLoading(false));
  };

  return (
    <Route
      {...rest}
      render={(props) => {
        const user = getUserFromStorage();
        if (loading) {
          return <Loader />;
        } else {
          if (!user || token === "") {
            return (
              <Redirect
                to={{
                  pathname: "/auth/signin",
                  state: { from: props.location },
                }}
              />
            );
          }

          return children;
        }
      }}
    />
  );
};

export default PrivateRoute;
