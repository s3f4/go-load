import { makeReq } from "../api";

export interface Test {
  id?: number;
  testGroupId?: number;
  requestCount: number;
  goroutineCount: number;
  url: string;
  method: string;
  payload: string;
  expectedResponseCode: number;
  expectedResponseBody: string;
  transportConfig: TransportConfig;
}

export interface TransportConfig {
  DisableKeepAlives: boolean;
}

export const saveTest = async (test: Test) => {
  return await makeReq("/test", "POST", test);
};

export const getTest = async (test: Test) => {
  return await makeReq(`/test/${test.id}`, "GET");
};

export const deleteTest = async (test: Test) => {
  return await makeReq("/test", "DELETE", test);
};

export const updateTest = async (test: Test) => {
  return await makeReq("/test", "PUT", test);
};
