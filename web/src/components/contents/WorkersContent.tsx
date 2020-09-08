/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Worker } from "../../api/entity/worker";
import Loader from "../basic/Loader";
import { Box, Sizes } from "../style";
import Button from "../basic/Button";

interface Props {
  workers?: Worker[];
  handleStop: (worker: Worker) => any;
  loader: boolean;
}

const WorkersContent: React.FC<Props> = (props: Props) => {
  const workersDiv = () =>
    props.workers?.map((worker: Worker) => {
      if (worker.Names[0].startsWith("/worker")) {
        return (
          <div css={workerCard} key={worker.Id}>
            <h1 css={workerTitle}>{worker.Names[0].substr(1)}</h1>
            <br />
            {worker.Id.substr(0, 7)} <br />
            {worker.Status} <br />
            {worker.State}
            <br />
            <Button
              onClick={props.handleStop(worker)}
              text={"Stop Container"}
            />
          </div>
        );
      }
    });

  return <div css={workers}>{!props.loader ? workersDiv() : <Loader />}</div>;
};

const workers = css`
  display: flex;
  flex-wrap: wrap;
  height: 100%;
`;

const workerCard = css`
  width: 28rem;
  height: 25rem;
  margin: 1rem 1rem;
  border: 1px solid black;
  text-align: center;
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1}
`;

const workerTitle = css`
  background-color: #007d9c;
  color: white;
  width: 100%;
  height: 100;
  padding: 0.5rem;
`;

export default WorkersContent;
