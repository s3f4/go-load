/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Sizes, Colors } from "../style";
import BasicProps from "./basicProps";
import { validate, Validate } from "./validate";

interface Props extends BasicProps {
  label?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  validate?: Validate;
  isValid?: boolean;
  disabled?: boolean;
}

const TextInput: React.FC<Props> = (props: Props) => {
  const [valid, setValid] = useState<boolean>(false);

  const onChange = React.useCallback(() => {
    if (props.validate && props.value) {
      setValid(validate(props.value, props.validate));
    }

    return props.onChange;
  }, [props.value]);

  console.log(props.name);
  return (
    <React.Fragment>
      <div css={inputDiv}>
        {props.label ? <label css={label}>{props.label}</label> : ""}
        <input
          name={props.name}
          css={textInput(valid ?? true)}
          type={props.type ?? "text"}
          value={props.value}
          onChange={onChange()}
          disabled={props.disabled}
        />
        {!props.isValid && props.validate?.message ? (
          <span css={validateMessage}>{props.validate.message}</span>
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

const textInput = (valid: boolean) => css`
  border: ${Borders.textInputBorder(valid)};
  height: ${Sizes.inputHeight};
  width: 100%;
  border-radius: ${Sizes.borderRadius1};
  font-size: ${Sizes.textInputFontSize};
  padding: ${Sizes.textInputPadding};
`;

export default React.memo(TextInput);
