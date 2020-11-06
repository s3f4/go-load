const URL = `http://${process.env.REACT_APP_API_BASE_URL}`;
// makeReq makes requests
export const makeReq = async (url: string, method?: any, body?: any) => {
  if (!url.startsWith("/")) {
    throw new Error("URL must start with / character");
  }

  const request = {
    url: `${URL}${url}`,
    config: {
      method: method ? method : "GET",
      headers: {
        Accept: "application/json",
      },
    },
  };

  if (body && request.config.method !== "GET") {
    (request.config as any).body = JSON.stringify(body);
  }

  return await fetch(request.url, request.config).then((response: Response) => {
    return new Promise((resolve, reject) => {
      if (response) {
        if (response.ok) {
          response.json().then((json) => resolve(json));
        } else {
          response
            .json()
            .then((json) => reject(json))
            .catch(() =>
              reject({
                status: response.status,
                message: response.statusText,
              }),
            );
        }
      }
    });
  });
};
