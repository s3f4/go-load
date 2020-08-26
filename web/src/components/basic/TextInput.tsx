/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Sizes, Colors } from "../style";
import BasicProps from "./basicProps";

interface Props extends BasicProps {
  label?: string;
}

const TextInput: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <div css={inputDiv}>
        {props.label ? <label css={label}>{props.label}</label> : ""}
        <input css={textInput} type="text" value={props.value} />
      </div>
    </React.Fragment>
  );
};

const inputDiv = css`
  margin-top: 0.3rem;
`;

const label = css`
  font-size: ${Sizes.label};
  color: ${Colors.textPrimary};
  margin: 0.4rem;
`;

const textInput = css`
  border: ${Borders.textInputBorder(true)};
  height: ${Sizes.inputHeight};
  width: 100%;
  border-radius: ${Sizes.borderRadius1};
`;

export default TextInput;
