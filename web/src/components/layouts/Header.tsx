/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { ReactComponent as ReactLogo } from "../img/gopher.svg";

interface Props {}

const Header: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <div css={header}>
        <div css={logoDiv}>
          <a css={headerLink} href="/">
            <ReactLogo css={logo} />
            <span css={logoFont}>go-load</span>
          </a>
        </div>
        <div css={headerDiv}>
          <a href="content">Instances</a>
          <a href="content">Workers</a>
          <a href="content">Stats</a>
        </div>
      </div>
    </React.Fragment>
  );
};

const header = css`
  background-color: #007d9c;
  width: 100%;
  height: 4.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 0.1rem solid gray;
  & a {
    color: white;
    text-decoration: none;
  }
`;

const logoDiv = css`
  height: 100%;
  width: 30%;
  padding: 0.2rem 1rem 0.5rem 1rem;
`;

const logo = css`
  height: 80%;
  width: 20%;
`;

const logoFont = css`
  height: 70%;
  width: 80%;
  font-size: 2.5rem;
`;

const headerLink = css`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
`;

const headerDiv = css`
  width: 70%;
  height: 70%;
  font-size: 2rem;
  display: flex;
  align-items: center;
  & a {
    margin-left: 1.5rem;
  }
`;

export default Header;
