import { makeReq } from "../api";

export const getCsrfToken = async () => {
  try {
    const response = await makeReq("/form");
    return response.data;
  } catch (e) {
    return "";
  }
};
