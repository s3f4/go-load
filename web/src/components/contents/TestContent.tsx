/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  return <div css={container}>Create Test</div>;
};

const container = css``;

export default TestContent;
