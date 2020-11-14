export interface BaseForm {
  beforeSubmit?: () => void;
  afterSubmit?: () => void;
}

export const validateAll = (validateObj: any) => {
  let valid = true;
  Object.keys(validateObj).forEach(function (key) {
    if (!validateObj[key]) {
      valid = false;
    }
  });
  return valid;
};
