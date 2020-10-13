/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import {
  deleteTestGroup,
  listTestGroup,
  runTestGroup,
  TestGroup,
} from "../../../api/entity/test_group";
import Table from "../../basic/Table";
import Button from "../../basic/Button";
import { leftContent } from "../../style";
import Message from "../../basic/Message";
import TestForm from "./test_form";
import { deleteTest, Test, updateTest } from "../../../api/entity/test";

interface Props {
  testGroup?: TestGroup;
}

const Show: React.FC<Props> = (props: Props) => {
  const [configs, setConfigs] = useState<TestGroup[]>();
  const [selectedTestGroup, setSelectedTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const [selectedTest, setSelectedTest] = useState<Test | null>(null);
  const [message, setMessage] = useState<string>("");
  const [updateSelectedGroupName, setUpdateSelectedGroupName] = useState<
    TestGroup
  >();

  React.useEffect(() => {
    listTestGroup()
      .then((response) => {
        setConfigs(response.data);
        setSelectedTestGroup(response.data[0]);
      })
      .catch((error) => console.log(error));
  }, []);

  const run = (testConfig: TestGroup) => (e: React.FormEvent) => {
    e.preventDefault();

    runTestGroup(testConfig)
      .then(() => {})
      .catch(() => {});
  };

  const buildTable = () => {
    const content: any[] = [];

    if (selectedTestGroup) {
      selectedTestGroup.tests.map((test: Test) => {
        const row: any[] = [
          test.url,
          test.method,
          test.requestCount,
          buttons("Edit", test),
          buttons("Delete", test),
        ];
        content.push(row);
      });
    }

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
              editTest(test!);
            }}
          />
        );
    }
  };

  const onUpdateTest = (test: Test) => {
    updateTest(test).then(() => {
      setMessage(JSON.stringify(test));
    });
  };

  const onDeleteTest = (test: Test): void => {
    deleteTest(test)
      .then(() => {
        setSelectedTestGroup({
          ...selectedTestGroup,
          tests: selectedTestGroup.tests.filter(
            (selectedTest: Test) => test.id !== selectedTest.id,
          ),
        });
        if (selectedTestGroup.tests.length <= 1) {
          deleteTestGroup(selectedTestGroup)
            .then(() => {
              setConfigs(
                configs?.filter(
                  (conf: TestGroup) => conf.id !== selectedTestGroup.id,
                ),
              );
            })
            .catch((error) => console.log(error));
        }
      })
      .catch((error) => setMessage(error));
  };

  const editTest = (test: Test): void => {
    setSelectedTest(test);
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        <h3 css={h3title}>Test Groups</h3>
        {configs?.map((config: TestGroup) => (
          <div
            css={css`
              ${leftContent}
            `}
            key={config.id}
            onClick={(e: React.MouseEvent) => {
              e.preventDefault();
              setSelectedTestGroup(config);
            }}
          >
            <div>
              <span>
                <b>{config.name}</b> Total Requests: <b>0</b>
              </span>
            </div>
          </div>
        ))}
        <Link to="/tests/create">
          <Button text="New Test Group" />
        </Link>
      </div>
      <div css={rightColumn}>
        {message ? <Message type="error" message={message} /> : ""}
        <Button text="Run" onClick={run(selectedTestGroup)} />
        <Button text="Update" />
        <Table
          title={["URL", "Method", "Requests Count", "", ""]}
          content={buildTable()}
        />
        {selectedTest && (
          <TestForm test={selectedTest} updateTest={onUpdateTest} />
        )}
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

const h3title = css`
  border-bottom: 0.1rem solid grey;
  margin-bottom: 0.5rem;
  padding-bottom: 0.5rem;
`;

export default Show;
