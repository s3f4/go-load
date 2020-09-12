/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders } from "../style";

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
  display: inline-block;
  border: ${Borders.border1};
  min-width: 10rem;
  min-height: 4rem;
  color: white;
  background-color: #007d9c;
  border-radius: ${Sizes.borderRadius1};
  padding: 1rem;
  margin: 0.5rem 0.5rem;
  font-size: 1.6rem;
  font-weight: 600;
  cursor: pointer;
  &:hover {
    background-color: #4caf50;
  }
`;

export default Button;
