import { makeReq } from "../api";
export interface Service {
  Id: string;
  Status: string;
  State: string;
  Names: string[];
}

export const list = async () => {
  return await makeReq("/services");
};

export const stop = async (service: Service) => {
  return await makeReq("/services", "POST", service);
};
