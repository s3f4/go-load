/** @jsx jsx */
import React, { Fragment } from "react";
import { jsx, css } from "@emotion/core";
import Paginator from "../../basic/Paginator";
import { Link } from "react-router-dom";
import { leftContent } from "../../style";
import { listTestGroup, TestGroup } from "../../../api/entity/test_group";
import Button, { ButtonType } from "../../basic/Button";
import { FiPlusCircle } from "react-icons/fi";

interface Props {
  testGroups: TestGroup[] | undefined;
  selectedTestGroup: TestGroup;
  setSelectedTestGroup: (testGroup: TestGroup) => any;
  setTestGroups: (testGroups: TestGroup[]) => any;
}

const TestGroupLeftMenu: React.FC<Props> = (props: Props) => {
  return (
    <Fragment>
      <div css={titleDiv}>
        <h3 css={h3title}>Test Groups</h3>
        <Link to="/tests/create">
          <Button
            type={ButtonType.iconButton}
            icon={<FiPlusCircle />}
            text="New Test Group"
          />
        </Link>
      </div>
      {props.testGroups &&
        props.testGroups.map((testGroup: TestGroup) => (
          <div
            css={leftContent(testGroup.id === props.selectedTestGroup.id)}
            key={testGroup.id}
            onClick={(e: React.MouseEvent) => {
              e.preventDefault();
              props.setSelectedTestGroup(testGroup);
            }}
          >
            <div>
              <span>
                <b>{testGroup.name}</b>
              </span>
            </div>
          </div>
        ))}
      <Paginator
        fetcher={listTestGroup}
        setter={(data) => {
          props.setTestGroups(data);
          props.setSelectedTestGroup(data[0]);
        }}
      />
    </Fragment>
  );
};

const h3title = css`
  margin-bottom: 0.5rem;
  padding-bottom: 0.5rem;
`;

const titleDiv = css`
  display: flex;
  justify-content: space-between;
`;

export default TestGroupLeftMenu;
