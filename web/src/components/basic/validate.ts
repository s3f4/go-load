import { debug } from "console";
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

export interface ValidationResult {
  isValid: boolean;
  message?: string;
}

const parseValidationRules = (validate: string): Validate => {
  const validateObj: Validate = {};

  const rules = validate.split("|");
  rules.forEach((rule) => {
    const ruleArr = rule.split(":");
    if (ruleArr.length === 1) {
      validateObj[ruleArr[0]] = true;
    } else if (ruleArr.length === 2) {
      if (ruleArr[0].includes("min") || ruleArr[0].includes("max")) {
        validateObj[ruleArr[0]] = toNum(ruleArr[1].trim());
      } else {
        validateObj[ruleArr[0]] = ruleArr[1].trim();
      }
    }
  });

  return validateObj;
};

// validate is using to validation message
export const validate = (
  value: any,
  validationStr: string,
): ValidationResult => {
  const validate = parseValidationRules(validationStr);
  if (validate.min && toNum(value) < validate.min) {
    return { message: validate.message, isValid: false };
  }
  if (validate.max && toNum(value) > validate.max) {
    return { message: validate.message, isValid: false };
  }
  if (validate.minLength && value.length < validate.minLength) {
    return { message: validate.message, isValid: false };
  }
  if (validate.maxLength && value.length > validate.maxLength) {
    return { message: validate.message, isValid: false };
  }

  if (validate.url) {
    const regex = new RegExp(
      /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_.~#?&//=]*)/,
    );

    if (!value.match(regex)) {
      return { message: validate.message, isValid: false };
    }
  }

  if (validate.email) {
    const regex = new RegExp(
      /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
    );

    if (!value.match(regex)) {
      return { message: validate.message, isValid: false };
    }
  }

  return { message: undefined, isValid: true };
};
