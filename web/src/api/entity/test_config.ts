import {makeReq} from '../api'

export interface TestConfig {
  ID?: number;
  Tests: Test[]
}
export interface Test {
  requestCount: number;
  goroutineCount: number;
  url: string;
  method: string;
  payload: string;
  expectedResponseCode: string;
  expectedResponseBody: string;
  transportConfig: TransportConfig,
}

export interface TransportConfig {
  DisableKeepAlives: boolean;
}

export const runTests = async (testConfig: TestConfig) => {
  return await makeReq('/test/run', 'POST', testConfig)
}