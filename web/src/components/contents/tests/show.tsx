/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import {
  deleteTestReq,
  deleteTestsReq,
  listTests,
  runTests,
  Test,
  TestConfig,
  updateTestReq,
} from "../../../api/entity/test_config";
import Table from "../../basic/Table";
import Button from "../../basic/Button";
import { leftContent } from "../../style";
import Message from "../../basic/Message";
import TestForm from "./test_form";

interface Props {
  testConfg?: TestConfig;
}

const Show: React.FC<Props> = (props: Props) => {
  const [configs, setConfigs] = useState<TestConfig[]>();
  const [selectedConfig, setSelectedConfig] = useState<TestConfig>({
    name: "",
    tests: [],
  });
  const [selectedTest, setSelectedTest] = useState<Test | null>(null);
  const [message, setMessage] = useState<string>("");
  const [showLinkList, setShowLinkList] = useState<TestConfig>();

  React.useEffect(() => {
    listTests()
      .then((response) => {
        setConfigs(response.data);
        setSelectedConfig(response.data[0]);
      })
      .catch((error) => console.log(error));
  }, []);

  const run = (testConfig: TestConfig) => (e: React.FormEvent) => {
    e.preventDefault();

    runTests(testConfig)
      .then(() => {})
      .catch(() => {});
  };

  const buildTable = () => {
    const content: any[] = [];

    if (selectedConfig) {
      selectedConfig.tests.map((test: Test) => {
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
    }
  };

  const updateTest = (test: Test) => {
    updateTestReq(test).then(() => {
      setMessage(JSON.stringify(test));
    });
  };

  const deleteTest = (test: Test): void => {
    deleteTestReq(test)
      .then(() => {
        setSelectedConfig({
          ...selectedConfig,
          tests: selectedConfig.tests.filter(
            (selectedTest: Test) => test.id !== selectedTest.id,
          ),
        });
        if (selectedConfig.tests.length <= 1) {
          deleteTestsReq(selectedConfig)
            .then(() => {
              setConfigs(
                configs?.filter(
                  (conf: TestConfig) => conf.id !== selectedConfig.id,
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
        {configs?.map((config: TestConfig) => (
          <div
            css={css`
              ${leftContent}
            `}
            key={config.id}
            onClick={(e: React.MouseEvent) => {
              e.preventDefault();
              setSelectedConfig(config);
            }}
            onMouseEnter={() => {
              setShowLinkList(config);
            }}
            onMouseLeave={() => {
              setShowLinkList(undefined);
            }}
          >
            {showLinkList === undefined || showLinkList !== config ? (
              <div>
                <span>
                  <b>{config.name}</b> Total Requests: <b>0</b>
                </span>
              </div>
            ) : (
              <div>
                <b>{config.name}</b>
                <Button type={1} text="Run" onClick={run(config)} />
                <Button type={1} text="Update" />
              </div>
            )}
          </div>
        ))}
        <Link to="/tests/create">
          <Button text="New Test Group" />
        </Link>
      </div>
      <div css={rightColumn}>
        {message ? <Message type="error" message={message} /> : ""}
        <Table
          title={["URL", "Method", "Requests Count", "", ""]}
          content={buildTable()}
        />
        {selectedTest && (
          <TestForm test={selectedTest} updateTest={updateTest} />
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

const linkList = css`
  display: block;
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
