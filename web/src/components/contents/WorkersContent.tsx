/** @jsx jsx */
import { jsx, css } from "@emotion/core";
import React from "react";
import { Worker } from "../../api/entity/worker";
import Loader from "../basic/Loader";
import {
  card,
  cardContainer,
  cardContent,
  cardItem,
  cardTitle,
  MediaQuery,
} from "../style";
import Button from "../basic/Button";

interface Props {
  workers?: Worker[];
  handleStop: (worker: Worker) => any;
  loader: boolean;
}

const WorkersContent: React.FC<Props> = (props: Props) => {
  const workersDiv = () =>
    props.workers?.map((worker: Worker) => {
      if (worker.Names[0].includes("worker")) {
        return (
          <div css={card} key={worker.Id}>
            <h1 css={cardTitle}>{worker.Names[0].substr(1)}</h1>
            <div css={cardContent}>
              <div css={cardItem}>
                <b>ID:</b>
                <span>{worker.Id.substr(0, 7)} </span>
              </div>
              <div css={cardItem}>
                <b>Status:</b>
                <span>{worker.Status}</span>
              </div>
              <div css={cardItem}>
                <b>State:</b>
                <span>{worker.State}</span>
              </div>
              <div css={cardItem}>
                <Button
                  onClick={props.handleStop(worker)}
                  text={"Stop Container"}
                />
              </div>
            </div>
          </div>
        );
      }
    });

  return (
    <div>
      <div css={title}>Swarm Container List</div>
      <div css={cardContainer}>
        {!props.loader ? (
          workersDiv()
        ) : (
          <Loader message={"workers is loading..."} />
        )}
      </div>
    </div>
  );
};

const title = css`
  width: 70%;
  text-align: center;
  margin: 1rem auto;
  padding: 1rem;
  background-color: #efefef;

  ${MediaQuery[1]} {
    height: 4rem;
  }
`;

export default WorkersContent;
