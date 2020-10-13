import { makeReq } from "../api";
import { Test } from "./test";

export interface TestGroup {
  id?: number;
  name: string;
  tests: Test[];
}

export const saveTestGroup = async (testGroup: TestGroup) => {
  return await makeReq("/test_group", "POST", testGroup);
};

export const updateTestGroup = async (testGroup: TestGroup) => {
  return await makeReq("/test_group", "PUT", testGroup);
};

export const runTestGroup = async (testGroup: TestGroup) => {
  return await makeReq("/test_group/start", "POST", testGroup);
};

export const listTestGroup = async () => {
  return await makeReq("/test_group", "Get");
};

export const deleteTestGroup = async (testGroup: TestGroup) => {
  return await makeReq("/test_group", "DELETE", testGroup);
};
