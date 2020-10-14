/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import SelectBox from "../../basic/SelectBox";
import { toNum } from "../../basic/helper";
import Loader from "../../basic/Loader";
import { BaseForm } from "../../basic/BaseForm";
import {
  spinUp,
  listAvailableRegions,
  InstanceConfig,
  showAccount,
} from "../../../api/entity/instance";
import InstanceConfigCards from "./InstanceConfigCards";

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

  React.useEffect(() => {
    let mount = true;
    regionsRequest(mount);
    accountRequest(mount);
    return () => {
      mount = false;
    };
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

  const regionsRequest = (isMounted: boolean) => {
    setRegionsLoading(true);
    listAvailableRegions()
      .then((response) => {
        if (isMounted) {
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
        }
      })
      .catch((error) => {
        setRegionsLoading(false);
        console.log(error);
      });
  };

  const accountRequest = (isMounted: boolean) => {
    showAccount()
      .then((response) => {
        if (isMounted) {
          const data = JSON.parse(response.data);
          setInstanceLimit(data.droplet_limit - 2);
        }
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const add = (config: any) => (e: React.FormEvent) => {
    e.preventDefault();
    const newConfigs = configs.map((conf) => {
      if (conf.region === config.region) {
        conf.instanceCount++;
      }
      return conf;
    });
    setConfigs([...newConfigs]);
  };

  const remove = (config: any) => (e: React.FormEvent) => {
    e.preventDefault();
    const newConfigs: any[] = [];
    configs.forEach((conf) => {
      if (conf.region === config.region) {
        if (conf.instanceCount === 1) {
        } else {
          conf.instanceCount--;
          newConfigs.push(conf);
        }
      } else {
        newConfigs.push(conf);
      }
    });
    setConfigs([...newConfigs]);
  };

  const totalInstanceCount = (): number => {
    let count = 0;
    configs.forEach((config) => (count += config.instanceCount));
    return count;
  };

  const formContent = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h2 css={formTitle}>Set up Testing Infrastructure</h2>
          <div css={title}>
            Your droplet limit is <b>{instanceLimit + 2}</b> and already used{" "}
            <b>2</b>, you can increase this on digitalocean.
          </div>
          <TextInput
            name="instanceCount"
            label={"Instance Count"}
            type="text"
            onChange={handleChange}
            value={instanceCount}
            validate={{
              min: 1,
              max: instanceLimit,
              message: "Your can create " + instanceLimit + " instances.",
              validationFunction: validation("instanceCount"),
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
              validationFunction: validation("region"),
            }}
            isValid={isValid["region"]}
          />

          <div css={buttons}>
            <Button
              loading={regionsLoading}
              text="Add New Instance"
              onClick={addNewInstance}
              disabled={
                !isValid["instanceCount"] ||
                !isValid["region"] ||
                totalInstanceCount() + instanceCount > instanceLimit
              }
            />

            <Button
              text="Spin Up"
              onClick={sendRequest}
              disabled={
                !isValid["instanceCount"] ||
                !isValid["region"] ||
                configs.length === 0 ||
                totalInstanceCount() > instanceLimit
              }
            />
          </div>
        </div>
        <InstanceConfigCards configs={configs} add={add} remove={remove} />
      </div>
    );
  };

  return loading ? (
    <Loader message="Instances will be created in a few minutes.." />
  ) : (
    formContent()
  );
};

const title = css`
  width: 100%;
  height: 4rem;
  text-align: center;
  margin: 1rem auto;
  padding: 1rem;
  background-color: #efefef;
`;

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

const buttons = css`
  margin-top: 1rem;
  height: 6.5rem;
`;

export default SpinUp;
