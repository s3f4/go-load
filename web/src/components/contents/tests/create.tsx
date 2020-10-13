/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import Message from "../../basic/Message";
import Table from "../../basic/Table";
import { useHistory } from "react-router-dom";
import { isEqual } from "lodash";
import TestForm from "./test_form";
import { Test } from "../../../api/entity/test";
import { saveTestGroup, TestGroup } from "../../../api/entity/test_group";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  const [editTest, setEditTest] = useState<Test | undefined>(undefined);
  const [message, setMessage] = useState<string>("");
  const [testGroupName, setTestGroupName] = useState<string>("");
  const [testGroup, setTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const history = useHistory();

  const setConfig = (e: React.FormEvent) => {
    e.preventDefault();
    setTestGroup({
      ...testGroup,
      name: testGroupName,
    });
  };
  const addNewTest = (test: Test) => {
    if (!testGroup.name) {
      setMessage("Please set test group name on the left menu.");
      return;
    }

    setEditTest(undefined);

    let equal = false;
    testGroup.tests.forEach((t: Test) => {
      if (isEqual(t, test)) {
        equal = true;
      }
    });

    if (equal) {
      setMessage("This test was already created");
      return;
    }
    test.id = new Date().getUTCMilliseconds();
    setTestGroup({
      ...testGroup,
      tests: [...testGroup.tests, test],
    });
  };

  const updateNewTest = (test: Test) => {
    const index = testGroup.tests.findIndex((t) => t.id === test.id);
    if (index !== -1) {
      setTestGroup({
        ...testGroup,
        tests: [
          ...testGroup.tests.slice(0, index),
          Object.assign({}, testGroup.tests[index], test),
          ...testGroup.tests.slice(index + 1),
        ],
      });
    }
    setEditTest(undefined);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage("");
    setTestGroupName(e.target.value);
  };

  const totalRequests = (): number => {
    let count = 0;
    if (testGroup && testGroup.tests.length) {
      testGroup.tests.forEach((test: Test) => {
        count += test.requestCount;
      });
    }
    return count;
  };

  const buildTable = () => {
    const content: any[] = [];

    testGroup?.tests.map((test: Test) => {
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
              onDeleteTest(test!);
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

  const onDeleteTest = (test: Test) => {
    setTestGroup({
      ...testGroup,
      tests: testGroup.tests.filter((t: Test) => !isEqual(t, test)),
    });
  };

  const deleteAllTests = () => {
    setTestGroup({
      name: "",
      tests: [],
    });
  };

  const onSaveTestGroup = () => {
    if (!testGroup.tests.length) {
      setMessage("Please create a test to save test group");
      return;
    }

    saveTestGroup(testGroup)
      .then(() => {
        history.push("/tests");
      })
      .catch((error) => {
        setMessage(error);
      });
  };

  const onUpdateTestGroup = () => {
    setTestGroup({
      ...testGroup,
      name: "",
    });
  };

  const triggerMessage = (message: string) => () => {
    setMessage(message);
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        {testGroup && testGroup.name ? (
          <div css={leftConfigDiv}>
            <h3 css={h3title}>Test Group</h3>
            <span>
              Name: <b>{testGroup.name}</b>
            </span>
            <span>
              Total Requests: <b>{totalRequests()}</b>
            </span>
            <div>
              <Button text="Save" onClick={onSaveTestGroup} />
              <Button text="Update" onClick={onUpdateTestGroup} />
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
        {testGroup && testGroup.tests.length > 0 && (
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

        <TestForm
          test={editTest}
          setMessage={triggerMessage("")}
          addTest={addNewTest}
          updateTest={updateNewTest}
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
