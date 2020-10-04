/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TestForm from "../forms/TestForm";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      <TestForm instanceInfo={null} />
    </div>
  );
};

const container = css``;

export default TestContent;
