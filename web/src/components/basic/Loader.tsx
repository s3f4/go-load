/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import LoaderSvg from "./LoaderSvg";
import { ReactComponent as LoaderComp } from "../img/loader.svg";

interface Props {
  message?: string;
  inlineLoading?: boolean;
  fill?: string;
}

const Loader: React.FC<Props> = (props: Props) => {
  return props.inlineLoading ? (
    <div css={inline}>
      <LoaderSvg width={"16"} height={"16"} fill={props.fill ?? "#fff"} />
    </div>
  ) : (
    <div css={loaderContainer}>
      <div css={loaderCss}>
        <LoaderComp />
        <span>{props.message ? props.message : ""}</span>
      </div>
    </div>
  );
};

const loaderContainer = css`
  display: block;
  position: relative;
  min-width: 10rem;
  min-height: 10rem;
  width: 100%;
  height: 100%;
`;

const loaderCss = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

const inline = css`
  display: inline-block;
  margin: 0.2rem 0.5rem;
`;

export default Loader;
