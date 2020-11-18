/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Sizes, Borders } from "../style";
import Loader from "./Loader";

export enum ButtonType {
  small = 1,
  mid,
  big,
  iconButton,
  iconTextButton,
}

export const ButtonColorType = {
  primary: {
    color: "#fff",
    background_color: "#007d9c",
  },
  secondary: {
    color: "#fff",
    background_color: "#6c757d",
  },
  success: {
    color: "#fff",
    background_color: "#28a745",
  },
  danger: {
    color: "#fff",
    background_color: "#dc3545",
  },
  warning: {
    color: "#212529",
    background_color: "#ffc107",
  },
  info: {
    color: "#fff",
    background_color: "#17a2b8",
  },
  dark: {
    color: "#fff",
    background_color: "#343a40",
  },
  light: {
    color: "#212529",
    background_color: "#f8f9fa",
  },
};

interface ColorType {
  color: string;
  background_color: string;
}

interface Props {
  text?: string;
  onClick?: (e: any) => any;
  loading?: boolean;
  disabled?: boolean;
  type?: ButtonType;
  colorType?: ColorType;
  icon?: any;
}

const Button: React.FC<Props> = (props: Props) => {
  const renderContent = () => {
    if (!props.loading && props.type === ButtonType.iconTextButton) {
      return <span css={icon}>{props.icon}</span>;
    }

    if (props.loading) {
      return <Loader inlineLoading={true} />;
    }
  };

  return (
    <React.Fragment>
      <button
        css={button(
          props.loading || props.disabled,
          props.type,
          props.colorType ?? ButtonColorType.primary,
        )}
        type="submit"
        onClick={props.onClick}
        disabled={props.loading || props.disabled}
      >
        {renderContent()}
        {props.type === ButtonType.iconButton ? (
          props.icon
        ) : (
          <span css={text}>{props.text}</span>
        )}
      </button>
    </React.Fragment>
  );
};

const button = (
  disabled: boolean | undefined,
  type?: ButtonType,
  colorType?: any,
) => {
  let gen = `
    display: inline-block;
    border: ${Borders.border1};
    color: ${colorType.color};
    background-color: ${colorType.background_color};
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
    case ButtonType.iconButton:
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
    case ButtonType.iconTextButton:
      return css`
        ${gen}
        display:inline-flex;
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

const icon = css`
  display: inline-block;
  font-size: 1.6rem;
  margin: 0 0.4rem 0 0;
`;

export default React.memo(Button);
