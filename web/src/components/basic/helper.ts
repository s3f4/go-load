export const toNum = (val: any): number => {
  let intVal = Number(val);

  if (isNaN(intVal)) {
    return 0;
  }

  return intVal;
};

export const defaultFormat = (): string => "DD.MM.YYYY hh:mm:ss";
export const preciseFormat = (): string => "DD.MM.YYYY hh:mm:ss.SSS";

export const findInAOO = (arr: object[], key: string) => {
  for (const item of arr) {
    if (item[key]) {
      return true;
    }
  }
  return false;
};

export const nanoToMilli = (nanoseconds: number): number =>
  nanoseconds / 1000000;
