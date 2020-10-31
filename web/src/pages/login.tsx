import React from "react";
import LoginContent from "../components/contents/LoginContent";
import MainLayout from "../components/layouts/MainLayout";

const Login: React.FC = () => {
  return (
    <React.Fragment>
      <MainLayout content={LoginContent} />
    </React.Fragment>
  );
};

export default Login;
