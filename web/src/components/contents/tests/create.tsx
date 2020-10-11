/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { saveTests, Test, TestConfig } from "../../../api/entity/test_config";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import CreateTest from "./CreateTest";
import Message from "../../basic/Message";
import Table from "../../basic/Table";
import { useHistory } from "react-router-dom";
import { isEqual } from "lodash";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [editTest, setEditTest] = useState<Test | undefined>(undefined);
  const [message, setMessage] = useState<string>("");
  const [configName, setConfigName] = useState<string>("");
  const [testConfig, setTestConfig] = useState<TestConfig>({
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
  const addNewTest = (test: Test) => {
    if (!testConfig.name) {
      setMessage("Please set test group name on the left menu.");
      return;
    }

    setEditTest(undefined);

    let equal = false;
    testConfig.tests.forEach((t: Test) => {
      if (isEqual(t, test)) {
        equal = true;
      }
    });

    if (equal) {
      setMessage("This test was already created");
      return;
    }
    test.id = new Date().getUTCMilliseconds();
    setTestConfig({
      ...testConfig,
      tests: [...testConfig.tests, test],
    });
  };

  const updateNewTest = (test: Test) => {
    const index = testConfig.tests.findIndex((t) => t.id === test.id);
    if (index !== -1) {
      setTestConfig({
        ...testConfig,
        tests: [
          ...testConfig.tests.slice(0, index),
          Object.assign({}, testConfig.tests[index], test),
          ...testConfig.tests.slice(index + 1),
        ],
      });
    }
    setEditTest(undefined);
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
              setEditTest(test!);
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

  const deleteTest = (test: Test) => {
    setTestConfig({
      ...testConfig,
      tests: testConfig.tests.filter((t: Test) => !isEqual(t, test)),
    });
  };

  const deleteAllTests = () => {
    setTestConfig({
      name: "",
      tests: [],
    });
  };

  const saveTestConfig = () => {
    if (!testConfig.tests.length) {
      setMessage("Please create a test to save test group");
      return;
    }

    saveTests(testConfig)
      .then(() => {
        history.push("/tests");
      })
      .catch((error) => {
        setMessage(error);
      });
  };

  const updateTestConfig = () => {
    setTestConfig({
      ...testConfig,
      name: "",
    });
  };

  const triggerMessage = (message: string) => () => {
    setMessage(message);
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
              <Button text="Save" onClick={saveTestConfig} />
              <Button text="Update" onClick={updateTestConfig} />
            </div>
          </div>
        ) : (
          <React.Fragment>
            <TextInput
              name={"Test Group Name"}
              label={"Test Group Name"}
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

        <CreateTest
          test={editTest}
          setMessage={triggerMessage("")}
          addNewTest={addNewTest}
          updateNewTest={updateNewTest}
        />
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
