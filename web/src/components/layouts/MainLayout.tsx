/** @jsx jsx */
import React from "react";
import { Helmet } from "react-helmet";
import { jsx, Global, css } from "@emotion/core";
import { html } from "../style";
import Header from "./Header";
import Content from "./Content";
import Footer from "./Footer";

interface Props {
  title?: string;
  header?: string;
  content?: string;
  footer?: string;
}

const MainLayout: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <Helmet>
        <title>
          {props.title ? props.title : "Go-Load Load Testing Helper"}
        </title>
        <Global styles={html} />
      </Helmet>
      <div css={container}>
        <Header />
        <Content />
        <Footer />
      </div>
    </React.Fragment>
  );
};

const container = css`
  width: 1200px;
  height: auto;
  margin: 0 auto;
`;

export default MainLayout;
