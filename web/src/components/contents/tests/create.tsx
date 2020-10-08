/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Test, TestConfig } from "../../../api/entity/test_config";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import CreateTest from "../../forms/tests/CreateTest";
import Message from "../../basic/Message";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [message, setMessage] = React.useState<string>("");
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
    if (!testConfig.Name) {
      setMessage("Please set test group name on the left menu.");
      return;
    }
    setTestConfig({
      ...testConfig,
      Tests: [...testConfig.Tests, test],
    });
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage("");
    setConfigName(e.target.value);
  };
  return (
    <div css={container}>
      <div css={leftColumn}>
        {testConfig && testConfig.Name ? (
          <div css={leftConfigDiv}>
            <h3 css={h3title}>Test Group</h3>
            <span>
              {" "}
              Name: <b>{testConfig.Name}</b>
            </span>
            <div>
              <Button type={1} text="Save" onClick={() => {}} />
              <Button type={1} text="Update" onClick={() => {}} />
            </div>
            {testConfig?.Tests.map((test: Test) => {
              return (
                <div css={configCss} key={test.url}>
                  URL : {test.url} - Method: {test.method} - Request Count:{" "}
                  {test.requestCount}
                </div>
              );
            })}
          </div>
        ) : (
          <React.Fragment>
            <TextInput
              name={"Test Config Name"}
              label={"Test Config Name"}
              onChange={handleChange}
            />
            <Button text="Create" onClick={setConfig} />
          </React.Fragment>
        )}
      </div>
      <div css={rightColumn}>
        {message ? <Message type="error" message={message} /> : ""}
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

const leftConfigDiv = css`
  width: 100%;
  height: 5rem;
  display: flex;
  flex-direction: column;
`;

const h3title = css`
  border-bottom: 0.1rem solid grey;
  margin-bottom: 2rem;
`;

export default Create;
