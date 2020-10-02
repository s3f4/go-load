/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { toNum } from "../basic/helper";
import { BaseForm } from "./BaseForm";
import { runWorkers } from "../../api/entity/worker";
import Loader from "../basic/Loader";
import { RunConfig, TransportConfig } from "../../api/entity/run_config";
import SelectBox from "../basic/SelectBox";
import { destroyAll, InstanceConfig } from "../../api/entity/instance";

interface Props extends BaseForm {
  instanceInfo: InstanceConfig | null;
}

const RunWorkers = (props: Props) => {
  const [requestCount, setRequestCount] = React.useState<number>(0);
  const [url, setUrl] = React.useState<string>("");
  const [goroutineCount, setGoroutineCount] = React.useState<number>(1);
  const [loading, setLoading] = React.useState<boolean>(false);
  const [transportConfig, setTransportConfig] = React.useState<TransportConfig>(
    {
      DisableKeepAlives: true,
    },
  );

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    if (!e.target && e.hasOwnProperty("value") && e.hasOwnProperty("label")) {
      setTransportConfig({
        ...transportConfig,
        DisableKeepAlives: e.value === "true",
      });
      return;
    }

    switch (e.target.name) {
      case "url":
        setUrl(e.target.value);
        break;
      case "requestCount":
        setRequestCount(toNum(e.target.value));
        break;
      case "goroutineCount":
        let val = toNum(e.target.value);
        if (val <= 0) val = 1;
        setGoroutineCount(val);
        break;
    }
  };

  const run = (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    const runConfig: RunConfig = {
      requestCount,
      goroutineCount,
      url,
      transportConfig,
    };

    runWorkers(runConfig)
      .then(() => {
        setLoading(false);
        props.afterSubmit?.();
      })
      .catch(() => {
        setLoading(false);
        props.afterSubmit?.();
      });
  };

  const destroyRequest = (e: any) => {
    e.preventDefault();
    destroyAll()
      .then((data) => console.log(data))
      .catch((error) => console.log(error));
  };

  const formContent = () => {
    return (
      <div css={formDiv}>
        <h2 css={formTitle}>Run Workers</h2>
        InstanceInfo: {JSON.stringify(props.instanceInfo)}
        <Button text="Destroy" onClick={destroyRequest} />
        <TextInput
          onChange={handleChange}
          label="Target URL"
          name="url"
          value={url}
        />
        <TextInput
          onChange={handleChange}
          label="Total Request"
          name="requestCount"
          value={requestCount}
        />
        <TextInput
          onChange={handleChange}
          label="Goroutine per worker (up to 10)"
          name="goroutineCount"
          value={goroutineCount}
        />
        <SelectBox
          name={"disableKeepAlives"}
          label={"Disable Keep-alive connections"}
          onChange={handleChange}
          options={[
            { value: "true", label: "True" },
            { value: "false", label: "False" },
          ]}
          value={transportConfig.DisableKeepAlives ? "true" : "false"}
        />
        <Button text="run" onClick={run} />
      </div>
    );
  };

  return loading ? <Loader message="Workers are running..." /> : formContent();
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
export default RunWorkers;
