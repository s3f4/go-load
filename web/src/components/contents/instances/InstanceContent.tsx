/** @jsx jsx */
import React, { useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "./SpinUpForm";
import {
  destroyAll,
  getInstanceInfoFromTerraform,
  InstanceTerra,
} from "../../../api/entity/instance";
import { card, cardTitle, cardContainer } from "../../style";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { FiX } from "react-icons/fi";

const InstanceContent: React.FC = () => {
  const [loading, setLoading] = useState<boolean>(true);
  const [instances, setInstances] = useState<InstanceTerra[]>();

  useEffect(() => {
    let mount = true;
    onGetInstances(mount);
    return () => {
      mount = false;
    };
  }, []);

  const onGetInstances = (mount?: boolean) => {
    setLoading(true);
    getInstanceInfoFromTerraform()
      .then((response) => {
        if (mount) {
          let arr: any = [];
          const obj = JSON.parse(response.data);
          Object.keys(obj).map(function (key) {
            arr.push(obj[key]);
            return arr;
          });
          setInstances(arr);
          setLoading(false);
        }
      })
      .catch(() => {
        setLoading(false);
      });
  };

  const spinUpAfterHandle = () => {
    onGetInstances(true);
  };

  // spinUpForm
  const spinUpForm = () => {
    if (loading) {
      return "";
    }

    return <SpinUp afterSubmit={spinUpAfterHandle} />;
  };

  const onDestroyAll = () => {
    setLoading(true);
    destroyAll()
      .then(() => {
        onGetInstances(true);
        setLoading(false);
      })
      .catch(() => {
        onGetInstances(true);
        setLoading(false);
      });
  };

  const instanceList = () => (
    <div>
      <div css={center}>
        <Button
          type={ButtonType.iconTextButton}
          colorType={ButtonColorType.danger}
          icon={<FiX />}
          loading={loading}
          text="Destroy All"
          onClick={onDestroyAll}
        />
      </div>
      <div css={cardContainer}>
        {instances &&
          instances.length > 0 &&
          instances.map((instance: InstanceTerra) => {
            return (
              <div css={card} key={instance.region}>
                <h1 css={cardTitle}>{instance.name}</h1>
                Region: {instance.region} <br />
                Size: {instance.size} <br />
                Memory: {instance.memory} <br />
                Disk: {instance.disk}GB <br />
                Image: {instance.image} <br />
                Size: {instance.size} <br />
                IPv4: {instance.ipv4_address} <br />
                Created: {instance.created_at}
              </div>
            );
          })}
      </div>
    </div>
  );

  return (
    <React.Fragment>
      {instances && instances.length ? instanceList() : spinUpForm()}
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
