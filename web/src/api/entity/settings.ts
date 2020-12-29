import { makeReq } from "../api";
export interface Settings {
  Id: string;
  Setting: string;
  Value: string;
}

export const getSettings = async (setting: string) => {
  return await makeReq(`/settings/${setting}`);
};
