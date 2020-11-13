import { toNum } from "./helper";

export interface Validate {
  maxLength?: number;
  minLength?: number;
  max?: number;
  min?: number;
  message?: string;
  url?: boolean;
  email?: boolean;
}

// validate is using to validation message
export const validate = (value: any, validate: Validate): boolean => {
  if (validate.min && toNum(value) < validate.min) {
    return false;
  }
  if (validate.max && toNum(value) > validate.max) {
    return false;
  }
  if (validate.minLength && value.length < validate.minLength) {
    return false;
  }
  if (validate.maxLength && value.length > validate.maxLength) {
    return false;
  }

  if (validate.url) {
    const regex = new RegExp(
      /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_.~#?&//=]*)/,
    );

    if (!value.match(regex)) {
      return false;
    }
  }

  if (validate.email) {
    const regex = new RegExp(
      /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
    );

    if (!value.match(regex)) {
      return false;
    }
  }

  return true;
};
