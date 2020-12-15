/** @jsx jsx */
import React, { useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import SpinUp from "./SpinUpForm";
import {
  destroyAll,
  getInstanceInfoFromTerraform,
  InstanceTerra,
} from "../../../api/entity/instance";
import {
  card,
  cardTitle,
  cardContainer,
  cardItem,
  cardContent,
} from "../../style";
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
                <div css={cardContent}>
                  <div css={cardItem}>
                    <b>Region:</b>
                    <span> {instance.region}</span>
                  </div>
                  <div css={cardItem}>
                    <b>Size:</b>
                    <span> {instance.size}</span>
                  </div>
                  <div css={cardItem}>
                    <b>Memory:</b>
                    <span> {instance.memory}</span>
                  </div>
                  <div css={cardItem}>
                    <b>Disk: </b>
                    <span> {instance.disk}GB</span>
                  </div>
                  <div css={cardItem}>
                    <b>Image:</b>
                    <span> {instance.image}</span>
                  </div>
                  <div css={cardItem}>
                    <b>IPv4:</b>
                    <span> {instance.ipv4_address}</span>
                  </div>
                  <div css={cardItem}>
                    <b>Created:</b>
                    <span> {instance.created_at}</span>
                  </div>
                </div>
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
