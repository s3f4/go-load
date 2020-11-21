/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

type messageType = "success" | "error" | "warning";
export interface IMessage {
  type: messageType;
  message: string;
}
interface Props {
  type: messageType;
  message: string;
}

const Message: React.FC<Props> = (props: Props) => {
  const messageContent = () => {
    return (
      <div css={container(props.type)}>
        <div css={message}>{props.message}</div>
      </div>
    );
  };
  return props.message && props.message.length > 0 ? (
    messageContent()
  ) : (
    <React.Fragment />
  );
};

const container = (type: string) => css`
  display: block;
  position: relative;
  border: 0.1rem solid white;
  border-radius: 0.5rem;
  text-align: center;
  font-size: 1.5rem;
  height: 8rem;
  width: 100%;
  padding: 4rem;
  background-color: ${type === "success"
    ? "green"
    : type === "error"
    ? "#dc3545"
    : "grey"};
  opacity: 0.8;
  margin: 1rem auto;
`;

const message = css`
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
`;

export default Message;
