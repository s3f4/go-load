import { makeReq } from "../api";

export interface User {
  email: string;
  password: string;
}

export const login = async (user: User) => {
  return await makeReq("/login", "POST", user);
};

export const logout = async () => {
  return await makeReq("/logout");
};
