/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

interface Props {
  content?: React.ReactNode;
}
const Content: React.FC<Props> = (props: Props) => {
  return <div css={content}>{props.content ? props.content : ""}</div>;
};

const content = css`
  width: 100%;
  min-height: 20rem;
`;

export default Content;
