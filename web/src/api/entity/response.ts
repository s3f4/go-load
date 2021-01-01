import { Query } from "../../components/basic/query";
import { makeReq, QueryString } from "../api";

export interface Response {
  run_test_id: number;
  worker_host_name: string;
  start_time: number;
  total_time: number;
  first_byte: Date;
  first_byte_time: number;
  dns_start: Date;
  dns_done: Date;
  dns_time: number;
  tls_start: Date;
  tls_done: Date;
  tls_time: number;
  connect_start: Date;
  connect_done: Date;
  connect_time: number;
  status_code: number;
  body: string;
  reasons: string;
  passed: boolean;
}

export const listResponses = (runTestID: number) => async (query?: Query) => {
  return await makeReq(`/run_test/${runTestID}/stats?${QueryString(query)}`);
};
