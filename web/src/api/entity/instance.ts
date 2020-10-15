import { makeReq } from "../api";

export interface InstanceConfig {
  configs: Instance[];
}

export interface Instance {
  count: number;
  size: string;
  image: string;
  region: string;
}

export const spinUp = async (instanceConfig: InstanceConfig) => {
  return await makeReq("/instances", "POST", instanceConfig);
};

export const getInstanceInfo = async () => {
  return await makeReq(`/instances`);
};

export const destroyAll = async () => {
  return await makeReq(`/instances`, "DELETE");
};

export const listAvailableRegions = async () => {
  return await makeReq(`/instances/regions`);
};

export const showAccount = async () => {
  return await makeReq(`/instances/account`);
};

export const showSwarmNodes = async () => {
  return await makeReq("/instances/swarm-nodes");
};
