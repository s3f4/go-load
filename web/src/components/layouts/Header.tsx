/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { ReactComponent as ReactLogo } from "../img/gopher.svg";
import { Link } from "react-router-dom";
import {
  FiServer,
  FiClipboard,
  FiActivity,
  FiMonitor,
  FiUser,
} from "react-icons/fi";
import { MediaQuery } from "../style";
import { useLocation, useHistory } from "react-router-dom";
import { getUserFromStorage, signOut } from "../../api/entity/user";
import { setToken, token } from "../../api/entity/jwt";

const headerIconStyle = {
  width: "2rem",
  height: "1.8rem",
};

interface Props {}

const Header: React.FC<Props> = (props: Props) => {
  const location = useLocation();
  const history = useHistory();
  const user = getUserFromStorage();

  const onSignOut = (e: React.FormEvent) => {
    e.preventDefault();
    signOut()
      .then((response) => {
        localStorage.removeItem("user");
        setToken("");
        history.push("/auth/signin");
        console.log(response);
      })
      .catch((error) => {
        localStorage.removeItem("user");
        setToken("");
        history.push("/auth/signin");
        console.log(error);
      });
  };
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
          <Link
            css={headerLink(location.pathname === "/instances")}
            to="/instances"
          >
            <FiServer style={headerIconStyle} />
            Instances
          </Link>
          <Link css={headerLink(location.pathname === "/tests")} to="/tests">
            <FiClipboard style={headerIconStyle} />
            Tests
          </Link>
          <Link css={headerLink(location.pathname === "/swarm")} to="/swarm">
            <FiMonitor style={headerIconStyle} />
            Swarm
          </Link>
          <Link css={headerLink(location.pathname === "/stats")} to="/stats">
            <FiActivity style={headerIconStyle} />
            Stats
          </Link>
        </div>
        <div>
          <div css={authLink}>
            {user && token ? (
              <Link
                onClick={onSignOut}
                css={headerLink(false)}
                to="/auth/signin"
              >
                <FiUser style={headerIconStyle} />
                Sign Out
              </Link>
            ) : (
              <Link css={headerLink(false)} to="/auth/signin">
                <FiUser style={headerIconStyle} />
                Sign In
              </Link>
            )}
          </div>
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
    justify-content: left;
  }
`;

const headerDiv = css`
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 100%;
  width: 100%;
  font-size: 2.3rem;

  & a {
    padding: 0.5rem 0 0.5rem 0;
    margin-left: 0.5rem;
    width: 100%;
  }

  ${MediaQuery[1]} {
    flex-direction: row;
    justify-content: center;

    & a {
      width: auto;
      padding: 0.5rem 1rem 0.5rem 1rem;
    }
  }
`;

const headerLink = (selected?: boolean) => css`
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  padding: 0rem 1rem 0rem 1rem;
  ${selected ? "background-color:#17a2b8" : ""};

  ${MediaQuery[1]} {
    flex-direction: row;
  }
`;

const authLink = css`
  min-width: 12rem;
  font-size: 2rem;
  font-weight: 500;
`;

export default Header;
