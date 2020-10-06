/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import SelectBox from "../basic/SelectBox";
import { toNum } from "../basic/helper";
import Loader from "../basic/Loader";
import { BaseForm } from "./BaseForm";
import {
  spinUp,
  listAvailableRegions,
  InstanceConfig,
  showAccount,
} from "../../api/entity/instance";
import { Box, Sizes } from "../style";

interface Props extends BaseForm {}

const SpinUp: React.FC<Props> = (props: Props) => {
  const [instanceCount, setInstanceCount] = useState<number>(1);
  const [region, setRegion] = useState<string>("");
  const [regions, setRegions] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [regionsLoading, setRegionsLoading] = useState<boolean>(false);
  const [configs, setConfigs] = useState<any[]>([]);
  const [instanceLimit, setInstanceLimit] = useState<number>(0);
  const [isValid, setIsValid] = useState<any>({
    instanceCount: true,
    region: false,
  });

  console.log(isValid);

  React.useEffect(() => {
    regionsRequest();
    accountRequest();
  }, []);

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    if (!e.target && e.hasOwnProperty("value") && e.hasOwnProperty("label")) {
      setRegion(e.value);
      return;
    }

    switch (e.target.name) {
      case "instanceCount":
        setInstanceCount(toNum(e.target.value));
        break;
    }
  };

  const sendRequest = (e: any) => {
    e.preventDefault();
    setLoading(true);
    const instanceConfig: InstanceConfig = {
      Configs: configs,
    };

    spinUp(instanceConfig)
      .then(() => {
        setLoading(false);
        props.afterSubmit?.();
      })
      .catch((error) => {
        setLoading(false);
        console.log(error);
        props.afterSubmit?.();
      });
  };

  const addNewInstance = (e: React.FormEvent) => {
    e.preventDefault();
    const found = configs.find((config) => config.region === region);
    if (!found && region !== "") {
      setConfigs([
        ...configs,
        {
          instanceCount,
          region,
        },
      ]);
    }
  };

  const regionsRequest = () => {
    setRegionsLoading(true);
    listAvailableRegions()
      .then((response) => {
        if (response && response.status) {
          const jsonRes = JSON.parse(response.data);
          const regions = jsonRes.regions;
          const regionSelectBox = regions.map((region: any) => {
            return {
              label: region.name,
              value: region.slug,
            };
          });
          setRegionsLoading(false);
          setRegions(regionSelectBox);
        }
      })
      .catch((error) => {
        setRegionsLoading(false);
        console.log(error);
      });
  };

  const accountRequest = () => {
    showAccount()
      .then((response) => {
        const data = JSON.parse(response.data);
        setInstanceLimit(data.droplet_limit);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const formContent = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h2 css={formTitle}>Set up Testing Infrastructure</h2>
          <TextInput
            name="instanceCount"
            label={"Instance Count"}
            type="text"
            onChange={handleChange}
            value={instanceCount}
            validate={{
              min: 1,
              max: instanceLimit,
              message: "Your droplet limit is " + instanceLimit,
              isValid: validation("instanceCount"),
            }}
            isValid={isValid["instanceCount"]}
          />

          <SelectBox
            name={"region"}
            label={"Pick the region"}
            onChange={handleChange}
            options={regions}
            value={region}
            validate={{
              minLength: 3,
              message: "Please choose a region",
              isValid: validation("region"),
            }}
            isValid={isValid["region"]}
          />

          <div css={buttons}>
            <Button
              loading={regionsLoading}
              text="Add New Instance"
              onClick={addNewInstance}
              disabled={!isValid["instanceCount"] || !isValid["region"]}
            />

            <Button
              text="Spin Up"
              onClick={sendRequest}
              disabled={
                !isValid["instanceCount"] ||
                !isValid["region"] ||
                configs.length === 0
              }
            />
          </div>
        </div>
        <div css={configContainer}>
          {configs &&
            configs.map((config) => {
              return (
                <div css={configCss} key={config.region}>
                  <div css={instanceTitle}>Region: {config.region}</div>
                  <br />
                  Instance Count: <b>{config.instanceCount}</b>
                  <Button type={1} text="+" onClick={() => {}} />
                  <Button type={1} text="-" onClick={() => {}} />
                </div>
              );
            })}
        </div>
      </div>
    );
  };

  return loading ? (
    <Loader message="Instances will be created in a few minutes.." />
  ) : (
    formContent()
  );
};

const container = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const formDiv = css`
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  width: 60%;
  margin-bottom: 3rem;
`;

const formTitle = css`
  font-size: 2.3rem;
  text-decoration: none;
  text-align: center;
`;

const configContainer = css`
  width: 100%;
  display: flex;
  flex-wrap: wrap;
`;

const configCss = css`
  background-color: #efefef;
  width: 15rem;
  height: 15rem;
  margin: 1rem 1rem;
  border: 1px solid black;
  text-align: center;
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1}
`;

const instanceTitle = css`
  background-color: #007d9c;
  color: white;
  width: 100%;
  height: 100;
  padding: 0.5rem;
  font-weight: bold;
`;

const buttons = css`
  margin-top: 1rem;
  height: 6.5rem;
`;

export default SpinUp;
