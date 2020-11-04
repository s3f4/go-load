import { makeReq } from "../api";

export interface User {
  email: string;
  password: string;
}

export const signIn = async (user: User) => {
  return await makeReq("/signin", "POST", user);
};

export const signOut = async () => {
  return await makeReq("/signout");
};

export const refresh = async () => {
  return await makeReq("/_rt");
};
