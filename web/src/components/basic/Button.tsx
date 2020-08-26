/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";

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

const button = css``;

export default Button;
