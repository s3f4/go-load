/** @jsx jsx */
import React, { Fragment, useState } from "react";
import { jsx, css } from "@emotion/core";
import { IRTableRow, TableTitle } from "./RTable";
import { DisableSelect } from "../style";
import Modal from "./Modal";

interface Props {
  mobile: boolean;
  row: IRTableRow;
  title: TableTitle[];
}

const RTableRow: React.FC<Props> = (props: Props) => {
  const [show, setShow] = useState<boolean>(false);

  const toggle = (val: boolean) => {
    setShow(val);
  };

  const content = () => {
    if (props.mobile) {
      return (
        <Fragment>
          <div css={mobileRow}>
            {props.row.columns.map((column, colIndex) => (
              <div css={mobileFlex} key={colIndex}>
                <b>{props.title[colIndex].header}</b>
                {column.content}
              </div>
            ))}
          </div>
        </Fragment>
      );
    } else {
      return (
        <div css={row(false)}>
          <Modal show={show} setShow={toggle} />
          {props.row.columns.map((column, colIndex) => (
            <div
              onClick={() => {
                if (colIndex !== props.row.columns.length - 1) {
                  toggle(true);
                }
              }}
              css={columnStyle(props.title[colIndex].width)}
              key={colIndex}
            >
              {column.content}
            </div>
          ))}
        </div>
      );
    }
  };

  return content();
};

const mobileRow = css`
  margin-top: 1rem;
  padding: 2rem;
  background-color: #e1e1e1;
`;
const mobileFlex = css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 5rem;
  min-height: 4rem;
`;

const row = (title?: boolean) => css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 4.5rem;
  border-bottom: 1px solid black;
  background-color: ${title ? "#007d9c" : "none"};
  color: ${title ? "white" : "none"};
  ${title ? DisableSelect : ""}
  cursor:pointer;
`;

const columnStyle = (width?: string, sortable?: boolean) => css`
  flex: 0 1 ${width ? width : "20rem"};
  padding: 1rem 1rem 1rem 1rem;
  text-align: center;
  ${sortable ? "cursor:pointer;" : ""}
`;

export default RTableRow;
