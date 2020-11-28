/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { useHistory } from "react-router-dom";
import { TestGroup } from "../../../api/entity/test_group";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftColumn, MediaQuery, rightColumn } from "../../style";
import { Test, listTestsOfTestGroup } from "../../../api/entity/test";
import { FiActivity } from "react-icons/fi";
import RTable from "../../basic/RTable";
import TestGroupLeftMenu from "../tests/TestGroupLeftMenu";

const ListStats: React.FC = () => {
  const [testGroups, setTestGroups] = useState<TestGroup[]>();
  const [selectedTestGroup, setSelectedTestGroup] = useState<TestGroup>({
    name: "",
    tests: [],
  });
  const history = useHistory();

  const buildTable = (tests: Test[]): any[][] => {
    const content: any[] = [];

    tests.forEach((test: Test) => {
      const row: any[] = [
        <b>{test.name}</b>,
        test.method,
        test.request_count,
        <div>{buttons("Stats", test)}</div>,
      ];
      content.push(row);
    });

    return content;
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
