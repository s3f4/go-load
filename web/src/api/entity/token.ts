import Cookies from "js-cookie";

export interface Token {
  token: string;
}

let token;

export const setToken = (t: Token) => {
  token = t;
};


