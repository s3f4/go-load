import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import TestContent from "../components/contents/tests/TestContent";

const Tests: React.FC = () => {
  return (
    <React.Fragment>
      <MainLayout content={TestContent} />
    </React.Fragment>
  );
};

export default Tests;
