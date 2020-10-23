import React from "react";
import StatsContent from "../components/contents/StatsContent";
import MainLayout from "../components/layouts/MainLayout";
import { useParams } from "react-router-dom";

const Stats: React.FC = () => {
  const { id }: any = useParams();
  return (
    <React.Fragment>
      <MainLayout content={<StatsContent testID={id} />} />
    </React.Fragment>
  );
};

export default Stats;
