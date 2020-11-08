import { makeReq } from "../api";

export interface User {
  email: string;
  password: string;
}

let user: User | null = null;

export const setUser = (u: User) => {
  user = u;
};

export const getUser = (): User | null => user;

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
  return await makeReq("/auth/current_user");
};

export const refresh = async () => {
  return await makeReq("/auth/_rt");
};
