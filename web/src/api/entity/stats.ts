import { makeReq } from "../api";

export interface Response {
  run_test_id: number;
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
}

export const stats = async (runTestID: number) => {
  return await makeReq(`/run_test/${runTestID}/stats`);
};
