import { makeReq } from "../api";

export interface InstanceInfo {
  ID: string;
  InstanceCount: number;
  InstanceSize: string;
  Image: string;
  Region: string;
  MaxWorkingPeriod: number;
}

export const spinUp = async (item: any) => {
  return await makeReq("/instances", "POST", item);
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

export const showSwarmNodes = async () => {
  return await makeReq("/instances/swarm-nodes");
};
