/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { toNum } from "../basic/helper";

interface Props {}

const RunWorkers = (props: Props) => {
  const [requestCount, setRequestCount] = React.useState<number>(0);
  const [goroutineCount, setGoroutineCount] = React.useState<number>(0);

  const handleChange = (e: any) => {
    switch (e.target.name) {
      case "requestCount":
        setRequestCount(toNum(e.target.value));
        break;
      case "goroutineCount":
        setGoroutineCount(toNum(e.target.value));
        break;
    }
  };

  const run = (e: Event) => {
    console.log(e);
  };

  return (
    <div css={runWorkersForm}>
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

const runWorkersForm = css``;

export default RunWorkers;
