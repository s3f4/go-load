/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

interface Props {
  title: string[];
  content: string[][];
}

const Table: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      <table>
        <tr>
          {props.title.map((title) => (
            <th>{title}</th>
          ))}
        </tr>
        {props.content.map((rows, index) => (
          <tr key={index}>
            {rows.map((column, colIndex) => (
              <td key={colIndex}>{column}</td>
            ))}
          </tr>
        ))}
      </table>
    </div>
  );
};

const container = css`
  width: 100%;
  height: 4rem;
  padding: 1rem 0;
  border: 1px solid #e1e1e1;
  border-radius: 0.3rem;
  background-color: #e1e1e1;
  text-align: left;
`;

export default Table;
