import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import SpinUp from "../components/forms/SpinUpForm";
import RunWorkers from "../components/forms/RunWorkers";
import { useHistory } from "react-router-dom";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  const [showRunWorkerForm, setShowRunWorkerForm] = React.useState<boolean>();
  const history = useHistory();

  const routeToStats = () => {
    history.push("/stats");
  };

  const spinUpAfterHandle = () => {
    setShowRunWorkerForm(true);
  };

  // spinUpForm
  const spinUpForm: React.ReactNode = () => (
    <SpinUp afterHandle={spinUpAfterHandle} />
  );

  const runWorkersForm = () => <RunWorkers afterHandle={routeToStats} />;

  return (
    <React.Fragment>
      <MainLayout content={showRunWorkerForm ? runWorkersForm : spinUpForm} />
    </React.Fragment>
  );
};

export default Instances;
