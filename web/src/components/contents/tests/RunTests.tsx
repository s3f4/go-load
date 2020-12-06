/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Test } from "../../../api/entity/test";

interface Props {
  tests?: Test[];
}

const RunTests: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      {props.tests?.map((test: Test) => {
        return <div css={testLine}>{test.name}</div>;
      })}
    </div>
  );
};

const container = css`
  display: flex;
  flex-direction: column;
  margin-bottom: 2rem;
`;

const testLine = css`
  display: flex;
  width: 100%;
  margin: 0.2rem auto;
  background-color: #e3e3e3;
  border: 1px solid black;
  min-height: 3rem;
  padding: 1rem;
`;

export default RunTests;
