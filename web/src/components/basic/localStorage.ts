interface Condition {
  key: string;
  value: any;
}

export const addItem = (item: string, value: any, notJson?: boolean) => {
  const items = getItems(item);
  items.push(value);
  localStorage.setItem(items, notJson ? value : JSON.stringify(value));
};

export const setItems = (item: string, value: any, notJson?: boolean) => {
  localStorage.setItem(item, notJson ? value : JSON.stringify(value));
};

export const getItems = (item: string, notJson?: boolean) => {
  let value = localStorage.getItem(item);
  if (value) {
    return notJson ? value : JSON.parse(value);
  } else {
    return null;
  }
};

export const removeOne = (item: string, conditions: Condition[]) => {
  const items = getItems(item);
  const newItems = items.filter((i: any) => {
    return !filter(i, conditions);
  });
  setItems(item, newItems);
};

export const removeAll = (item: string) => {
  localStorage.removeItem(item);
};

export const search = (
  item: string,
  conditions: string | Condition[],
): number => {
  const isString = typeof search === "string";
  const values = getItems(item, isString);
  if (values) {
    if (isString) {
      return values.indexOf(search);
    } else {
      return (values as any[]).findIndex((val: any) => {
        return filter(val, conditions as Condition[]);
      });
    }
  }
  return -1;
};

const filter = (val: any, conditions: Condition[]) => {
  for (const condition of conditions) {
    // Search with two depth
    const keys = condition.key.split(".");
    if (keys.length === 2) {
      if (!jsonEqual(val[keys[0]][keys[1]], condition.value)) {
        return false;
      }
    } else {
      if (!jsonEqual(val[condition.key], condition.value)) {
        return false;
      }
    }
  }
  return true;
};

const jsonEqual = (a: any, b: any) => {
  return JSON.stringify(a) === JSON.stringify(b);
};
