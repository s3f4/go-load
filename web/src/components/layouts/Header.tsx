/** @jsx jsx */
import React from "react";
import { Helmet } from "react-helmet";
import { jsx, Global, css } from "@emotion/core";
import { html } from "../style";

interface Props {}

const Header: React.FC<Props> = () => {
  return (
    <React.Fragment>
      <Helmet>
        <title>shello there</title>
        <Global styles={html} />
      </Helmet>
      <div css={header}>
        <a css={headerLink} href="/">
          go-load
        </a>
      </div>
    </React.Fragment>
  );
};

const header = css`
  background-color: #007d9c;
  width: 100%;
  height: 45px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-bottom: 1px solid gray;
  & a {
    color: white;
    text-decoration: none;
  }
`;

const headerLink = css`
  color: white;
  font-size: 25px;
`;

export default Header;
