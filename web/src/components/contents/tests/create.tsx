/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { saveTests, Test, TestConfig } from "../../../api/entity/test_config";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import CreateTest from "./CreateTest";
import Message from "../../basic/Message";
import Table from "../../basic/Table";
import { useHistory } from "react-router-dom";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [message, setMessage] = React.useState<string>("");
  const [configName, setConfigName] = React.useState<string>("");
  const [testConfig, setTestConfig] = React.useState<TestConfig>({
    name: "",
    tests: [],
  });
  const history = useHistory();

  const setConfig = (e: React.FormEvent) => {
    e.preventDefault();
    setTestConfig({
      ...testConfig,
      name: configName,
    });
  };
  const addNewTest = (test: Test) => (e: React.FormEvent) => {
    e.preventDefault();
    if (!testConfig.name) {
      setMessage("Please set test group name on the left menu.");
      return;
    }
    setTestConfig({
      ...testConfig,
      tests: [...testConfig.tests, test],
    });
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage("");
    setConfigName(e.target.value);
  };

  const totalRequests = (): number => {
    let count = 0;
    if (testConfig && testConfig.tests.length) {
      testConfig.tests.forEach((test: Test) => {
        count += test.requestCount;
      });
    }
    return count;
  };

  const buildTable = () => {
    const content: any[] = [];

    testConfig?.tests.map((test: Test) => {
      const row: any[] = [
        test.url,
        test.method,
        test.requestCount,
        buttons("Edit", test),
        buttons("Delete", test),
      ];
      content.push(row);
    });
    console.log(content);
    return content;
  };

  const buttons = (text: string, test?: Test) => {
    switch (text) {
      case "Delete":
        return (
          <Button
            text={text}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              deleteTest(test!);
            }}
          />
        );
      case "Edit":
        return (
          <Button
            text={text}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              editTest(test!);
            }}
          />
        );
      case "Delete All":
        return (
          <Button
            text={text}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              deleteAllTests();
            }}
          />
        );
    }
  };

  const deleteTest = (test: Test) => {};
  const editTest = (test: Test) => {};
  const deleteAllTests = () => {
    setTestConfig({
      ...testConfig,
      tests: [],
    });
  };

  const save = () => {
    saveTests(testConfig)
      .then(() => {
        history.push("/tests");
      })
      .catch((error) => {
        setMessage(error);
      });
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        {testConfig && testConfig.name ? (
          <div css={leftConfigDiv}>
            <h3 css={h3title}>Test Group</h3>
            <span>
              Name: <b>{testConfig.name}</b>
            </span>
            <span>
              Total Requests: <b>{totalRequests()}</b>
            </span>
            <div>
              <Button type={1} text="Save" onClick={save} />
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
        {testConfig && testConfig.tests.length > 0 && (
          <Table
            title={[
              "URL",
              "Method",
              "Requests Count",
              "",
              buttons("Delete All"),
            ]}
            content={buildTable()}
          />
        )}

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
  min-height: 50rem;
  padding: 2rem;
`;

const rightColumn = css`
  width: 70%;
  padding: 2rem;
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
