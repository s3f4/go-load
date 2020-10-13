import { makeReq } from "../api";

export interface RunTest {
  id?: number;
  test_id: number;
  start_time: Date;
  end_time: Date;
  passed: boolean;
}

export const GetRunTest = async (runTest: RunTest) => {
  return await makeReq(`/run_test/${runTest.id}`, "GET");
};

export const listRunTest = async () => {
  return await makeReq("/run_test", "Get");
};

export const deleteRunTest = async (runTest: RunTest) => {
  return await makeReq("/run_test", "DELETE", runTest);
};
