/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import { listWorkers } from "../api/api";
import { Worker as WorkerModel } from "../api/entity/worker";

interface Props {}

const Workers: React.FC<Props> = (props: Props) => {
  const [workerContainers, setWorkerContainers] = React.useState<WorkerModel[]>(
    [],
  );

  React.useEffect(() => {
    listWorkers()
      .then((response) => {
        setWorkerContainers(response.data.containers);
      })
      .catch((err) => console.log(err));
    return () => {};
  }, []);

  return (
    <React.Fragment>
      <MainLayout content={<WorkerContent workers={workerContainers} />} />
    </React.Fragment>
  );
};

interface WorkerContentProps {
  workers?: WorkerModel[];
}

const WorkerContent: React.FC<WorkerContentProps> = (
  props: WorkerContentProps,
) => {
  return (
    <div css={workers}>
      {props.workers
        ? props.workers.map((worker: WorkerModel) => {
            console.log(worker);
            return (
              <div css={workerCard} key={worker.Id}>
                {worker.Id.substr(0, 7)} <br />
                {worker.Status} <br />
                {worker.State}
                <br />
              </div>
            );
          })
        : ""}
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
