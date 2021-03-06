/** @jsx jsx */
import React, { ChangeEvent, useState } from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Sizes, Colors } from "../style";
import BasicProps from "./basicProps";
import { ValidationResult, validate } from "./validate";

interface Props extends BasicProps {
  label?: string;
  onChange?: (
    e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLTextAreaElement>,
  ) => void;
  validate?: string;
  disabled?: boolean;
  validation?: (name: string, value: boolean) => void;
  textarea?: boolean;
  rows?: number;
}

const TextInput: React.FC<Props> = (props: Props) => {
  const [isValid, setIsValid] = useState<ValidationResult>();

  React.useEffect(() => {
    if (!props.validate) {
      setIsValid({ isValid: true });
    }
    if (props.validate && props.value !== undefined) {
      const validObj = validate(props.value, props.validate);
      setIsValid(validObj);
      props.validation?.(props.name, validObj.isValid);
    }
  }, [props.value, setIsValid]);

  return (
    <React.Fragment>
      <div css={inputDiv}>
        {props.label ? <label css={label}>{props.label}</label> : ""}
        {props.textarea ? (
          <textarea
            name={props.name}
            css={textInput(isValid?.isValid!, true)}
            value={props.value}
            onChange={props.onChange}
            disabled={props.disabled}
            rows={props.rows ?? 4}
          />
        ) : (
          <input
            name={props.name}
            css={textInput(isValid?.isValid!)}
            type={props.type ?? "text"}
            value={props.value}
            onChange={props.onChange}
            disabled={props.disabled}
          />
        )}
        {!isValid?.isValid && isValid?.message ? (
          <span css={validateMessage}>{isValid.message}</span>
        ) : (
          ""
        )}
      </div>
    </React.Fragment>
  );
};

const validateMessage = css`
  font-size: 1.3rem;
  color: red;
  text-align: right;
`;

const inputDiv = css`
  display: flex;
  width: 100%;
  flex-direction: column;
  margin: 0.9rem auto;
  padding: 0.9rem auto;
`;

const label = css`
  font-size: ${Sizes.label};
  color: ${Colors.textPrimary};
  margin: 0.4rem;
`;

const textInput = (valid: boolean, textarea?: boolean) => css`
  border: ${Borders.textInputBorder(valid)};
  ${!textarea ? `height: Sizes.inputHeight` : ""};
  width: 100%;
  border-radius: ${Sizes.borderRadius1};
  font-size: ${Sizes.textInputFontSize};
  padding: ${Sizes.textInputPadding};
  ${textarea
    ? `resize: vertical;
      font-family: "Roboto", Arial, Helvetica, sans-serif;
      `
    : ""}
`;

export default TextInput;
