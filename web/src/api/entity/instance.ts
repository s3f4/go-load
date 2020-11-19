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

export interface InstanceTerra {
  id: string;
  name: string;
  created_at: Date;
  disk: number;
  image: string;
  ipv4_address: string;
  ipv4_address_private: string;
  memory: number;
  region: string;
  size: string;
  status: string;
}

export const spinUp = async (instanceConfig: InstanceConfig) => {
  return await makeReq("/instances", "POST", instanceConfig);
};

export const getInstanceInfo = async () => {
  return await makeReq(`/instances`);
};

export const getInstanceInfoFromTerraform = async () => {
  return await makeReq(`/instances/terraform`);
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
