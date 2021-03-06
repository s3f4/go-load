import { getCsrfToken } from "./entity/csrf";
import { setToken, token } from "./entity/jwt";

const URL = `http://${process.env.REACT_APP_API_BASE_URL}/api`;

export interface ServerResponse {
  status: boolean;
  data?: any;
  total?: any;
  message?: string;
  status_code: number;
}

// makeReq makes requests
export const makeReq = async (url: string, method?: any, body?: any) => {
  if (!url.startsWith("/")) {
    throw new Error("URL must start with / character");
  }

  const request = {
    url: `${URL}${url}`,
    config: {
      allowedOrigins: URL,
      credentials: "include" as RequestCredentials,
      method: method ? method : "GET",
      headers: {
        Accept: "application/json",
        Authorization: "",
      },
    },
  };

  if (token) {
    request.config.headers.Authorization = `Bearer ${token}`;
  }

  if (body && request.config.method !== "GET") {
    (request.config as any).body = JSON.stringify(body);
  }

  if (["POST", "PUT", "DELETE", "PATCH"].includes(request.config.method)) {
    const csrf: string = await getCsrfToken();
    request.config.headers["X-CSRF-Token"] = csrf;
  }

  return await fetch(request.url, request.config).then((response: Response) => {
    return new Promise<ServerResponse>((resolve, reject) => {
      response
        .json()
        .then((json: ServerResponse) => {
          json.status_code = response.status;
          if (response.ok) {
            resolve(json);
          } else {
            if (response.status === 401) {
              localStorage.clear();
              setToken("");
            }
            reject(json);
          }
        })
        .catch((error) =>
          reject({
            status: response.status,
            message: error,
            status_code: response.status,
          }),
        );
    });
  });
};

export const QueryString = (params: any) => {
  if (params) {
    return Object.keys(params)
      .map(function (key) {
        return key + "=" + params[key];
      })
      .join("&");
  }
  return "";
};
