/** @jsx jsx */
import React, { MouseEventHandler } from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import { listWorkers, stopWorker } from "../api/api";
import { Worker as WorkerModel } from "../api/entity/worker";

interface Props {}

const Workers: React.FC<Props> = (props: Props) => {
  const [workerContainers, setWorkerContainers] = React.useState<WorkerModel[]>(
    [],
  );

  const handleStop = (worker: WorkerModel) => (e: any) => {
    e.preventDefault();
    stopWorker(worker)
      .then(() => {
        const newWorkers = workerContainers.filter(
          (workerContainer) => workerContainer.Id !== worker.Id,
        );
        console.log(newWorkers);
        setWorkerContainers(newWorkers);
      })
      .catch((error) => console.log(error));
  };

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
      <MainLayout
        content={
          <WorkerContent handleStop={handleStop} workers={workerContainers} />
        }
      />
    </React.Fragment>
  );
};

interface WorkerContentProps {
  workers?: WorkerModel[];
  handleStop: (worker: WorkerModel) => any;
}

const WorkerContent: React.FC<WorkerContentProps> = (
  props: WorkerContentProps,
) => {
  return (
    <div css={workers}>
      {props.workers
        ? props.workers.map((worker: WorkerModel) => {
            return (
              <div css={workerCard} key={worker.Id}>
                {worker.Names[0]}
                <br />
                {worker.Id.substr(0, 7)} <br />
                {worker.Status} <br />
                {worker.State}
                <br />
                <button onClick={props.handleStop(worker)}>
                  Stop Container
                </button>
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
