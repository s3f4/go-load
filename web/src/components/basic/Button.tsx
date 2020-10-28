/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders } from "../style";
import Loader from "./Loader";
import { FiPlay } from "react-icons/fi";

enum ButtonType {
  small = 1,
  mid,
  big,
}
interface Props {
  text: string;
  onClick?: (e: any) => any;
  loading?: boolean;
  disabled?: boolean;
  type?: ButtonType;
}

const Button: React.FC<Props> = (props: Props) => {
  return (
    <React.Fragment>
      <button
        css={button(props.loading || props.disabled, props.type)}
        type="submit"
        onClick={props.onClick}
        disabled={props.loading || props.disabled}
      >
        <FiPlay />
        {props.loading ? <Loader inlineLoading={true} /> : ""}
        <div css={text}>{props.text}</div>
      </button>
    </React.Fragment>
  );
};

const button = (disabled: boolean | undefined, type?: ButtonType) => {
  let gen = `
    display: inline-block;
    border: ${Borders.border1};
    color: white;
    background-color: #007d9c;
    border-radius: ${Sizes.borderRadius1};
    margin: 0.5rem 0.5rem;
    cursor: pointer;
    padding: 1rem;
    font-weight: 600;
    ${disabled ? "opacity: 0.7;cursor:auto;" : ""}
    &:hover {
      ${disabled ? "" : "background-color: #4caf50;"}
    }
  `;
  switch (type) {
    case ButtonType.small:
      return css`
        ${gen}
        height: 3rem;
        min-width: 3rem;
        font-size: 1.4rem;
        padding: 0.3rem;
      `;
    case ButtonType.mid:
      return css`
        ${gen}
        height: 3.5rem;
        width: 6rem;
        font-size: 1.6rem;
      `;
    case ButtonType.big:
      return css`
        ${gen}
        min-width: 10rem;
        min-height: 4rem;
        font-size: 1.6rem;
      `;
    default:
      return css`
        ${gen}
        min-width: 10rem;
        min-height: 4rem;
        font-size: 1.6rem;
      `;
  }
};

const text = css`
  display: inline-block;
`;

export default Button;
