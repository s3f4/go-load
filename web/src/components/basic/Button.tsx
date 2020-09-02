/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders, Colors } from "../style";

interface Props {
  text: string;
  onClick: (e: any) => any;
}

const Button: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <button css={button} type="submit" onClick={props.onClick}>
        {props.text}
      </button>
    </React.Fragment>
  );
};

const button = css`
  border: ${Borders.buttonBorder};
  border-radius: ${Sizes.buttonBorderRadius};
  margin-top: 0.9rem;
  width: ${Sizes.buttonWidth};
  height: ${Sizes.inputHeight};
  color: ${Colors.buttonTextColor};
  background-color: ${Colors.buttonColor};
  font-weight: ${Sizes.buttonFontWeight};
  font-size: ${Sizes.buttonFontSize};
`;

export default Button;
