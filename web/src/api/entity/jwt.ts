import { refresh } from "./user";

export let token = "";

export const setToken = (t: string) => {
  token = t;
};

export const getToken = (silentRefresh?: boolean): Promise<String> => {
  return new Promise((resolve, reject) => {
    if (token === "" || silentRefresh) {
      refresh()
        .then((response) => {
          setToken(response.data.token);
          resolve(response.data.token);
        })
        .catch((error) => {
          reject(error);
        });
    } else {
      resolve(token);
    }
  });
};
