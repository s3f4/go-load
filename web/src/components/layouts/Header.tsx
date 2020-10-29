/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { ReactComponent as ReactLogo } from "../img/gopher.svg";
import { Link } from "react-router-dom";
import { FiServer, FiClipboard, FiActivity, FiMonitor } from "react-icons/fi";
import { MediaQuery } from "../style";

const headerLinkStyle = {
  width: "2rem",
  height: "1.8rem",
};

interface Props {}

const Header: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <div css={header}>
        <div css={logoDiv}>
          <a css={logoLink} href="/">
            <ReactLogo css={logo} />
            <span css={logoFont}>go-load</span>
          </a>
        </div>
        <div css={headerDiv}>
          <Link css={headerLink} to="/instances">
            <FiServer style={headerLinkStyle} />
            Instances
          </Link>
          <Link css={headerLink} to="/tests">
            <FiClipboard style={headerLinkStyle} />
            Tests
          </Link>
          <Link css={headerLink} to="/swarm">
            <FiMonitor style={headerLinkStyle} />
            Swarm
          </Link>
          <Link css={headerLink} to="/stats">
            <FiActivity style={headerLinkStyle} />
            Stats
          </Link>
        </div>
      </div>
    </React.Fragment>
  );
};

const header = () => {
  return css`
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    background-color: #007d9c;
    border-bottom: 0.1rem solid gray;
    & a {
      color: white;
      text-decoration: none;
    }

    ${MediaQuery[1]} {
      flex-direction: row;
      justify-content: space-between;
      height: 5rem;
      align-items: center;
    }
  `;
};

const logoDiv = css`
  height: 100%;
  width: 100%;
  ${MediaQuery[1]} {
    width: 30%;
  }
  padding: 0.4rem 1rem 0.5rem 1rem;
`;

const logo = css`
  height: 4rem;
  width: 4rem;
`;

const logoFont = css`
  text-align: center;
  ${MediaQuery[1]} {
    text-align: left;
  }
  height: 80%;
  font-size: 2.7rem;
`;

const logoLink = css`
  display: flex;
  justify-content: center;
  width: 100%;
  ${MediaQuery[1]} {
    justify-content: center;
  }
`;

const headerDiv = css`
  font-size: 2.3rem;
  display: flex;
  align-items: center;
  justify-content: space-around;
  height: 70%;
  width: 100%;

  & a {
    margin-left: 1.5rem;
  }

  ${MediaQuery[1]} {
    flex-direction: row;
    justify-content: center;
  }
`;

const headerLink = css`
  display: flex;
  flex-direction: column;
  align-items: center;

  ${MediaQuery[1]} {
    flex-direction: row;
  }
`;

export default Header;
