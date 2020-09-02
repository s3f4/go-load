/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import { listWorkers } from "../api/api";

interface Props {}

const Workers: React.FC<Props> = (props: Props) => {
  const [workerContainers, setWorkerContainers] = React.useState<string>();

  React.useEffect(() => {
    listWorkers().then((response) => {
      setWorkerContainers(response);
      console.log(response);
    });
    return () => {};
  }, []);

  return (
    <React.Fragment>
      <MainLayout content={<WorkerContent content={workerContainers} />} />
    </React.Fragment>
  );
};

interface WorkerContentProps {
  content?: string;
}

const WorkerContent: React.FC<WorkerContentProps> = (
  props: WorkerContentProps,
) => {
  return (
    <div css={workers}>
      {JSON.stringify(props.content)}
      <div css={workerCard}>Worker1</div>
      <div css={workerCard}>Worker2</div>
      <div css={workerCard}>Worker3</div>
      <div css={workerCard}>Worker4</div>
    </div>
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
