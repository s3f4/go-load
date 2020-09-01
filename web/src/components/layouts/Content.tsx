/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import Up from "../forms/init";

interface Props {
  content?: React.ReactNode;
}
const Content: React.FC<Props> = (props: Props) => {
  return <div css={content}>{props.content ? props.content : <Up />}</div>;
};

const content = css`
  width: 100%;
  min-height: 20rem;
  padding: 2rem 0;
`;

export default Content;
