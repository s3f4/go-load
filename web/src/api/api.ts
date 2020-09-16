import { Worker } from "./entity/worker";

const URL = "http://localhost";

// fetcher
const makeReq = (url: any, method?: any, body?: any): Promise<Response> => {
  const request = {
    url: `${URL}:3001${url}`,
    config: {
      method: method ? method : "GET",
      headers: {
        Accept: "application/json",
      },
    },
  };

  if (body && request.config.method !== "GET") {
    (request.config as any).body = JSON.stringify(body);
  }

  return fetch(request.url, request.config);
};

export const initInstances = async (item: any) => {
  return await makeReq("/instances", "POST", item)
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const destroy = async () => {
  return await makeReq(`/instances`, "DELETE")
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const listWorkers = async () => {
  return await makeReq(`/workers`)
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const stopWorker = async (worker: Worker) => {
  return await makeReq(`/workers`, "POST", worker)
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const listAvailableRegions = async () => {
  return await makeReq(`/instances/regions`)
    .then((response) => response.json())
    .catch((error) => console.log(error));
};
