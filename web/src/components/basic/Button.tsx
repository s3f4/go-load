/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders } from "../style";

interface Props {
  text: string;
}

const Button = (props: Props) => {
  return (
    <React.Fragment>
      <button css={button} type="submit">
        {props.text}
      </button>
    </React.Fragment>
  );
};

const button = css`
  border: ${Borders.buttonBorder};
  border-radius: ${Sizes.borderRadius1};
  margin: 0.2rem;
  min-width: 2rem;
  height: ${Sizes.inputHeight};
  color: white;
  background-color: #007d9c;
`;

export default Button;
