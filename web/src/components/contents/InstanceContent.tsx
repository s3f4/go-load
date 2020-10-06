/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "../forms/SpinUpForm";
import TestForm from "../forms/TestForm";
import { useHistory } from "react-router-dom";
import { getInstanceInfo, InstanceConfig } from "../../api/entity/instance";
interface Props {}

const InstanceContent: React.FC<Props> = (props: Props) => {
  const [showRunWorkerForm, setShowRunWorkerForm] = React.useState<boolean>();
  const [instanceInfo, setInstanceInfo] = React.useState<InstanceConfig | null>(
    null,
  );

  const history = useHistory();

  React.useEffect(() => {
    getInstanceInfo()
      .then((response) => {
        setInstanceInfo(response.data);
      })
      .catch(() => {});
  }, []);

  const routeToStats = () => {
    history.push("/stats");
  };

  const spinUpAfterHandle = () => {
    setShowRunWorkerForm(true);
  };

  // spinUpForm
  const spinUpForm: React.ReactNode = (
    <SpinUp afterSubmit={spinUpAfterHandle} />
  );

  // runWorkersForm
  const runWorkersForm: React.ReactNode = (
    <TestForm instanceInfo={instanceInfo} afterSubmit={routeToStats} />
  );

  const content = () => {
    if (instanceInfo) {
      return runWorkersForm;
    } else {
      return showRunWorkerForm ? runWorkersForm : spinUpForm;
    }
  };

  return <div css={instanceContainer}>{content()}</div>;
};

const instanceContainer = css``;

export default InstanceContent;
