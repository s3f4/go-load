/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { toNum } from "../basic/helper";
import { BaseForm } from "./BaseForm";
import { runWorkers } from "../../api/entity/worker";

interface Props extends BaseForm {}

const RunWorkers = (props: Props) => {
  const [requestCount, setRequestCount] = React.useState<number>(0);
  const [url, setUrl] = React.useState<string>("");
  const [goroutineCount, setGoroutineCount] = React.useState<number>(0);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    switch (e.target.name) {
      case "url":
        setUrl(e.target.value);
        break;
      case "requestCount":
        setRequestCount(toNum(e.target.value));
        break;
      case "goroutineCount":
        setGoroutineCount(toNum(e.target.value));
        break;
    }
  };

  const run = (e: React.FormEvent) => {
    e.preventDefault();
    runWorkers({
      requestCount,
      goroutineCount,
      url,
    })
      .then((response) => console.log(response))
      .catch((error) => console.log(error));
  };

  return (
    <div css={formDiv}>
      <h2 css={formTitle}>Run Workers</h2>
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
      <Button text="run" onClick={run} />
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
export default RunWorkers;
