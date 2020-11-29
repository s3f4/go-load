import { Query } from "../../components/basic/query";
import { makeReq, QueryString } from "../api";
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
  return await makeReq(`/test_group/${testGroup.id}`, "PUT", testGroup);
};

export const runTestGroup = async (testGroup: TestGroup) => {
  return await makeReq(`/test_group/${testGroup.id}/start`, "POST", testGroup);
};

export const listTestGroup = async (query?: Query) => {
  return await makeReq(`/test_group?${QueryString(query)}`);
};

export const deleteTestGroup = async (testGroup: TestGroup) => {
  return await makeReq(`/test_group/${testGroup.id}`, "DELETE", testGroup);
};
