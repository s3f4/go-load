import {makeReq} from '../api'

export interface TestConfig {
  id?: number;
  name: string;
  tests: Test[]
}
export interface Test {
  requestCount: number;
  goroutineCount: number;
  url: string;
  method: string;
  payload: string;
  expectedResponseCode: number;
  expectedResponseBody: string;
  transportConfig: TransportConfig,
}

export interface TransportConfig {
  DisableKeepAlives: boolean;
}

export const saveTests =
    async (testConfig: TestConfig) => {
  return await makeReq('/tests', 'POST', testConfig);
}

export const runTests = async (testConfig: TestConfig) => {
  return await makeReq('/tests/start', 'POST', testConfig)
}

export const listTests = async () => {
  return await makeReq('/tests', 'Get');
}