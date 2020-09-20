import {makeReq} from '../api'

export interface Response {
  TotalTime: number;
  FirstByte: Date;
  DNSStart: Date;
  DNSDone: Date;
  TLSStart: Date;
  TLSDone: Date;
  ConnectStart: Date;
  ConnectDone: Date;
  StatusCode: number;
}

export const stats = async () => {
  return await makeReq('/stats')
}