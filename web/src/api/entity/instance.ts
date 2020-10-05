import {makeReq} from '../api';

export interface InstanceConfig {
  Configs: Instance[];
}

export interface Instance {
  InstanceCount: number;
  InstanceSize: string;
  Image: string;
  Region: string;
  MaxWorkingPeriod: number;
}

export const spinUp = async (instanceConfig: InstanceConfig) => {
  return await makeReq('/instances', 'POST', instanceConfig);
};

export const getInstanceInfo = async () => {
  return await makeReq(`/instances`);
};

export const destroyAll = async () => {
  return await makeReq(`/instances`, 'DELETE');
};

export const listAvailableRegions = async () => {
  return await makeReq(`/instances/regions`);
};

export const showAccount = async () => {
  return await makeReq(`/instances/account`);
};

export const showSwarmNodes = async () => {
  return await makeReq('/instances/swarm-nodes');
};
