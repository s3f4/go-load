/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Sizes } from "../style";

interface Props {
  name: string;
  value?: any;
  type?: string;
}
const TextInput = (props: Props) => {
  return (
    <React.Fragment>
      <input css={textInput} type="text" value={props.value} />
    </React.Fragment>
  );
};

const textInput = css`
  border: ${Borders.textInputBorder(true)};
  height: ${Sizes.inputHeight};
  width: 100%;
  border-radius: ${Sizes.borderRadius1};
`;

export default TextInput;
