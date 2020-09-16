import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import SpinUp from "../components/forms/SpinUpForm";
import RunWorkers from "../components/forms/RunWorkers";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  const [showRunWorkerForm, setShowRunWorkerForm] = React.useState<boolean>();

  const listInstances = () => {};

  const spinUpAfterHandle = () => {
    setShowRunWorkerForm(true);
  };

  // spinUpForm
  const spinUpForm: React.ReactNode = () => (
    <SpinUp afterHandle={spinUpAfterHandle} />
  );

  const runWorkersForm = () => <RunWorkers />;

  return (
    <React.Fragment>
      <MainLayout content={showRunWorkerForm ? runWorkersForm : spinUpForm} />
    </React.Fragment>
  );
};

export default Instances;
