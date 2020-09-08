/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import Up from "../components/forms/init";

interface Props {}

const Instances: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout content={<Up />} />
    </React.Fragment>
  );
};

const instances = css``;

export default Instances;
