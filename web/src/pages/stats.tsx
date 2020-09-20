import React from "react";
import StatsContent from "../components/contents/StatsContent";
import MainLayout from "../components/layouts/MainLayout";

interface Props {}

const Stats: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={StatsContent} />
    </React.Fragment>
  );
};

export default Stats;
