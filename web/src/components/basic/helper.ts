export const toNum = (val: any): number => {
  let intVal = Number(val)

  if (isNaN(intVal)) {
    return 0;
  }

  return intVal;
}