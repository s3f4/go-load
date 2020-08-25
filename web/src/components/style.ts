import { css } from "@emotion/core";

export enum Colors {
  textPrimary = "#000",
  textSecondary = "#ddd",
  textTertiary = "#ddd",
  textQuaternary = "#ddd",

  buttonPrimary = "rgba(240, 243, 91, 0.863)",

  borderPrimary = "#ddd",
  borderLight = "#e5e5e5",
  borderInvalid = "red",

  mainBackgroundColor = "#fff",
  qBoxBackgroundColor = "#faf9f9",
  inputFocus = "#ddd",
}

export enum Sizes {
  inputHeight = "4rem",
  textAreaHeight = "20rem",
  fontSizePrimary = "1.6rem",
  borderRadius1 = "0.5rem",
  smallText = "1.1rem",
}

export class Borders {
  public static border1 = `0.1rem solid ${Colors.borderPrimary}`;
  public static textInputBorder = (valid: boolean) =>
    `0.1rem solid ${valid ? Colors.borderPrimary : "red"}`;
}

export enum Box {
  boxShadow1 = `-webkit-box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);
                -moz-box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);
                box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);`,
}

export const html = css`
  *,
  *::after,
  *::before {
    margin: 0;
    padding: 0;
    box-sizing: inherit;
  }
  html {
    font-size: 62.5%;
    font-family: "Roboto", Arial, Helvetica, sans-serif;
    background-color: red;
  }

  body {
    box-sizing: border-box;
    background-color: white;
    font-size: 1.7rem;
    color: ${Colors.textPrimary};
  }

  a {
    text-decoration: none;
    color: ${Colors.textPrimary};
  }

  textarea {
    resize: none;
  }

  h1 {
    text-align: center;
    font-size: 2rem;
  }
`;
