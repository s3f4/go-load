export const toInt = (val: any): number => {
  let intVal = parseInt(val)

  if (isNaN(intVal)) {
    return 0;
  }

  return intVal;
}