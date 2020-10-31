/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import {
  deleteTestGroup,
  listTestGroup,
  runTestGroup,
  TestGroup,
  updateTestGroup,
} from "../../../api/entity/test_group";
import Table from "../../basic/Table";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftContent, MediaQuery } from "../../style";
import Message, { MessageObj } from "../../basic/Message";
import TestForm from "./test_form";
import {
  runTest,
  deleteTest,
  saveTest,
  Test,
  updateTest,
} from "../../../api/entity/test";
import TextInput from "../../basic/TextInput";
import { FiPlay, FiTrash2, FiEdit } from "react-icons/fi";

interface Props {
  testGroup?: TestGroup;
}

const Show: React.FC<Props> = (props: Props) => {
  const [testGroups, setTestGroups] = useState<TestGroup[]>();
  const [selectedTestGroup, setSelectedTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const [selectedTest, setSelectedTest] = useState<Test | null>(null);
  const [message, setMessage] = useState<MessageObj>();
  const [updateSelectedGroupName, setUpdateSelectedGroupName] = useState<
    string
  >("");
  const [addNewTest, setAddNewTest] = useState<boolean>(false);

  React.useEffect(() => {
    listTestGroup()
      .then((response) => {
        setTestGroups(response.data);
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
      selectedTestGroup.tests.forEach((test: Test) => {
        const row: any[] = [
          test.url,
          test.method,
          test.request_count,
          buttons("Run", test),
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
      case "Run":
        return (
          <Button
            colorType={ButtonColorType.success}
            type={ButtonType.iconButton}
            icon={<FiPlay />}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              onRunTest(test!);
            }}
          />
        );
      case "Delete":
        return (
          <Button
            colorType={ButtonColorType.danger}
            type={ButtonType.iconButton}
            icon={<FiTrash2 />}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              onDeleteTest(test!);
            }}
          />
        );
      case "Edit":
        return (
          <Button
            colorType={ButtonColorType.secondary}
            type={ButtonType.iconButton}
            icon={<FiEdit />}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              editTest(test!);
            }}
          />
        );
    }
  };

  const onAddTest = (test: Test) => {
    test.test_group_id = selectedTestGroup.id!;
    saveTest(test).then(() => {
      const newTestGroups = testGroups?.map((tg: TestGroup) => {
        if (tg.id === selectedTestGroup.id) {
          tg.tests = [...selectedTestGroup.tests, test];
        }
        return tg;
      });

      setTestGroups(newTestGroups);
    });
  };

  const onUpdateTest = (test: Test) => {
    updateTest(test).then(() => {
      setSelectedTestGroup({
        ...selectedTestGroup,
        tests: selectedTestGroup.tests.map((t: Test) => {
          if (t.id === test.id) {
            return test;
          }
          return t;
        }),
      });
      setMessage({
        type: "success",
        message: "Test's been updated.",
      });
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
              setTestGroups(
                testGroups?.filter(
                  (conf: TestGroup) => conf.id !== selectedTestGroup.id,
                ),
              );
            })
            .catch((error) => console.log(error));
        }
      })
      .catch((error) => setMessage(error));
  };

  const onRunTest = (test: Test) => {
    runTest(test)
      .then((response) => console.log(response))
      .catch((error) => console.log(error));
  };

  const editTest = (test: Test): void => {
    setSelectedTest(test);
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        <h3 css={h3title}>Test Groups</h3>
        {testGroups?.map((config: TestGroup) => (
          <div
            css={leftContent(config.id === selectedTestGroup.id)}
            key={config.id}
            onClick={(e: React.MouseEvent) => {
              e.preventDefault();
              setSelectedTestGroup(config);
            }}
          >
            <div>
              <span>
                <b>{config.name}</b>
              </span>
            </div>
          </div>
        ))}
        <Link to="/tests/create">
          <Button text="New Test Group" />
        </Link>
      </div>
      <div css={rightColumn}>
        {message ? <Message type="error" message={message.message} /> : ""}
        {selectedTestGroup && selectedTestGroup.tests.length > 0 ? (
          <React.Fragment>
            {updateSelectedGroupName ? (
              <React.Fragment>
                <TextInput
                  name="testGroupName"
                  type="text"
                  value={updateSelectedGroupName}
                  onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                    setUpdateSelectedGroupName(e.target.value);
                  }}
                />
                <Button
                  text="Save"
                  onClick={(e: React.FormEvent) => {
                    e.preventDefault();
                    updateTestGroup(selectedTestGroup).then(() => {
                      const newTestGroups = testGroups?.map((tg: TestGroup) => {
                        if (tg.id === selectedTestGroup.id) {
                          tg.name = updateSelectedGroupName;
                        }
                        return tg;
                      });

                      setTestGroups(newTestGroups);
                      setSelectedTestGroup({
                        ...selectedTestGroup,
                        name: updateSelectedGroupName,
                      });
                      setUpdateSelectedGroupName("");
                    });
                  }}
                />
              </React.Fragment>
            ) : (
              <React.Fragment>
                <Button
                  text="Add New Test"
                  onClick={() => {
                    setAddNewTest(true);
                  }}
                />
                <Button text="Run" onClick={run(selectedTestGroup)} />
                <Button
                  text="Update"
                  onClick={() => {
                    setUpdateSelectedGroupName(selectedTestGroup.name);
                  }}
                />
              </React.Fragment>
            )}
            <Table
              title={["URL", "Method", "Requests Count", "", "", ""]}
              content={buildTable()}
            />
          </React.Fragment>
        ) : (
          <Message
            type="warning"
            message="There is no tests here, Please create a new test group"
          />
        )}

        {addNewTest && (
          <TestForm testGroup={selectedTestGroup} addTest={onAddTest} />
        )}
        {selectedTest && (
          <TestForm
            testGroup={selectedTestGroup}
            test={selectedTest}
            updateTest={onUpdateTest}
          />
        )}
      </div>
    </div>
  );
};

const container = css`
  display: flex;
  width: 100%;
  flex-direction: column;
  ${MediaQuery[1]} {
    flex-direction: row;
  }
`;

const leftColumn = css`
  background-color: #e3e3e3;
  width: 100%;
  padding: 2rem;

  ${MediaQuery[1]} {
    width: 30%;
    min-height: 50rem;
  }
`;

const rightColumn = css`
  width: 100%;
  ${MediaQuery[1]} {
    width: 70%;
    min-height: 50rem;
  }
  padding: 2rem;
`;

const h3title = css`
  margin-bottom: 0.5rem;
  padding-bottom: 0.5rem;
`;

export default Show;
