import { makeReq } from "../api";

export interface User {
  ID?: number;
  email?: string;
  password?: string;
}

export const setUserStorage = (user: User) => {
  localStorage.setItem("user", JSON.stringify(user));
};

export const getUserFromStorage = () => {
  return localStorage.getItem("user");
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
  return await makeReq("/user/current_user");
};

export const refresh = async () => {
  return await makeReq("/auth/_rt", "POST");
};
