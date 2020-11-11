import { currentUser, refresh } from "./user";

export let token = "";

export const setToken = (t: string) => {
  token = t;
};

export const getToken = () => {
  if (token === "") {
    currentUser()
      .then((response) => (token = response.data.token))
      .catch((error) => {
        if (error.status === 401) {
          refresh()
            .then((response) => console.log(response))
            .catch((error) => console.log(error));
        }
      });
  }
};
