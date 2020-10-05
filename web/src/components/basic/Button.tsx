/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders } from "../style";
import Loader from "./Loader";

interface Props {
  text: string;
  onClick: (e: any) => any;
  loading?: boolean;
  disabled?: boolean;
}

const Button: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <button
        css={button(props.loading || props.disabled)}
        type="submit"
        onClick={props.onClick}
        disabled={props.loading || props.disabled}
      >
        {props.loading ? <Loader inlineLoading={true} /> : ""}
        <div css={text}>{props.text}</div>
      </button>
    </React.Fragment>
  );
};

const button = (disabled: boolean | undefined) => css`
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
  ${disabled ? "opacity: 0.5;" : ""}
  &:hover {
    background-color: #4caf50;
  }
`;

const text = css`
  display: inline-block;
`;

export default Button;
