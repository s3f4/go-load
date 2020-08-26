/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes } from "../style";

const Footer = () => {
  return (
    <div css={footer}>
      <div css={footerContent}>2020 go-load</div>
    </div>
  );
};

const footer = css`
  border-top: 1px solid #007d9c;
  text-align: center;
  padding: 2rem 0;
`;

const footerContent = css`
  margin: 0 auto;
  width: 30%;
  font-size: ${Sizes.smallText};
  color: #007d9c;
  font-weight: 600;
`;

export default Footer;
