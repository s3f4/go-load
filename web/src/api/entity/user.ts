import { makeReq } from "../api";

export interface User {
  email: string;
  password?: string;
  token?: string;
}

let user: User | null = null;

export const setUser = (u: User) => {
  user = u;
};

export const getUserObj = (): User | null => user;

export const getUser = (): User | null => {
  debugger;
  if (user != null) {
    return user;
  } else {
    currentUser()
      .then((response) => {
        console.log(response);
        user = response.data;
      })
      .catch((error) => {
        debugger;
        if (error.status_code == 401) {
          refresh()
            .then((response) => {
              user = response.data;
            })
            .catch((error) => {
              console.log(error);
            });
        }
      });
  }
  return user;
};

export const signUp = async (user: User) => {
  return await makeReq("/auth/signup", "POST", user);
};

export const signIn = async (user: User) => {
  return await makeReq("/auth/signin", "POST", user);
};

export const signOut = async () => {
  return await makeReq("/auth/signout");
};

export const currentUser = async () => {
  return await makeReq("/user/current_user", "POST");
};

export const refresh = async () => {
  return await makeReq("/auth/_rt", "POST");
};
