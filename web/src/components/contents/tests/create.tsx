/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Test, TestConfig } from "../../../api/entity/test_config";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import CreateTest from "../../forms/tests/CreateTest";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [configName, setConfigName] = React.useState<string>("");
  const [testConfig, setTestConfig] = React.useState<TestConfig>({
    Name: "",
    Tests: [],
  });

  const setConfig = (e: React.FormEvent) => {
    e.preventDefault();
    setTestConfig({
      ...testConfig,
      Name: configName,
    });
  };
  const addNewTest = (test: Test) => (e: React.FormEvent) => {
    e.preventDefault();
    setTestConfig({
      ...testConfig,
      Tests: [...testConfig.Tests, test],
    });
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setConfigName(e.target.value);
  };
  return (
    <div css={container}>
      <div css={leftColumn}>
        <TextInput
          name={"Test Config Name"}
          label={"Test Config Name"}
          onChange={handleChange}
        />
        <Button text="Create" onClick={setConfig} />
        <hr />
      </div>
      <div css={rightColumn}>
        {testConfig?.Tests.map((test: Test) => {
          return (
            <div css={configCss} key={test.url}>
              URL : {test.url} - Method: {test.method} - Request Count:{" "}
              {test.requestCount}
            </div>
          );
        })}
        <CreateTest addNewTest={addNewTest} />
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

export default Create;
