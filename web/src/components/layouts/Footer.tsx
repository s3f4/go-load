/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

const Footer = () => {
  return (
    <div css={footer}>
      <div css={footerContent}>test</div>
    </div>
  );
};

const footer = css`
  border-top: 1px solid gray;
  text-align: center;
`;

const footerContent = css`
  margin: 0 auto;
  width: 30%;
`;

export default Footer;
