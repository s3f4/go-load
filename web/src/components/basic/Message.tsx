/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

interface Props {
  type: "success" | "error" | "warning";
  message: string;
}

const Message: React.FC<Props> = (props: Props) => {
  return (
    <div css={container(props.type)}>
      <div css={message}>{props.message}</div>
    </div>
  );
};

const container = (type: string) => css`
  display: block;
  position: relative;
  border: 0.1rem solid white;
  border-radius: 0.3rem;
  text-align: center;
  font-size: 1.5rem;
  height: 8rem;
  width: 100%;
  padding: 4rem;
  background-color: ${type === "success"
    ? "green"
    : type === "error"
    ? "#dc3545"
    : "yellow"};
`;

const message = css`
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
`;

export default Message;
