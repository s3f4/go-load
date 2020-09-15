/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";

interface Props {}

const RunWorkers = (props: Props) => {
  const [requestCount, setRequestCount] = React.useState<number>(0);
  const [goroutineCount, setGoroutineCount] = React.useState<number>(0);

  const run = (e: Event) => {
    e.preventDefault();
  };

  return (
    <div css={runWorkersForm}>
      <TextInput
        label="Total Request"
        name="requestCount"
        value={requestCount}
      />
      
      <TextInput
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
