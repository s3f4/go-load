/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import {
  listTests,
  runTests,
  Test,
  TestConfig,
} from "../../../api/entity/test_config";
import Table from "../../basic/Table";
import Button from "../../basic/Button";
import { Borders } from "../../style";

interface Props {
  testConfg?: TestConfig;
}

const Show: React.FC<Props> = (props: Props) => {
  const [configs, setConfigs] = useState<TestConfig[]>();
  const [selectedConfig, setSelectedConfig] = useState<TestConfig>();
  console.log(configs);

  React.useEffect(() => {
    listTests()
      .then((response) => {
        setConfigs(response.data);
      })
      .catch((error) => console.log(error));
  }, []);

  const run = (e: React.FormEvent) => {
    e.preventDefault();

    runTests(props.testConfg!)
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
  const deleteAllTests = () => {};

  return (
    <div css={container}>
      <div css={leftColumn}>
        <h3 css={h3title}>Test Groups</h3>
        {configs?.map((config: TestConfig) => (
          <div
            css={leftConfigDiv}
            key={config.id}
            onClick={(e: React.MouseEvent) => {
              e.preventDefault();
              setSelectedConfig(config);
            }}
          >
            <span>
              Name: <b>{config.name}</b>
            </span>
            <span>
              Total Requests: <b>0</b>
            </span>
          </div>
        ))}
        <Link to="/tests/create">
          <Button text="New Test Group" />
        </Link>
      </div>
      <div css={rightColumn}>
        <Table
          title={["URL", "Method", "Requests Count", "", buttons("Delete All")]}
          content={buildTable()}
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
  min-height: 5rem;
  display: flex;
  flex-direction: column;
  border-bottom: ${Borders.border1};
  border-radius: 0.5rem;
  padding: 1rem;
  cursor: pointer;
`;

const h3title = css`
  border-bottom: 0.1rem solid grey;
  margin-bottom: 0.5rem;
  padding-bottom: 0.5rem;
`;

export default Show;
