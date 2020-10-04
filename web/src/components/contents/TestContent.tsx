/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import RunWorkers from "../forms/RunWorkers";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      <RunWorkers instanceInfo={null} />
    </div>
  );
};

const container = css``;

export default TestContent;
