/** @jsx jsx */
import React from "react";
import { Helmet } from "react-helmet";
import { jsx, Global, css } from "@emotion/core";
import { html } from "../style";

interface Props {
  title?: string;
}

const Header: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <Helmet>
        <title>
          {props.title ? props.title : "Go-Load Load Testing Helper"}
        </title>
        <Global styles={html} />
      </Helmet>
      <div css={header}>
        <div css={headerDiv}>
          <a css={headerLink} href="/">
            go-load
          </a>
        </div>
        <div css={headerDiv}>content1</div>
        <div css={headerDiv}></div>
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

const headerDiv = css`
  width: 33.333%;
  height: auto;
`;

export default Header;
