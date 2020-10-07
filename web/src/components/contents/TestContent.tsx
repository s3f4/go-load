/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TestForm from "../forms/TestForm";
import { destroyAll } from "../../api/entity/instance";
import Button from "../basic/Button";

interface Props {}

const TestContent: React.FC<Props> = (props: Props) => {
  const destroyRequest = (e: any) => {
    e.preventDefault();
    destroyAll()
      .then((data) => console.log(data))
      .catch((error) => console.log(error));
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        InstanceInfo: instances
        <Button text="Destroy" onClick={destroyRequest} />
      </div>
      <TestForm instanceInfo={null} />
    </div>
  );
};

const container = css`
  display: flex;
  width: 100%;
  flex-direction: row;
`;

const leftColumn = css`
  background-color: #e3e3e3;
  width: 30%;
`;

export default TestContent;
