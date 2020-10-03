/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import TestContent from "../components/contents/TestContent";

interface Props {}

const Test: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={TestContent} />
    </React.Fragment>
  );
};

export default Test;
