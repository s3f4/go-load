import React from "react";
import SigninContent from "../components/contents/SigninContent";
import MainLayout from "../components/layouts/MainLayout";

const Signin: React.FC = () => {
  return (
    <React.Fragment>
      <MainLayout content={SigninContent} />
    </React.Fragment>
  );
};

export default Signin;
