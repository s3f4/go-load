/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Colors } from "../style";

interface Props {
  title: string[];
  content: string[][];
}

const Table: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      <table css={table}>
        <thead css={th}>
          <tr css={trTitle}>
            {props.title.map((title) => (
              <th>{title}</th>
            ))}
          </tr>
        </thead>
        {props.content.map((rows, index) => (
          <tr css={tr} key={index}>
            {rows.map((column, colIndex) => (
              <td css={td(colIndex)} key={colIndex}>
                {column}
              </td>
            ))}
          </tr>
        ))}
      </table>
    </div>
  );
};

const container = css`
  width: 100%;
  border: 1px solid #e1e1e1;
  border-radius: 0.5rem;
  background-color: #e1e1e1;
  text-align: left;
  padding: 1rem 1rem 1rem 1rem;
`;

const table = css`
  width: 100%;
  border-collapse: collapse;
`;

const trTitle = css`
  background-color: #007d9c;
  border: 0.1rem solid #e1e1e1;
  height: 4rem;
  color: white;
  text-align: center;
`;

const th = css``;

const tr = css`
  height: 4rem;
  border-bottom: 0.1rem solid ${Colors.borderPrimary};
`;

const td = (index?: number) => css`
  width: 33%;
  ${index === 0
    ? "padding-left:2rem;"
    : "text-align:center;font-weight:bold;text-transform:uppercase;"}
`;

export default Table;
