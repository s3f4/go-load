/** @jsx jsx */
import React, { ReactNode } from "react";
import { jsx, css } from "@emotion/core";
import { IRTableColumn, IRTableRow, TableTitle } from "./RTable";

interface Props {
  show: boolean;
  setShow: (val: boolean) => any;
  title?: TableTitle[];
  row?: IRTableRow;
  content?: ReactNode;
}

const Modal: React.FC<Props> = (props: Props) => {
  return (
    <div css={container(props.show)}>
      <div css={modalContent}>
        <span css={close} onClick={() => props.setShow(false)}>
          &times;
        </span>
        <div css={contentDiv}>
          {props.row && props.row.allColumns
            ? props.row.allColumns.map((col: IRTableColumn, index: number) => {
                return (
                  <div css={rowDiv} key={index}>
                    <b>{props.title ? props.title[index].header : ""}:</b>
                    {col.content}
                  </div>
                );
              })
            : props.content
            ? props.content
            : ""}
        </div>
      </div>
    </div>
  );
};

const container = (show: boolean) => css`
  display: ${show ? "block" : "none"};
  position: fixed;
  z-index: 1;
  padding-top: 100px;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
`;

const modalContent = css`
  background-color: #fefefe;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  width: 60%;
  height: 80%;
  overflow: auto;
`;

const close = css`
  color: #aaaaaa;
  float: right;
  font-size: 28px;
  font-weight: bold;

  &:hover,
  &:focus {
    color: #000;
    text-decoration: none;
    cursor: pointer;
  }
`;

const rowDiv = css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 5rem;
  margin: 0 auto;
  min-height: 4rem;
  padding: 2rem 2rem 2rem 2rem;
  width: 90%;
  border-bottom: 1px solid #e3e3e3;
`;

const contentDiv = css`
  width: 96%;
`;

export default Modal;
