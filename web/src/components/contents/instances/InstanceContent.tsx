/** @jsx jsx */
import React, { useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "./SpinUpForm";
import {
  destroyAll,
  getInstanceInfo,
  getInstanceInfoFromTerraform,
  InstanceConfig,
  InstanceTerra,
} from "../../../api/entity/instance";
import { card, cardTitle, cardContainer } from "../../style";
import Button from "../../basic/Button";

const InstanceContent: React.FC = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [showInstances, setShowInstances] = useState<boolean>(false);
  const [instanceInfo, setInstanceInfo] = useState<InstanceConfig | null>(null);

  const [instanceTerra, setInstanceTerra] = useState<InstanceTerra[]>();

  useEffect(() => {
    let mount = true;
    onGetInstanceInfo(mount);
    onGetInstanceInfoFromTerraform(true);
    return () => {
      mount = false;
    };
  }, []);

  const onGetInstanceInfoFromTerraform = (mount?: boolean) => {
    getInstanceInfoFromTerraform()
      .then((response) => {
        if (mount) {
          let arr: any = [];
          const obj = JSON.parse(response.data);
          Object.keys(obj).map(function (key) {
            arr.push(obj[key]);
            return arr;
          });
          setInstanceTerra(arr);
        }
      })
      .catch(() => {});
  };

  const onGetInstanceInfo = (mount?: boolean) => {
    getInstanceInfo()
      .then((response) => {
        if (mount) {
          setInstanceInfo(response.data);
          setShowInstances(true);
        }
      })
      .catch(() => {});
  };

  const spinUpAfterHandle = () => {
    onGetInstanceInfo();
    setShowInstances(true);
  };

  // spinUpForm
  const spinUpForm = () => {
    return <SpinUp afterSubmit={spinUpAfterHandle} />;
  };

  const onDestroyAll = () => {
    setLoading(true);
    destroyAll()
      .then(() => {
        setLoading(false);
        setShowInstances(false);
      })
      .catch(() => {
        setLoading(false);
        setShowInstances(false);
      });
  };

  const instanceList = () => (
    <div>
      <div css={center}>
        <Button loading={loading} text="Destroy All" onClick={onDestroyAll} />
      </div>
      <div css={cardContainer}>
        {instanceTerra &&
          instanceTerra.length > 0 &&
          instanceTerra.map((instance: InstanceTerra) => {
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
