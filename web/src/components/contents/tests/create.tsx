/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Test, TestConfig } from "../../../api/entity/test_config";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import CreateTest from "../../forms/tests/CreateTest";
import Message from "../../basic/Message";
import Table from "../../basic/Table";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [message, setMessage] = React.useState<string>("");
  const [configName, setConfigName] = React.useState<string>("");
  const [testConfig, setTestConfig] = React.useState<TestConfig>({
    Name: "",
    Tests: [
      {
        url: "dehaa.com/stests.com",
        requestCount: 1000,
        method: "get",
        goroutineCount: 1,
        expectedResponseBody: "",
        expectedResponseCode: -1,
        payload: "",
        transportConfig: { DisableKeepAlives: true },
      },
    ],
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

  const totalRequests = (): number => {
    let count = 0;
    if (testConfig && testConfig.Tests.length) {
      testConfig.Tests.map((test: Test) => {
        count += test.requestCount;
      });
    }
    return count;
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        {testConfig && testConfig.Name ? (
          <div css={leftConfigDiv}>
            <h3 css={h3title}>Test Group</h3>
            <span>
              Name: <b>{testConfig.Name}</b>
            </span>
            <span>
              Total Requests: <b>{totalRequests()}</b>
            </span>
            <div>
              <Button type={1} text="Save" onClick={() => {}} />
              <Button type={1} text="Update" onClick={() => {}} />
            </div>
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
        <Table
          title={["URL", "Method", "Requests Count"]}
          content={[
            ["x", "y", "z"],
            ["a", "b", "c"],
          ]}
        ></Table>
        {testConfig?.Tests.map((test: Test) => {
          return (
            <div css={configCss} key={test.url}>
              <div>
                <b>URL:</b>
                <span>{test.url}</span>
              </div>
              <div>
                <b>Method:</b>
                <span>{test.method.toUpperCase()}</span>
              </div>
              <div>
                <b>Requests:</b>
                <span>
                  <b>{test.requestCount}</b>
                </span>
              </div>
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
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  width: 100%;
  height: 4rem;
  padding: 1rem 0;
  border: 1px solid #e1e1e1;
  border-radius: 0.3rem;
  background-color: #e1e1e1;
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
  padding-bottom: 0.5rem;
`;

export default Create;
