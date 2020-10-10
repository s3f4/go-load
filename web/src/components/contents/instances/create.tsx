/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

interface Props {}

const Create: React.FC<Props> = (props: Props) => {
  return <div css={container}></div>;
};

const container = css``;

export default Create;
