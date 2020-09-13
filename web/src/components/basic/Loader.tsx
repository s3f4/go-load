/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import loaderSrc from "../img/loader.svg";

interface Props {
  message?: string;
}

const Loader = (props: Props) => {
  React.useEffect(() => {}, []);

  return (
    <div css={loaderContainer}>
      <div css={loaderCss}>
        <img alt={"loader"} src={loaderSrc} />
        <span>{props.message ? props.message : ""}</span>
      </div>
    </div>
  );
};

const loaderContainer = css`
  display: block;
  min-width: 10rem;
  min-height: 10rem;
  position: relative;
  width: 100%;
  height: 100%;
`;

const loaderCss = css`
  position: absolute;
  display: flex;
  flex-direction: column;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

export default Loader;
