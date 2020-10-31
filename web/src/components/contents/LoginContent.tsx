/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { MediaQuery } from "../style";

const initialLoginState = {
  email: "",
  password: "",
  error: "",
};

const LoginContent: React.FC = () => {
  const [login, setLogin] = useState(initialLoginState);
  const [isValid, setIsValid] = useState<any>({
    email: true,
    password: true,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    setLogin({
      ...login,
      [e.target.name]: e.target.value,
    });
  };

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const onLogin = () => {};

  const loginForm = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h2 css={formTitle}>Login</h2>
          <TextInput
            name="email"
            label={"E-Mail:"}
            type="text"
            onChange={handleChange}
            value={login.email}
            validate={{
              message: "Provide an valid email",
              validationFunction: validation("email"),
            }}
            isValid={isValid["email"]}
          />

          <TextInput
            name="password"
            label={"Password:"}
            type="password"
            onChange={handleChange}
            value={login.password}
            validate={{
              min: 6,
              max: 9,
              message: "Please provide a valid passowrd",
              validationFunction: validation("password"),
            }}
            isValid={isValid["password"]}
          />

          <div css={buttons}>
            <Button
              text="Login"
              onClick={onLogin}
              disabled={!isValid["email"] || !isValid["password"]}
            />
          </div>
        </div>
      </div>
    );
  };
  return loginForm();
};

const container = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const formDiv = css`
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  margin-bottom: 3rem;
  width: 95%;
  ${MediaQuery[1]} {
    width: 60%;
  }
`;

const formTitle = css`
  margin-top: 1rem;
  font-size: 2.3rem;
  text-decoration: none;
  text-align: center;
`;

const buttons = css`
  margin-top: 1rem;
  height: 6.5rem;
`;

export default LoginContent;
