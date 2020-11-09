const URL = `http://${process.env.REACT_APP_API_BASE_URL}`;
export let csrf: string | null = "";

export interface ServerResponse {
  status: boolean;
  data?: any;
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
        "X-CSRF-Token": csrf ? csrf : "",
      },
    },
  };

  const token = "bearer";
  if (token) {
    request.config.headers.Authorization = `Bearer ${token}`;
  }

  if (body && request.config.method !== "GET") {
    (request.config as any).body = JSON.stringify(body);
  }

  return await fetch(request.url, request.config).then((response: Response) => {
    return new Promise<ServerResponse>((resolve, reject) => {
      for (var pair of response.headers.entries()) {
        console.log(pair);
      }

      csrf = response.headers.get("X-CSRF-Token");

      console.log(csrf);
      response
        .json()
        .then((json: ServerResponse) => {
          json.status_code = response.status;
          if (response.ok) {
            resolve(json);
          } else {
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
