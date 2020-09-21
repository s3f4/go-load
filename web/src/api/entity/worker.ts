import {makeReq} from '../api';
import {RunConfig} from './run_config';
export interface Worker {
  Id: string;
  Status: string;
  State: string;
  Names: string[];
}

export const list = async () => {
  return await makeReq('/workers');
};

export const stop = async (worker: Worker) => {
  return await makeReq('/workers', 'POST', worker);
};

export const runWorkers = async (runConfig: RunConfig) => {
  return await makeReq('/workers/run', 'POST', runConfig)
}