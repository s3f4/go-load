import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import SwarmContent from "../components/contents/SwarmContent";

interface Props {}

const Swarm: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={SwarmContent} />
    </React.Fragment>
  );
};

export default Swarm;
