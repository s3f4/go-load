/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { useHistory } from "react-router-dom";
import { TestGroup } from "../../../api/entity/test_group";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftColumn, MediaQuery, rightColumn } from "../../style";
import { Test, listTestsOfTestGroup } from "../../../api/entity/test";
import { FiActivity } from "react-icons/fi";
import RTable, { RTableRow } from "../../basic/RTable";
import TestGroupLeftMenu from "../tests/TestGroupLeftMenu";
import Message from "../../basic/Message";

const ListStats: React.FC = () => {
  const [testGroups, setTestGroups] = useState<TestGroup[]>();
  const [selectedTestGroup, setSelectedTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const history = useHistory();

  const buildTable = (tests: Test[]): RTableRow[] => {
    const rows: RTableRow[] = [];

    tests.forEach((test: Test) => {
      const row: RTableRow = {
        rowStyle:`${test}`,
        columns: [
          { content: <b>{test.name}</b> },
          { content: test.method },
          { content: test.request_count },
          { content: <div>{buttons("Stats", test)}</div> },
        ],
      };
      rows.push(row);
    });

    return rows;
  };

  const buttons = (text: string, test?: Test) => {
    switch (text) {
      case "Stats":
        return (
          <Button
            colorType={ButtonColorType.info}
            type={ButtonType.iconButton}
            icon={<FiActivity />}
            onClick={(e: React.FormEvent) => {
              e.preventDefault();
              history.push(`/stats/${test?.id}`);
            }}
          />
        );
    }
  };

  const fetcher = () => {
    if (selectedTestGroup && selectedTestGroup.id) {
      return listTestsOfTestGroup(selectedTestGroup?.id!);
    }
    return listTestsOfTestGroup(1);
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        <TestGroupLeftMenu
          testGroups={testGroups}
          selectedTestGroup={selectedTestGroup}
          setSelectedTestGroup={setSelectedTestGroup}
          setTestGroups={setTestGroups}
        />
      </div>
      <div css={rightColumn}>
        {selectedTestGroup && selectedTestGroup.tests.length > 0 ? (
          <RTable
            builder={buildTable}
            fetcher={fetcher()}
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
                width: "15%",
              },
              {
                header: "Request Count",
                accessor: "request_count",
                sortable: true,
                width: "25%",
              },
              {
                header: "Actions",
                sortable: false,
                width: "10%",
              },
            ]}
          />
        ) : (
          <Message
            type="warning"
            message="There is no tests here, Please create a new test group"
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

export default ListStats;
