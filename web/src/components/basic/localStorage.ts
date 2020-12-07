interface Search {
  key: string;
  value: any;
}

export const setItem = (item: string, value: any, notJson?: boolean) => {
  localStorage.setItem(item, notJson ? value : JSON.stringify(value));
};

export const getItem = (item: string, notJson?: boolean) => {
  let value = localStorage.getItem(item);
  if (value) {
    return notJson ? value : JSON.parse(value);
  } else {
    return null;
  }
};

export const removeItem = (item: string) => {
  localStorage.removeItem(item);
};

export const search = (item: string, searchItem: string | Search[]): number => {
  const isString = typeof search === "string";

  const values = getItem(item, isString);
  if (values) {
    if (isString) {
      return values.indexOf(search);
    } else {
      return (values as any[]).findIndex((val: any) => {
        return searchConditions(searchItem as Search[], val);
      });
    }
  }
  return -1;
};

const searchConditions = (searchItem: Search[], val: any) => {
  let found = false;
  searchItem.map((condition: Search) => {
    found = jsonEqual(val[condition.key], condition.value);
  });
  return found;
};

function jsonEqual(a: any, b: any) {
  return JSON.stringify(a) === JSON.stringify(b);
}
