/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import { runTests, Test, TestConfig } from "../../../api/entity/test_config";

interface Props {
  testConfg?: TestConfig;
}

const Show: React.FC<Props> = (props: Props) => {
  const run = (e: React.FormEvent) => {
    e.preventDefault();

    runTests(props.testConfg!)
      .then(() => {})
      .catch(() => {});
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        configName
        <hr />
        <Link to="/tests/create"> New Test Group</Link>
      </div>
      <div css={rightColumn}>
        {props.testConfg?.Tests.map((test: Test) => {
          return (
            <div css={configCss} key={test.url}>
              URL : {test.url} - Method: {test.method} - Request Count:{" "}
              {test.requestCount}
            </div>
          );
        })}
      </div>
    </div>
  );
};

const container = css`
  display: flex;
  width: 100%;
  flex-direction: row;
`;

const leftColumn = css`
  background-color: #e3e3e3;
  width: 30%;
  padding: 2rem;
`;

const rightColumn = css`
  width: 70%;
  padding: 2rem;
`;

const configCss = css`
  width: 100%;
  height: 5rem;
  padding: 2rem 0;
  border-bottom: 1px solid black;
  text-align: left;
`;

export default Show;
