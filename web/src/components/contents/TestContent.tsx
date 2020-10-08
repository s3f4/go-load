/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TestForm from "../forms/TestForm";
import { destroyAll } from "../../api/entity/instance";
import Button from "../basic/Button";
import TextInput from "../basic/TextInput";
import { runTests, Test, TestConfig } from "../../api/entity/test_config";
import { Box, Sizes } from "../style";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  const [configName, setConfigName] = React.useState<string>("");
  const [loading, setLoading] = React.useState<boolean>(false);
  const [testConfig, setTestConfig] = React.useState<TestConfig>({
    Name: "",
    Tests: [],
  });

  console.log(testConfig);

  const destroyRequest = (e: any) => {
    e.preventDefault();
    destroyAll()
      .then((data) => console.log(data))
      .catch((error) => console.log(error));
  };

  const setConfig = (e: React.FormEvent) => {
    e.preventDefault();
    setTestConfig({
      ...testConfig,
      Name: configName,
    });
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setConfigName(e.target.value);
  };

  const run = (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    runTests(testConfig)
      .then(() => {
        setLoading(false);
      })
      .catch(() => {
        setLoading(false);
      });
  };

  const addNewTest = (test: Test) => (e: React.FormEvent) => {
    e.preventDefault();
    setTestConfig({
      ...testConfig,
      Tests: [...testConfig.Tests, test],
    });
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
        <TestForm addNewTest={addNewTest} />
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

export default TestContent;
