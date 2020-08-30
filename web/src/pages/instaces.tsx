/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";

interface Props {}

const Instances = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout />
    </React.Fragment>
  );
};

const instances = css``;

export default Instances;
