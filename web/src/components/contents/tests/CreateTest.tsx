/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import { toNum } from "../../basic/helper";
import { BaseForm } from "../../basic/BaseForm";
import { Test } from "../../../api/entity/test_config";
import SelectBox from "../../basic/SelectBox";

interface Props extends BaseForm {
  addNewTest: (test: Test) => void;
  updateNewTest: (test: Test) => void;
  setMessage?: () => void;
  test?: Test;
}

// type methodType = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

const initialTest: Test = {
  url: "",
  requestCount: 1,
  method: "GET",
  expectedResponseCode: 0,
  expectedResponseBody: "",
  payload: "",
  goroutineCount: 1,
  transportConfig: { DisableKeepAlives: true },
};

const TestForm = (props: Props) => {
  const [test, setTest] = React.useState<Test>(initialTest);
  const [isValid, setIsValid] = React.useState<any>({
    requestCount: true,
    url: false,
    method: false,
  });

  React.useEffect(() => {
    props.test && setTest(props.test);
  }, [props.test]);

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    if (props.setMessage) {
      props.setMessage();
    }
    if (!e.target && e.hasOwnProperty("value") && e.hasOwnProperty("label")) {
      if (e.value === "true" || e.value === "false") {
        setTest({
          ...test,
          ["transportConfig"]: { DisableKeepAlives: e.value === "true" },
        });
        return;
      }
      setTest({
        ...test,
        ["method"]: e.value,
      });
      return;
    }

    setTest({
      ...test,
      [e.target.name]:
        typeof test[e.target.name] === "number"
          ? toNum(e.target.value)
          : e.target.value,
    });
  };

  const formContent = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h3 css={formTitle}>Create Tests</h3>
          <TextInput
            onChange={handleChange}
            label="Target URL"
            name="url"
            value={test.url}
            validate={{
              url: true,
              message: "Please write a valid URL",
              validationFunction: validation("url"),
            }}
            isValid={isValid["url"]}
          />
          <TextInput
            onChange={handleChange}
            label="Total Request"
            name="requestCount"
            value={test.requestCount}
            validate={{
              min: 1,
              validationFunction: validation("requestCount"),
              message: "Request must be greather than 0.",
            }}
            isValid={isValid["requestCount"]}
          />
          <SelectBox
            name="method"
            label={"HTTP Method"}
            onChange={handleChange}
            options={[
              { value: "GET", label: "GET" },
              { value: "POST", label: "POST" },
              { value: "PUT", label: "PUT" },
              { value: "PATCH", label: "PATCH" },
              { value: "DELETE", label: "DELETE" },
            ]}
            value={test.method}
            validate={{
              minLength: 3,
              validationFunction: validation("method"),
              message: "Please select a method for HTTP requests.",
            }}
            isValid={isValid["method"]}
          />
          <TextInput
            onChange={handleChange}
            label="Request Payload"
            name="payload"
            value={test.payload}
          />
          <TextInput
            onChange={handleChange}
            label="Expected Response Code"
            name="expectedResponseCode"
            value={test.expectedResponseCode}
          />
          <TextInput
            onChange={handleChange}
            label="Expected Response Body"
            name="expectedResponseBody"
            value={test.expectedResponseBody}
          />

          <TextInput
            onChange={handleChange}
            label="Goroutine per worker (up to 10)"
            name="goroutineCount"
            value={test.goroutineCount}
          />
          <SelectBox
            name={"disableKeepAlives"}
            label={"Disable Keep-alive connections"}
            onChange={handleChange}
            options={[
              { value: "true", label: "True" },
              { value: "false", label: "False" },
            ]}
            value={test.transportConfig.DisableKeepAlives ? "true" : "false"}
          />
          {props.test ? (
            <Button
              text="Update"
              onClick={() => {
                props.updateNewTest(test);
                setTest(initialTest);
              }}
              disabled={!isValid["url"]}
            />
          ) : (
            <Button
              text="Add New Test"
              onClick={() => {
                props.addNewTest(test);
                setTest(initialTest);
              }}
              disabled={!isValid["url"]}
            />
          )}
        </div>
      </div>
    );
  };

  return formContent();
};

const container = css`
  display: block;
  width: 100%;
`;

const formDiv = css`
  margin: 0 auto;
  width: 80%;
  padding: 1rem 0 3rem 0;
`;

const formTitle = css`
  font-size: 2rem;
  text-decoration: none;
  text-align: center;
  padding: 0 0 1rem 0;
  border-bottom: 0.1rem solid #e3e3e3;
`;

export default TestForm;
