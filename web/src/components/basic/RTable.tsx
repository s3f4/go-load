/** @jsx jsx */
import React, { Fragment, useState } from "react";
import { jsx, css } from "@emotion/core";
import { Colors, MediaQuery } from "../style";

export interface TableTitle {
  header: string;
  accessor: string;
  sortable: boolean;
  row: object;
}
interface Props {
  title: any[];
  content: any[][];
}

const RTable: React.FC<Props> = (props: Props) => {
  const [total, setTotal] = useState<number>();
  const [page, setPage] = useState<number>();
  const [sort, setSort] = useState<string>();

  return (
    <Fragment>
      <div css={container}>
        <div css={row(true)}>
          {props.title.map((title, index) => (
            <div css={columnStyle} key={index}>
              <b>{title}</b>
            </div>
          ))}
        </div>

        {props.content.map((rows, index) => {
          return (
            <div key={index} css={row(false)}>
              {rows.map((column, colIndex) => (
                <div css={columnStyle} key={colIndex}>
                  {column}
                </div>
              ))}
            </div>
          );
        })}
      </div>

      <div css={mobileContainer}>
        {props.content.map((rows, index) => {
          return (
            <div key={index} css={mobileRow}>
              {rows.map((column, colIndex) => (
                <div css={mobileFlex} key={colIndex}>
                  <b>{props.title[colIndex]}</b>
                  {column}
                </div>
              ))}
            </div>
          );
        })}
      </div>
    </Fragment>
  );
};

const mobileContainer = css`
  display: block;
  ${MediaQuery[2]} {
    display: none;
  }
  width: 100%;
  border: 1px solid #e1e1e1;
  border-radius: 0.5rem;
  text-align: left;
  padding: 1rem 1rem 1rem 1rem;
`;

const mobileFlex = css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 5rem;
  min-height: 4rem;
`;

const mobileRow = css`
  margin-top: 1rem;
  padding: 2rem;
  background-color: #e1e1e1;
`;

const row = (title?: boolean) => css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 4.5rem;
  border-bottom: 1px solid black;
  background-color: ${title ? "#007d9c" : "none"};
  color: ${title ? "white" : "none"};
`;

const columnStyle = css`
  flex: 0 1 20rem;
  padding-left: 1rem;
  padding-top: 1rem;
  text-align: center;
  width: 7rem;
`;

const container = css`
  display: none;
  ${MediaQuery[2]} {
    display: flex;
    flex-direction: column;
    width: 100%;
    border: 1px solid #e1e1e1;
    border-radius: 0.5rem;
    background-color: #e1e1e1;
    text-align: left;
    padding: 1rem 1rem 1rem 1rem;
  }
`;

const tableTitle = css`
  font-weight: bold;
`;

export default RTable;
