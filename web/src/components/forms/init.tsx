/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { initInstances } from "../../api/api";

interface Props {}

const Up: React.FC<Props> = () => {
  const [requestCount, setRequestCount] = useState<number>(0);
  const [instanceCount, setInstanceCount] = useState<number>(0);
  const [instanceSize, setInstanceSize] = useState<string>("");
  const [maxWorkingPeriod, setMaxWorkingPeriod] = useState<number>(0);
  const [region, setRegion] = useState<string>("AMS3");

  const handleChange = (name?: string) => (e: any) => {
    switch (name) {
      case "instanceCount":
        setInstanceCount(parseInt(e.target.value));
        break;
      case "requestCount":
        setRequestCount(parseInt(e.target.value));
        break;
      case "region":
        setRegion(e.target.value);
        break;
      case "maxWorkingPeriod":
        setMaxWorkingPeriod(parseInt(e.target.value));
        break;
    }
  };

  const sendRequest = (e: any) => {
    e.preventDefault();
    const instances = { requestCount, instanceCount, region, maxWorkingPeriod };
    initInstances(instances).then((data) => console.log(data));
  };

  return (
    <div css={formDiv}>
      <h2 css={formTitle}>Set up Testing Infrastructure</h2>
      <TextInput
        label={"Request Count"}
        type="text"
        name="requestCount"
        onChange={handleChange("requestCount")}
        value={requestCount}
      />
      <TextInput
        label={"Instance Count"}
        type="text"
        name="instanceCount"
        onChange={handleChange("instanceCount")}
        value={instanceCount}
      />
      <TextInput
        label={"Instance Size"}
        type="text"
        name="instanceSize"
        onChange={handleChange("instanceSize")}
        value={instanceCount}
      />
      <TextInput
        label={"Region"}
        type="text"
        name="region"
        value={region}
        onChange={handleChange("region")}
      />
      <TextInput
        label={"Max working period(minutes)"}
        type="text"
        name="maxWorkingPeriod"
        value={maxWorkingPeriod}
        onChange={handleChange("maxWorkingPeriod")}
      />
      <Button text="Up" onClick={sendRequest} />
    </div>
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

export default Up;
