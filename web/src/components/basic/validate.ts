import {toNum} from './helper';

export interface Validate {
  maxLength?: number;
  minLength?: number;
  max?: number;
  min?: number;
  message?: string;
  url?: boolean;
  isValid: (condition: boolean) => void;
}

// validate is using to validation message
export const validate = (value: any, validate: Validate): void => {
  if (validate.min && toNum(value) < validate.min) {
    validate.isValid(false);
    return;
  }
  if (validate.max && toNum(value) > validate.max) {
    validate.isValid(false);
    return;
  }
  if (validate.minLength && value.length < validate.minLength) {
    validate.isValid(false);
    return;
  }
  if (validate.maxLength && value.length > validate.maxLength) {
    validate.isValid(false);
    return;
  }

  if (validate.url) {
    const regex = new RegExp(
        /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/);

    if (!value.match(regex)) {
      validate.isValid(false);
      return;
    }
  }

  validate.isValid(true);
};
