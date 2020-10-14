/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "./SpinUpForm";
import {
  getInstanceInfo,
  Instance,
  InstanceConfig,
} from "../../../api/entity/instance";
import { card, cardTitle, cardContainer } from "../../style";

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
  const runWorkersForm: React.ReactNode = (
    <div>
      <div css={cardContainer}>
        {instanceInfo &&
          instanceInfo.configs &&
          instanceInfo.configs.map((instance: Instance) => {
            return (
              <div css={card} key={instance.region}>
                <h1 css={cardTitle}>{instance.region}</h1>
                Size: {instance.instance_size} <br />
                <br />
                <br />
              </div>
            );
          })}
      </div>
    </div>
  );

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
