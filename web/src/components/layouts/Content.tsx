/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import Up from "../forms/Up";

const Content = () => {
  return (
    <div css={content}>
      <Up />
    </div>
  );
};

const content = css`
  width: 100%;
  min-height: 20rem;
`;

export default Content;
