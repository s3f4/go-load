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
  destroyAll,
  listAvailableRegions,
} from "../../api/entity/instance";

interface Props extends BaseForm {}

const SpinUp: React.FC<Props> = (props: Props) => {
  const [instanceCount, setInstanceCount] = useState<number>(0);
  const [maxWorkingPeriod, setMaxWorkingPeriod] = useState<number>(0);
  const [region, setRegion] = useState<string>("");
  const [regions, setRegions] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(false);

  React.useEffect(() => {
    regionsRequest();
  }, []);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    switch (e.target.name) {
      case "instanceCount":
        setInstanceCount(toNum(e.target.value));
        break;
      case "maxWorkingPeriod":
        setMaxWorkingPeriod(toNum(e.target.value));
        break;
      case "regions":
        setRegion(e.target.value);
        break;
    }
  };

  const sendRequest = (e: any) => {
    e.preventDefault();
    setLoading(true);
    const instances = { instanceCount, maxWorkingPeriod, region };
    spinUp(instances)
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

  const destroyRequest = (e: any) => {
    e.preventDefault();
    destroyAll().then((data) => console.log(data));
  };

  const regionsRequest = () => {
    listAvailableRegions()
      .then((response) => {
        if (response && response.status) {
          const jsonRes = JSON.parse(response.message);
          const regions = jsonRes.regions;
          const regionSelectBox = regions.map((region: any) => {
            return {
              text: region.name,
              value: region.slug,
            };
          });
          setRegions(regionSelectBox);
        }
      })
      .catch((error) => console.log(error));
  };

  const formContent = () => {
    return (
      <div css={formDiv}>
        <h2 css={formTitle}>Set up Testing Infrastructure</h2>
        <TextInput
          name="instanceCount"
          label={"Instance Count"}
          type="text"
          onChange={handleChange}
          value={instanceCount}
        />

        <TextInput
          name="maxWorkingPeriod"
          label={"Max working period(minutes)"}
          type="text"
          value={maxWorkingPeriod}
          onChange={handleChange}
        />

        <SelectBox
          name={"regions"}
          label={"Pick the region"}
          onChange={handleChange}
          options={regions}
          value={region}
        />

        <Button text="Spin Up" onClick={sendRequest} />
        <Button text="Destroy" onClick={destroyRequest} />
      </div>
    );
  };

  return loading ? (
    <Loader message="Instances will be created in a few minutes.." />
  ) : (
    formContent()
  );
};

const formDiv = css`
  margin: 0 auto;
  width: 50%;
`;

const formTitle = css`
  font-size: 2.3rem;
  text-decoration: none;
  text-align: center;
`;

export default SpinUp;
