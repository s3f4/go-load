/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";

interface Props {}

const Stats: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <MainLayout />
    </React.Fragment>
  );
};

const stats = css``;

export default Stats;
