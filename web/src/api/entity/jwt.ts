import { refresh } from "./user";

export let token = "";

export const setToken = (t: string) => {
  token = t;
};

export const getToken = (): Promise<String> => {
  return new Promise((resolve, reject) => {
    if (token === "") {
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
