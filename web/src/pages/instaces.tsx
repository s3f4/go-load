import React from "react";
import InstanceContent from "../components/contents/instances/InstanceContent";
import MainLayout from "../components/layouts/MainLayout";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={InstanceContent} />
    </React.Fragment>
  );
};

export default Instances;
