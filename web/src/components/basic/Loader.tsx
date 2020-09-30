/** @jsx jsx */
import React, { SVGProps } from "react";
import { jsx, css } from "@emotion/core";
import loaderSrc from "../img/loader.svg";
import LoaderSvg from "./LoaderSvg";

interface Props {
  message?: string;
  inlineLoading?: boolean;
}

const Loader = (props: Props) => {
  return props.inlineLoading ? (
    <div css={inline}>
      <LoaderSvg width={"16"} height={"16"} fill={"#fff"} />
    </div>
  ) : (
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

const inline = css`
  display: inline-block;
  margin-right: 0.5rem;
`;

export default Loader;
