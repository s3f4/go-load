/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { destroy, initInstances, listAvailableRegions } from "../../api/api";
import SelectBox from "../basic/SelectBox";
import { toNum } from "../basic/helper";
import Loader from "../basic/Loader";

interface Props {
  afterHandle?: () => void;
}

const SpinUp: React.FC<Props> = (props: Props) => {
  const [instanceCount, setInstanceCount] = useState<number>(0);
  const [maxWorkingPeriod, setMaxWorkingPeriod] = useState<number>(0);
  const [region, setRegion] = useState<string>("");
  const [regions, setRegions] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(false);

  React.useEffect(() => {
    regionsRequest();
  }, []);

  const handleChange = (name?: string) => (e: any) => {
    switch (name) {
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
    initInstances(instances).then((data) => {
      setLoading(false);
      console.log(data);
      props.afterHandle?.();
    });
  };

  const destroyRequest = (e: any) => {
    e.preventDefault();
    destroy().then((data) => console.log(data));
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
      .catch((err) => {
        console.log(err);
        setLoading(false);
      });
  };

  const formContent = () => {
    return (
      <div css={formDiv}>
        <h2 css={formTitle}>Set up Testing Infrastructure</h2>
        <TextInput
          label={"Instance Count"}
          type="text"
          name="instanceCount"
          onChange={handleChange("instanceCount")}
          value={instanceCount}
        />

        <TextInput
          label={"Max working period(minutes)"}
          type="text"
          name="maxWorkingPeriod"
          value={maxWorkingPeriod}
          onChange={handleChange("maxWorkingPeriod")}
        />

        <SelectBox
          name={"regions"}
          label={"Pick the region"}
          onChange={handleChange("regions")}
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
