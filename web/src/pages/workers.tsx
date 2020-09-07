/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import MainLayout from "../components/layouts/MainLayout";
import { listWorkers, stopWorker } from "../api/api";
import { Worker as WorkerModel } from "../api/entity/worker";
import Loader from "../components/basic/Loader";
import { Box, Sizes, Borders } from "../components/style";

interface Props {}

const Workers: React.FC<Props> = (props: Props) => {
  const [workerContainers, setWorkerContainers] = React.useState<WorkerModel[]>(
    [],
  );
  const [loader, setLoader] = React.useState<boolean>(false);

  const handleStop = (worker: WorkerModel) => (e: any) => {
    e.preventDefault();
    setLoader(true);
    stopWorker(worker)
      .then(() => {
        const newWorkers = workerContainers.filter(
          (workerContainer) => workerContainer.Id !== worker.Id,
        );
        setLoader(false);
        setWorkerContainers(newWorkers);
      })
      .catch((error) => console.log(error));
  };

  React.useEffect(() => {
    setLoader(true);
    listWorkers()
      .then((response) => {
        setWorkerContainers(response.data.containers);
        setLoader(false);
      })
      .catch((err) => console.log(err));
    return () => {};
  }, []);

  return (
    <React.Fragment>
      <MainLayout
        content={
          <WorkerContent
            loader={loader}
            handleStop={handleStop}
            workers={workerContainers}
          />
        }
      />
    </React.Fragment>
  );
};

interface WorkerContentProps {
  workers?: WorkerModel[];
  handleStop: (worker: WorkerModel) => any;
  loader: boolean;
}

const WorkerContent: React.FC<WorkerContentProps> = (
  props: WorkerContentProps,
) => {
  const workersDiv = () =>
    props.workers?.map((worker: WorkerModel) => {
      if (worker.Names[0].startsWith("/worker")) {
        return (
          <div css={workerCard} key={worker.Id}>
            <h1 css={workerTitle}>{worker.Names[0].substr(1)}</h1>
            <br />
            {worker.Id.substr(0, 7)} <br />
            {worker.Status} <br />
            {worker.State}
            <br />
            <button css={btn} onClick={props.handleStop(worker)}>
              Stop Container
            </button>
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

const btn = css`
  border: ${Borders.border1};
  color: white;
  background-color: #007d9c;
  border-radius: ${Sizes.borderRadius1};
  padding: 1rem;
  margin: 0.5rem auto;
  font-size: 1.7rem;
  font-weight: 600;
`;

export default Workers;
