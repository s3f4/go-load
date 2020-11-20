export const toNum = (val: any): number => {
  let intVal = Number(val);

  if (isNaN(intVal)) {
    return 0;
  }

  return intVal;
};

export const defaultFormat = (): string => "DD.MM.YYYY hh:mm:ss";
export const preciseFormat = (): string => "hh:mm:ss.SSSS";
