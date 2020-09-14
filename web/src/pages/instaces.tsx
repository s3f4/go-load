import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import Up from "../components/forms/init";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={<Up />} />
    </React.Fragment>
  );
};

export default Instances;
