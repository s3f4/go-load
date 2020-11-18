/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "./SpinUpForm";
import {
  destroyAll,
  getInstanceInfo,
  Instance,
  InstanceConfig,
} from "../../../api/entity/instance";
import { card, cardTitle, cardContainer } from "../../style";
import Button from "../../basic/Button";

const InstanceContent: React.FC = () => {
  const [loading, setLoading] = React.useState<boolean>(false);
  const [showInstances, setShowInstances] = React.useState<boolean>();
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
    setShowInstances(true);
  };

  // spinUpForm
  const spinUpForm = () => <SpinUp afterSubmit={spinUpAfterHandle} />;

  const onDestroyAll = () => {
    setLoading(true);
    destroyAll()
      .then((response) => {
        debugger;
        setLoading(false);
        console.log(response);
        setShowInstances(false);
      })
      .catch((error) => {
        debugger;
        setLoading(false);
        console.log(error);
        setShowInstances(false);
      });
  };

  const instanceList = () => (
    <div>
      <div css={center}>
        <Button loading={loading} text="Destroy All" onClick={onDestroyAll} />
      </div>
      <div css={cardContainer}>
        {instanceInfo &&
          instanceInfo.configs &&
          instanceInfo.configs.map((instance: Instance) => {
            return (
              <div css={card} key={instance.region}>
                <h1 css={cardTitle}>{instance.region}</h1>
                Size: {instance.size} <br />
              </div>
            );
          })}
      </div>
    </div>
  );

  return (
    <React.Fragment>
      {showInstances ? instanceList() : spinUpForm()}
    </React.Fragment>
  );
};

const center = css`
  height: 3em;
  display: flex;
  align-items: center;
  justify-content: center;
`;

export default InstanceContent;
