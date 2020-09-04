/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import loaderSrc from "../img/loader.svg";

interface Props {}

const Loader = (props: Props) => {
  return (
    <div css={loaderContainer}>
      <img css={loaderCss} alt={"loader"} src={loaderSrc} />
    </div>
  );
};

const loaderContainer = css`
  position: relative;
  width: 100%;
  height: 100%;
`;
const loaderCss = css`
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

export default Loader;
