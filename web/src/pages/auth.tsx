import React, { useEffect, useState } from "react";
import AuthContent from "../components/contents/AuthContent";
import MainLayout from "../components/layouts/MainLayout";
import { useHistory, useParams } from "react-router-dom";
import { getSettings } from "../api/entity/settings";
import { ServerResponse } from "../api/api";

const Auth: React.FC = () => {
  const [signupDisable, setSignupDisable] = useState<boolean>(false);
  const { type }: any = useParams();
  const history = useHistory();

  useEffect(() => {
    getSettings("signup")
      .then((response: ServerResponse) => {
        if (response.data.value === "Forbidden") {
          if (type === "signup") {
            history.push("/signin");
          } else {
            setSignupDisable(true);
          }
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <React.Fragment>
      <MainLayout
        content={<AuthContent type={type} signupDisable={signupDisable} />}
      />
    </React.Fragment>
  );
};

export default Auth;
