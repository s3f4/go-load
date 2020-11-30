import { Query } from "../../components/basic/query";
import { makeReq, QueryString } from "../api";
import { RunTest } from "./runtest";
import { TestGroup } from "./test_group";

export interface Test {
  id?: number;
  name: string;
  test_group_id?: number;
  request_count: number;
  goroutine_count: number;
  url: string;
  method: string;
  payload: string;
  expected_response_code: number;
  expected_response_body: string;
  expected_first_byte_time: number;
  expected_connection_time: number;
  expected_dns_time: number;
  expected_tls_time: number;
  test_group?: TestGroup;
  transport_config: TransportConfig;
  headers?: Header[];
  run_tests?: RunTest[];
}

export interface Header {
  id?: number;
  key: string;
  value: string;
  is_request_header: boolean;
}

export interface TransportConfig {
  disable_keep_alives: boolean;
}

export const runTest = async (test: Test) => {
  return await makeReq(`/test/${test.id}/start`, "POST", test);
};

export const saveTest = async (test: Test) => {
  return await makeReq("/test", "POST", test);
};

export const getTest = async (testID: number) => {
  return await makeReq(`/test/${testID}`, "GET");
};

export const listTests = async (query?: Query) => {
  return await makeReq(`/test?${QueryString(query)}`);
};

export const listTestsOfTestGroup = (testID: number) => async (
  query?: Query,
) => {
  return await makeReq(`/test_group/${testID}/tests?${QueryString(query)}`);
};

export const deleteTest = async (test: Test) => {
  return await makeReq(`/test/${test.id}/`, "DELETE", test);
};

export const updateTest = async (test: Test) => {
  return await makeReq(`/test/${test.id}/`, "PUT", test);
};
