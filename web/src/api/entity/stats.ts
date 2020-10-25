import { makeReq } from "../api";

export interface Response {
  RunTestID: number;
  TotalTime: number;
  FirstByte: Date;
  FirstByteTime: number;
  DNSStart: Date;
  DNSDone: Date;
  DNSTime: number;
  TLSStart: Date;
  TLSDone: Date;
  TLSTime: number;
  ConnectStart: Date;
  ConnectDone: Date;
  ConnectTime: number;
  StatusCode: number;
  Body: string;
}

export const stats = async (runTestID: number) => {
  return await makeReq(`/run_test/${runTestID}`);
};
