import { Worker } from "./entity/worker";

const URL = "http://localhost";

export const initInstances = async (item: any) => {
  return await fetch(`${URL}:3001/instances`, {
    method: "POST",
    headers: {
      Accept: "application/json",
    },
    body: JSON.stringify(item),
  })
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const destroy = async () => {
  return await fetch(`${URL}:3001/instances`, {
    method: "DELETE",
    headers: {
      Accept: "application/json",
    },
  })
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const listWorkers = async () => {
  return await fetch(`${URL}:3001/workers`, {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  })
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const stopWorker = async (worker: Worker) => {
  return await fetch(`${URL}:3001/workers`, {
    method: "POST",
    headers: {
      Accept: "application/json",
    },
    body: JSON.stringify(worker),
  })
    .then((response) => response.json())
    .catch((error) => console.log(error));
};

export const listAvailableRegions = async () => {
  return await fetch(`${URL}:3001/instances/regions`, {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  })
    .then((response) => response.json())
    .catch((error) => console.log(error));
};
