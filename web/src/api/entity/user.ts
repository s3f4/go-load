import { makeReq } from "../api";

export interface User {
  email: string;
  password: string;
}

export const signUp = async (user: User) => {
  return await makeReq("/user/signin", "POST", user);
};

export const signIn = async (user: User) => {
  return await makeReq("/user/signin", "POST", user);
};

export const signOut = async () => {
  return await makeReq("/user/signout");
};

export const refresh = async () => {
  return await makeReq("/user/_rt");
};
