/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";

interface Props {}

const Workers = (props: Props) => {
  return (
    <React.Fragment>
      {/* <MainLayout /> */}
      <div css={workers}>
        <div css={workerCard}>Worker1</div>
        <div css={workerCard}>Worker2</div>
        <div css={workerCard}>Worker3</div>
        <div css={workerCard}>Worker4</div>
      </div>
    </React.Fragment>
  );
};

const workers = css`
  display: flex;
`;

const workerCard = css`
  width: 25rem;
  height: 25rem;
`;

export default Workers;
