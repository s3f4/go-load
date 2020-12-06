/** @jsx jsx */
import React, { useState, useEffect, ChangeEvent } from "react";
import { jsx, css } from "@emotion/core";
import {
  deleteTestGroup,
  runTestGroup,
  TestGroup,
  updateTestGroup,
} from "../../../api/entity/test_group";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftColumn, MediaQuery, rightColumn } from "../../style";
import Message, { IMessage } from "../../basic/Message";
import TestForm from "./TestForm";
import {
  runTest,
  deleteTest,
  saveTest,
  Test,
  updateTest,
  listTestsOfTestGroup,
} from "../../../api/entity/test";
import TextInput from "../../basic/TextInput";
import {
  FiPlay,
  FiTrash2,
  FiEdit,
  FiPlayCircle,
  FiEdit2,
  FiPlusCircle,
} from "react-icons/fi";
import { getInstanceInfo, Instance } from "../../../api/entity/instance";
import RTable from "../../basic/RTable";
import TestGroupLeftMenu from "./TestGroupLeftMenu";
import RunTests from "./RunTests";

const ShowTests: React.FC = () => {
  const [instances, setInstances] = useState<Instance[] | undefined>();
  const [testGroups, setTestGroups] = useState<TestGroup[]>();
  const [selectedTestGroup, setSelectedTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const [selectedTest, setSelectedTest] = useState<Test | undefined>(undefined);
  const [message, setMessage] = useState<IMessage>();
  const [
    updateSelectedGroupName,
    setUpdateSelectedGroupName,
  ] = useState<string>("");
  const [addNewTest, setAddNewTest] = useState<boolean>(false);

  useEffect(() => {
    getInstanceInfo()
      .then((response) => {
        setInstances(response.data.configs);
      })
      .catch(() => {});
    return () => {};
  }, []);

  const run = (testConfig: TestGroup) => (e: React.FormEvent) => {
    e.preventDefault();

    runTestGroup(testConfig)
      .then(() => {})
      .catch(() => {});
  };

  const buildTable = (tests: Test[]): any[][] => {
    const content: any[] = [];

    tests.forEach((test: Test) => {
      const row: any[] = [
        <b>{test.name}</b>,
        test.method,
        test.request_count,
        <div>
          {buttons("Run", test)}
          {buttons("Edit", test)}
          {buttons("Delete", test)}
        </div>,
      ];
      content.push(row);
    });

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
            disabled={!instances}
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
      setSelectedTest(undefined);
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
    setAddNewTest(false);
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        <TestGroupLeftMenu
          testGroups={testGroups}
          selectedTestGroup={selectedTestGroup}
          setSelectedTestGroup={(testGroup) => {
            setSelectedTest(undefined);
            setMessage(undefined);
            setSelectedTestGroup(testGroup);
          }}
          setTestGroups={setTestGroups}
        />
      </div>
      <div css={rightColumn}>
        {!instances ? (
          <Message
            type="error"
            message={"You have to create instances to run tests"}
          />
        ) : (
          ""
        )}
        {message ? (
          <Message type={message.type} message={message.message} />
        ) : (
          ""
        )}
        {selectedTestGroup && selectedTestGroup.tests.length > 0 ? (
          <React.Fragment>
            {updateSelectedGroupName ? (
              <div css={updateTestGroupNameDiv}>
                <div css={updateTestGroupNameTI}>
                  <TextInput
                    name="testGroupName"
                    type="text"
                    value={updateSelectedGroupName}
                    onChange={(
                      e:
                        | ChangeEvent<HTMLInputElement>
                        | ChangeEvent<HTMLTextAreaElement>,
                    ) => {
                      setUpdateSelectedGroupName(e.target.value);
                    }}
                  />
                </div>
                <div
                  css={css`
                    margin: 0.9rem;
                  `}
                >
                  <Button
                    text="Save"
                    onClick={(e: React.FormEvent) => {
                      e.preventDefault();
                      updateTestGroup(selectedTestGroup).then(() => {
                        const newTestGroups = testGroups?.map(
                          (tg: TestGroup) => {
                            if (tg.id === selectedTestGroup.id) {
                              tg.name = updateSelectedGroupName;
                            }
                            return tg;
                          },
                        );

                        setTestGroups(newTestGroups);
                        setSelectedTestGroup({
                          ...selectedTestGroup,
                          name: updateSelectedGroupName,
                        });
                        setUpdateSelectedGroupName("");
                      });
                    }}
                  />
                </div>
              </div>
            ) : (
              <div css={buttonsDiv}>
                <Button
                  text="New Test"
                  colorType={ButtonColorType.primary}
                  type={ButtonType.iconTextButton}
                  icon={<FiPlusCircle />}
                  onClick={() => {
                    setAddNewTest(true);
                    setSelectedTest(undefined);
                  }}
                />
                <Button
                  text="Run All"
                  colorType={ButtonColorType.success}
                  type={ButtonType.iconTextButton}
                  icon={<FiPlayCircle />}
                  onClick={run(selectedTestGroup)}
                  disabled={!instances}
                />
                <Button
                  text="Update Name"
                  colorType={ButtonColorType.secondary}
                  type={ButtonType.iconTextButton}
                  icon={<FiEdit2 />}
                  onClick={() => {
                    setUpdateSelectedGroupName(selectedTestGroup.name);
                  }}
                />
              </div>
            )}
            <RunTests tests={selectedTestGroup.tests} />
            <RTable
              builder={buildTable}
              fetcher={listTestsOfTestGroup(selectedTestGroup?.id!)}
              trigger={selectedTestGroup}
              title={[
                {
                  header: "Name",
                  accessor: "name",
                  sortable: true,
                  width: "50%",
                },
                {
                  header: "Method",
                  accessor: "Method",
                  sortable: true,
                  width: "10%",
                },
                {
                  header: "Request Count",
                  accessor: "request_count",
                  sortable: true,
                  width: "20%",
                },
                {
                  header: "Actions",
                  sortable: false,
                  width: "20%",
                },
              ]}
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

const buttonsDiv = css`
  display: flex;
  justify-content: space-between;
  width: 38rem;
  margin: 0.5rem auto;
  ${MediaQuery[1]} {
    margin: 0.5rem;
  }
`;

const updateTestGroupNameDiv = css`
  display: flex;
  justify-content: space-around;
  margin: 2rem 0.5rem;
`;

const updateTestGroupNameTI = css`
  width: 80%;
`;
export default ShowTests;
