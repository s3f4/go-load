const URL = "http://localhost";

// makeReq makes requests
export const makeReq = async (url: any, method?: any, body?: any) => {
  const request = {
    url: `${URL}:3001${url}`,
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

  return await fetch(request.url, request.config).then((response) => {
    if (response.ok) {
      return response.json();
    } else {
      return Promise.reject({
        status: response.status,
        statusText: response.statusText,
      });
    }
  });
};
