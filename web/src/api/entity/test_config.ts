import { makeReq } from "../api";

export interface TestConfig {
  id?: number;
  name: string;
  tests: Test[];
}
export interface Test {
  id?: number;
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

export const saveTests = async (testConfig: TestConfig) => {
  return await makeReq("/tests", "POST", testConfig);
};

export const updateTests = async (testConfig: TestConfig) => {
  return await makeReq("/tests", "PUT", testConfig);
};

export const runTests = async (testConfig: TestConfig) => {
  return await makeReq("/tests/start", "POST", testConfig);
};

export const listTests = async () => {
  return await makeReq("/tests", "Get");
};

export const deleteTestsReq = async (testConfig: TestConfig) => {
  return await makeReq("/tests", "DELETE", testConfig);
};

export const deleteTestReq = async (test: Test) => {
  return await makeReq("/tests/test", "DELETE", test);
};

export const updateTestReq = async (test: Test) => {
  return await makeReq("/tests/test", "PUT", test);
};
