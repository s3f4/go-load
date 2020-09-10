import {Worker} from './entity/worker';

const URL = 'http://localhost';

export const initInstances =
    async (item: any) => {
  const response = await fetch(`${URL}:3001/instances`, {
                     method: 'POST',
                     headers: {
                       Accept: 'application/json',
                     },
                     body: JSON.stringify(item)
                   })
                       .then(response => response.json())
                       .catch(error => console.log(error))

  return response;
}

export const destroy =
    async () => {
  const response = await fetch(`${URL}:3001/instances`, {
                     method: 'DELETE',
                     headers: {
                       Accept: 'application/json',
                     },
                   })

                       .then(response => response.json())
                       .catch(error => console.log(error))
  return response;
}

export const listWorkers =
    async () => {
  try {
    const response = await fetch(`${URL}:3001/workers`, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
      },
    })

    return response.json();
  } catch (err) {
    return {error: err};
  }
}

export const stopWorker = async (worker: Worker) => {
  try {
    const response = await fetch(`${URL}:3001/workers`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
      },
      body: JSON.stringify(worker)
    })
    return response.json();
  } catch (err) {
    return {error: err};
  }
}