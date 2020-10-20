import { makeReq } from "../api";
import { RunTest } from "./runtest";

export interface Test {
  id?: number;
  test_group_id?: number;
  request_count: number;
  goroutine_count: number;
  url: string;
  method: string;
  payload: string;
  expected_response_code: number;
  expected_response_body: string;
  transport_config: TransportConfig;
  headers?: Header[];
  run_tests?: RunTest[];
}

export interface Header {
  id?: number;
  key: string;
  value: string;
}

export interface TransportConfig {
  disable_keep_alives: boolean;
}

export const runTest = async (test: Test) => {
  return await makeReq("/test/run", "POST", test);
};

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
