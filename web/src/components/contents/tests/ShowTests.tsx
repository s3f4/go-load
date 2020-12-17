/** @jsx jsx */
import React, { useState, useEffect, ChangeEvent } from "react";
import { jsx, css } from "@emotion/core";
import {
  deleteTestGroup,
  TestGroup,
  updateTestGroup,
} from "../../../api/entity/test_group";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftColumn, MediaQuery, rightColumn } from "../../style";
import Message, { IMessage } from "../../basic/Message";
import TestForm from "./TestForm";
import {
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
import RTable, { IRTableRow } from "../../basic/RTable";
import TestGroupLeftMenu from "./TestGroupLeftMenu";
import RunTests from "./RunTests";
import { getItems, search, removeOne } from "../../basic/localStorage";

export interface RunConfig {
  test: Test;
  loading: boolean;
  passed: boolean;
  started: boolean;
  finished: boolean;
  error: any;
}

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
  // Test Run States
  const [testGroupRun, setTestGroupRun] = useState<TestGroup>();
  const [testRun, setTestRun] = useState<Test>();
  const [runConfigs, setRunConfigs] = useState<RunConfig[]>(
    getItems("run_configs") || [],
  );

  useEffect(() => {
    getInstanceInfo()
      .then((response) => {
        setInstances(response.data.configs);
      })
      .catch(() => {});
    return () => {};
  }, []);

  const onRunTestGroup = (testGroup: TestGroup) => (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(undefined);
    setTestGroupRun(testGroup);
  };

  const onRunTest = (test: Test) => {
    setMessage(undefined);
    setSelectedTest(undefined);
    setAddNewTest(false);
    setTestRun(test);
  };

  const buildTable = (tests: Test[]): IRTableRow[] => {
    const rows: IRTableRow[] = [];

    tests.forEach((test: Test) => {
      const row: IRTableRow = {
        rowStyle: "",
        allColumns: [
          { content: test.id },
          { content: test.name },
          { content: test.test_group?.name },
          { content: test.url },
          { content: test.method },
          { content: test.request_count },
          { content: test.goroutine_count },
          { content: JSON.stringify(test.headers) },
          { content: JSON.stringify(test.transport_config) },
          { content: test.expected_response_code },
          { content: test.expected_response_body },
          { content: test.expected_first_byte_time },
          { content: test.expected_connection_time },
          { content: test.expected_dns_time },
          { content: test.expected_tls_time },
          { content: test.expected_total_time },
        ],
        columns: [
          { content: <b>{test.name}</b> },
          { content: test.method },
          { content: test.request_count },
          {
            content: (
              <div>
                {buttons("Run", test)}
                {buttons("Edit", test)}
                {buttons("Delete", test)}
              </div>
            ),
          },
        ],
      };
      rows.push(row);
    });

    return rows;
  };

  const buttons = (text: string, test?: Test) => {
    switch (text) {
      case "Run":
        return (
          <Button
            colorType={ButtonColorType.success}
            type={ButtonType.iconButton}
            icon={<FiPlay />}
            disabled={
              !instances ||
              search("run_configs", [{ key: "test.id", value: test?.id }]) !==
                -1
            }
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
      setSelectedTestGroup({
        ...selectedTestGroup,
        tests: [...selectedTestGroup.tests, test],
      });
    });
    setAddNewTest(false);
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
    removeOne("run_configs", [{ key: "test.id", value: test.id }]);
    setRunConfigs(getItems("run_configs"));
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
        <RunTests
          test={testRun}
          testGroup={testGroupRun}
          setRunConfigs={setRunConfigs}
          runConfigs={runConfigs}
          clear={() => {
            setTestRun(undefined);
            setTestGroupRun(undefined);
            setRunConfigs([]);
          }}
        />
        {selectedTestGroup &&
        selectedTestGroup.tests.length > 0 &&
        !testGroupRun ? (
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
                  onClick={onRunTestGroup(selectedTestGroup)}
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
            <RTable
              builder={buildTable}
              fetcher={listTestsOfTestGroup(selectedTestGroup?.id!)}
              trigger={{ selectedTestGroup }}
              title={[
                {
                  header: "Name",
                  accessor: "name",
                  sortable: true,
                  width: "45%",
                },
                {
                  header: "Method",
                  accessor: "Method",
                  sortable: true,
                  width: "15%",
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
              allTitles={[
                { header: "ID" },
                { header: "Test Name" },
                { header: "Test Group Name" },
                { header: "URL" },
                { header: "Method" },
                { header: "Request Count" },
                { header: "Goroutine Count" },
                { header: "Headers" },
                { header: "Transport Config" },
                { header: "Expected Response Status Code" },
                { header: "Expected Response Body" },
                { header: "Expected Response First Byte Time" },
                { header: "Expected Response Connect Time" },
                { header: "Expected Response DNS Time" },
                { header: "Expected Response TLS Time" },
                { header: "Expected Response Total Time" },
              ]}
            />
          </React.Fragment>
        ) : !testGroupRun ? (
          <Message
            type="warning"
            message="There is no tests here, Please create a new test group"
          />
        ) : (
          ""
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
