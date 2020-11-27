/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { useHistory } from "react-router-dom";
import { listTestGroup, TestGroup } from "../../../api/entity/test_group";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { leftColumn, leftContent, MediaQuery, rightColumn } from "../../style";
import { Test, listTestsOfTestGroup } from "../../../api/entity/test";
import { FiActivity } from "react-icons/fi";
import RTable from "../../basic/RTable";
import Paginator from "../../basic/Paginator";

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
        test.name,
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

  return (
    <div css={container}>
      <div css={leftColumn}>
        <h3 css={h3title}>Test Groups</h3>
        {testGroups &&
          testGroups.map((config: TestGroup) => (
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
        <Paginator
          fetcher={listTestGroup}
          setter={(data) => {
            setTestGroups(data);
            setSelectedTestGroup(data[0]);
          }}
        />
      </div>
      <div css={rightColumn}>
        <RTable
          builder={buildTable}
          fetcher={listTestsOfTestGroup(selectedTestGroup?.id!)}
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

const h3title = css`
  margin-bottom: 0.5rem;
  padding-bottom: 0.5rem;
`;

export default ListStats;
