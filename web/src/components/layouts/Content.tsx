/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import Up from "../forms/init";

interface Props {}
const Content: React.FC<Props> = () => {
  return (
    <div css={content}>
      <Up />
    </div>
  );
};

const content = css`
  width: 100%;
  min-height: 20rem;
  padding: 2rem 0;
`;

export default Content;
