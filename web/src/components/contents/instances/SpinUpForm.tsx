/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import SelectBox from "../../basic/SelectBox";
import { toNum } from "../../basic/helper";
import Loader from "../../basic/Loader";
import { BaseForm, validateAll } from "../../basic/BaseForm";
import {
  spinUp,
  listAvailableRegions,
  InstanceConfig,
  showAccount,
  Instance,
} from "../../../api/entity/instance";
import InstanceConfigCards from "./InstanceConfigCards";
import { MediaQuery } from "../../style";

interface Props extends BaseForm {}

const SpinUp: React.FC<Props> = (props: Props) => {
  const [count, setCount] = useState<number>(1);
  const [region, setRegion] = useState<string>("");
  const [regions, setRegions] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [regionsLoading, setRegionsLoading] = useState<boolean>(false);
  const [configs, setConfigs] = useState<any[]>([]);
  const [instanceLimit, setInstanceLimit] = useState<number>(0);
  const [isValid, setIsValid] = useState<any>({
    count: true,
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
      case "count":
        setCount(toNum(e.target.value));
        break;
    }
  };

  const sendRequest = (e: any) => {
    e.preventDefault();
    setLoading(true);
    const instanceConfig: InstanceConfig = {
      configs: configs,
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
          count,
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
    const newConfigs = configs.map((conf: Instance) => {
      if (conf.region === config.region) {
        conf.count++;
      }
      return conf;
    });
    setConfigs([...newConfigs]);
  };

  const remove = (config: Instance) => (e: React.FormEvent) => {
    e.preventDefault();
    const newConfigs: Instance[] = [];
    configs.forEach((conf: Instance) => {
      if (conf.region === config.region) {
        if (conf.count === 1) {
        } else {
          conf.count--;
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
    configs.forEach((config: Instance) => (count += config.count));
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
            name="count"
            label={"Instance Count"}
            type="text"
            onChange={handleChange}
            value={count}
            validate={`min:1|max:${instanceLimit}|message:You can create ${instanceLimit} instances.`}
            validation={validation}
          />

          <SelectBox
            name={"region"}
            label={"Pick the region"}
            onChange={handleChange}
            options={regions}
            value={region}
            validate="minLength:3|message:Please choose a region"
            validation={validation}
          />

          <div css={buttons}>
            <Button
              loading={regionsLoading}
              text="Add New Instance"
              onClick={addNewInstance}
              disabled={
                validateAll(isValid) ||
                totalInstanceCount() + count > instanceLimit
              }
            />

            <Button
              text="Spin Up"
              onClick={sendRequest}
              disabled={
                validateAll(isValid) ||
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
  text-align: center;
  margin: 1rem auto;
  padding: 1rem;
  background-color: #efefef;

  ${MediaQuery[1]} {
    height: 4rem;
  }
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
  margin-bottom: 3rem;
  width: 95%;
  ${MediaQuery[1]} {
    width: 60%;
  }
`;

const formTitle = css`
  margin-top: 1rem;
  font-size: 2.3rem;
  text-decoration: none;
  text-align: center;
`;

const buttons = css`
  display: flex;
  justify-content: space-around;
  margin: 1rem auto;
  width: 60%;
  ${MediaQuery[1]} {
    width: 40%;
  }
`;

export default SpinUp;
