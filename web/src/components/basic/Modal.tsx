/** @jsx jsx */
import React, { useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";

interface Props {
  show: boolean;
  onOpen?: () => any;
  onClose?: () => any;
}

const Modal: React.FC<Props> = (props: Props) => {
  const [showModal, setShowModal] = useState<boolean>(false);

  useEffect(() => {}, [showModal]);

  return (
    <div css={container(showModal)}>
      <div css={modalContent}>
        <span
          css={close}
          onClick={() => {
            setShowModal(false);
          }}
        >
          &times;
        </span>
        <p>some content</p>
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
  width: 80%;
  height: 80%;
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

export default Modal;
