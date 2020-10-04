import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import TestContent from "../components/contents/TestContent";

const Test: React.FC = () => {
  return (
    <React.Fragment>
      <MainLayout content={TestContent} />
    </React.Fragment>
  );
};

export default Test;
