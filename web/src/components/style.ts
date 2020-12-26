import { css } from "@emotion/core";

export const MediaQuery = [640, 768, 1024, 1280].map(
  (bp) => `@media (min-width: ${bp}px)`,
);

export enum Colors {
  textPrimary = "#3e4042",
  textSecondary = "#ddd",
  textTertiary = "#ddd",
  textQuaternary = "#ddd",

  buttonPrimary = "rgba(240, 243, 91, 0.863)",

  borderPrimary = "#abb",
  borderLight = "#e5e5e5",
  borderInvalid = "red",

  mainBackgroundColor = "#fff",
  qBoxBackgroundColor = "#faf9f9",
  inputFocus = "#ddd",

  buttonTextColor = "#fff",
  buttonColor = "#007d9c",
  passed = "#87b666",
  notPassed = "#ff6961",
}

export enum Sizes {
  buttonFontWeight = "bold",
  buttonWidth = "10rem",
  buttonFontSize = "1.2rem",

  inputHeight = "4rem",
  textInputPadding = "0.8rem 1rem",
  textAreaHeight = "20rem",
  fontSizePrimary = "1.6rem",
  borderRadius1 = "0.5rem",
  buttonBorderRadius = "0.5rem",
  smallText = "1.4rem",

  textInputFontSize = "1.4rem",
  label = "1.4rem",
}

export class Borders {
  public static buttonBorder = `0.1rem solid ${Colors.borderPrimary}`;
  public static border1 = `0.1rem solid ${Colors.borderPrimary}`;
  public static textInputBorder = (valid: boolean) =>
    `0.2rem solid ${valid ? Colors.borderPrimary : "red"}`;
}

export enum Box {
  boxShadow1 = `-webkit-box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);
                -moz-box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);
                box-shadow: 0.3rem 0.3rem 0.2rem 0px rgba(0,0,0,0.25);`,
}

export const html = `
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
  textarea, select, input, button { outline: none; }
`;

export const leftContent = (selected?: boolean) => css`
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 3rem;
  border-bottom: ${Borders.border1};
  padding: 1rem;
  cursor: pointer;
  ${selected ? "background-color:#e3e3c3" : ""};
  &:first-of-type {
    border-top: ${Borders.border1};
  }
`;

export const cardContainer = css`
  display: flex;
  flex-wrap: wrap;
  height: 100%;
`;

export const card = css`
  background-color: #efefef;
  width: 28rem;
  height: 25rem;
  margin: 1rem 1rem;
  border: 1px solid black;
  text-align: center;
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1}
`;

export const cardTitle = css`
  background-color: #007d9c;
  color: white;
  width: 100%;
  height: 100;
  padding: 0.5rem;
`;

export const cardContent = css`
  display: flex;
  flex-direction: column;
  padding: 1rem;
`;

export const cardItem = css`
  display: flex;
  justify-content: space-between;
`;

export const leftColumn = css`
  background-color: #e3e3e3;
  width: 100%;
  padding: 2rem;

  ${MediaQuery[1]} {
    width: 25%;
    min-height: 50rem;
  }
`;

export const rightColumn = css`
  width: 100%;
  ${MediaQuery[1]} {
    width: 75%;
    min-height: 50rem;
    padding: 2rem;
  }
`;

export const DisableSelect = `
  -webkit-user-select: none;  
  -moz-user-select: none;    
  -ms-user-select: none;      
  user-select: none;
`;
