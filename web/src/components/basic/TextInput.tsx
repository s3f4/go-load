/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Sizes, Colors } from "../style";
import BasicProps from "./basicProps";
import { validate, Validate } from "./validate";

interface Props extends BasicProps {
  label?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  validate?: Validate;
  isValid?: boolean;
}

const TextInput: React.FC<Props> = (props: Props) => {
  if (props.validate) validate(props.value, props.validate);

  return (
    <React.Fragment>
      <div css={inputDiv}>
        {props.label ? <label css={label}>{props.label}</label> : ""}
        <input
          name={props.name}
          css={textInput(props.isValid!)}
          type="text"
          value={props.value}
          onChange={props.onChange}
        />
      </div>
    </React.Fragment>
  );
};

const inputDiv = css`
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

export default TextInput;
