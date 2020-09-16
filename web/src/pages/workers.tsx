import React from "react";
import MainLayout from "../components/layouts/MainLayout";
import { Worker, list, stop } from "../api/entity/worker";
import WorkersContent from "../components/contents/WorkersContent";

interface Props {}

const Workers: React.FC<Props> = (props: Props) => {
  const [workerContainers, setWorkerContainers] = React.useState<Worker[]>([]);
  const [loader, setLoader] = React.useState<boolean>(false);

  const handleStop = (worker: Worker) => (e: any) => {
    e.preventDefault();
    setLoader(true);
    stop(worker)
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
    list()
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
          <WorkersContent
            loader={loader}
            handleStop={handleStop}
            workers={workerContainers}
          />
        }
      />
    </React.Fragment>
  );
};

export default Workers;
