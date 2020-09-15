import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import SpinUp from "../components/forms/SpinUpForm";
import RunWorkers from "../components/forms/RunWorkers";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  const spinUpAfterHandle = () => {
    console.log("handled");
  };

  // spinUpForm
  const spinUpForm: React.ReactNode = () => (
    <SpinUp afterHandle={spinUpAfterHandle} />
  );

  const runWorkersForm = () => <RunWorkers />;

  return (
    <React.Fragment>
      <MainLayout content={spinUpForm} />
    </React.Fragment>
  );
};

export default Instances;
