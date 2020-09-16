import { makeReq } from "../api";

export const spinUp = async (item: any) => {
  return await makeReq("/instances", "POST", item);
};

export const destroyAll = async () => {
  return await makeReq(`/instances`, "DELETE");
};

export const listAvailableRegions = async () => {
  return await makeReq(`/instances/regions`);
};
