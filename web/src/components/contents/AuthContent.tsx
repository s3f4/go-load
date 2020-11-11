/** @jsx jsx */
import React, { useState, useEffect } from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";
import { MediaQuery } from "../style";
import { setUserStorage, signIn, signUp } from "../../api/entity/user";
import Message from "../basic/Message";
import { Link, useHistory } from "react-router-dom";
import { setToken } from "../../api/entity/jwt";

const initialUserState = {
  email: "",
  password: "",
};

interface Props {
  type: string;
}

const AuthContent: React.FC<Props> = (props: Props) => {
  const [user, setUser] = useState(initialUserState);
  const [error, setError] = useState<string>("");
  const [isValid, setIsValid] = useState<any>({
    email: true,
    password: true,
  });
  const history = useHistory();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser({
      ...user,
      [e.target.name]: e.target.value,
    });
  };

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const onSignIn = () => {
    signIn(user)
      .then((response) => {
        setToken(response.data.token);
        setUserStorage(response.data);
        history.push("/instances");
      })
      .catch((error) => {
        setError(error.message);
      });
  };

  const onSignUp = () => {
    signUp(user)
      .then((response) => {
        setToken(response.data.token);
        setUserStorage(response.data);
        history.push("/instances");
      })
      .catch((error) => {
        setError(error.message);
        console.log(error);
      });
  };

  const button = () => {
    let buttonText = "Sign In";
    let buttonFunc = onSignIn;
    let text = (
      <React.Fragment>
        If you don't have an account <Link to="/auth/signup">Sign Up</Link>
      </React.Fragment>
    );

    if (props.type === "signup") {
      buttonText = "Sign Up";
      buttonFunc = onSignUp;
      text = (
        <React.Fragment>
          If you already have an account <Link to="/auth/signin">Sign In</Link>
        </React.Fragment>
      );
    }
    return (
      <React.Fragment>
        <Button
          text={buttonText}
          onClick={buttonFunc}
          disabled={!isValid["email"] || !isValid["password"]}
        />
        <span css={buttonRightText}>{text}</span>
      </React.Fragment>
    );
  };

  const signInForm = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h2 css={formTitle}>
            {props.type === "signin" ? "Sign In" : "Sign Up"}
          </h2>
          <Message type="error" message={error} />
          <TextInput
            name="email"
            label={"E-Mail:"}
            type="text"
            onChange={handleChange}
            value={user.email}
            validate={{
              email: true,
              message: "Provide a valid email",
              validationFunction: validation("email"),
            }}
            isValid={isValid["email"]}
          />

          <TextInput
            name="password"
            label={"Password:"}
            type="password"
            onChange={handleChange}
            value={user.password}
            validate={{
              minLength: 4,
              maxLength: 9,
              message: "Please provide a valid passowrd",
              validationFunction: validation("password"),
            }}
            isValid={isValid["password"]}
          />

          <div css={buttons}>{button()}</div>
        </div>
      </div>
    );
  };
  return signInForm();
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

const buttonRightText = css`
  font-size: 1.5rem;
`;

export default AuthContent;
