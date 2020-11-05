import React from "react";
import AuthContent from "../components/contents/AuthContent";
import MainLayout from "../components/layouts/MainLayout";
import { useParams } from "react-router-dom";

const Auth: React.FC = () => {
  const { type }: any = useParams();
  return (
    <React.Fragment>
      <MainLayout content={<AuthContent type={type} />} />
    </React.Fragment>
  );
};

export default Auth;
