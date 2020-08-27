export enum Colors {
  textPrimary = '#3e4042',
  textSecondary = '#ddd',
  textTertiary = '#ddd',
  textQuaternary = '#ddd',

  buttonPrimary = 'rgba(240, 243, 91, 0.863)',

  borderPrimary = '#abb',
  borderLight = '#e5e5e5',
  borderInvalid = 'red',

  mainBackgroundColor = '#fff',
  qBoxBackgroundColor = '#faf9f9',
  inputFocus = '#ddd',

  buttonTextColor = '#fff',
  buttonColor = '#007d9c',
}

export enum Sizes {
  buttonFontWeight = 'bold',
  buttonWidth = '10rem',
  buttonFontSize = '1.2rem',

  inputHeight = '2.5rem',
  textAreaHeight = '20rem',
  fontSizePrimary = '1.6rem',
  borderRadius1 = '0.5rem',
  buttonBorderRadius = '0.5rem',
  smallText = '1.4rem',

  textInputFontSize = '1.4rem',
  label = '1.4rem',
}

export class Borders {
  public static buttonBorder = `0.1rem solid ${Colors.borderPrimary}`;
  public static border1 = `0.1rem solid ${Colors.borderPrimary}`;
  public static textInputBorder = (valid: boolean) =>
      `0.1rem solid ${valid ? Colors.borderPrimary : 'red'}`;
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
`;
