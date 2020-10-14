import React from "react";
import SpinUp from "./SpinUpForm";
import { getInstanceInfo, InstanceConfig } from "../../../api/entity/instance";

const InstanceContent: React.FC = () => {
  const [showRunWorkerForm, setShowRunWorkerForm] = React.useState<boolean>();
  const [instanceInfo, setInstanceInfo] = React.useState<InstanceConfig | null>(
    null,
  );

  React.useEffect(() => {
    let mount = true;
    getInstanceInfo()
      .then((response) => {
        if (mount) {
          setInstanceInfo(response.data);
        }
      })
      .catch(() => {});
    return () => {
      mount = false;
    };
  }, []);

  const spinUpAfterHandle = () => {
    setShowRunWorkerForm(true);
  };

  // spinUpForm
  const spinUpForm: React.ReactNode = (
    <SpinUp afterSubmit={spinUpAfterHandle} />
  );

  // runWorkersForm
  const runWorkersForm: React.ReactNode = <div></div>;

  const content = () => {
    if (instanceInfo) {
      return runWorkersForm;
    } else {
      return showRunWorkerForm ? runWorkersForm : spinUpForm;
    }
  };

  return <React.Fragment>{content()}</React.Fragment>;
};

export default InstanceContent;
