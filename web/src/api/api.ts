import { EROFS } from "constants";
import { Worker } from "./entity/worker";

const URL = "http://localhost";

// fetcher
const makeReq = async (url: any, method?: any, body?: any) => {
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

  return await fetch(request.url, request.config).then((response) => {
    debugger;
    if (response.ok) {
      return response.json();
    } else {
      return Promise.reject({
        status: response.status,
        statusText: response.statusText,
      });
    }
  });
};

export const initInstances = async (item: any) => {
  return await makeReq("/instances", "POST", item);
};

export const destroy = async () => {
  return await makeReq(`/instances`, "DELETE");
};

export const listWorkers = async () => {
  return await makeReq(`/workers`);
};

export const stopWorker = async (worker: Worker) => {
  return await makeReq(`/workers`, "POST", worker);
};

export const listAvailableRegions = async () => {
  return await makeReq(`/instances/regions`);
};
