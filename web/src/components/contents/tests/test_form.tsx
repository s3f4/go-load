/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import { toNum } from "../../basic/helper";
import { BaseForm } from "../../basic/BaseForm";
import SelectBox from "../../basic/SelectBox";
import { Header, Test } from "../../../api/entity/test";
import { TestGroup } from "../../../api/entity/test_group";

interface Props extends BaseForm {
  testGroup?: TestGroup;
  addTest?: (test: Test) => void;
  updateTest?: (test: Test) => void;
  setMessage?: () => void;
  test?: Test;
}

// type methodType = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

const initialTest: Test = {
  url: "",
  request_count: 1,
  method: "GET",
  expected_response_code: 0,
  expected_response_body: "",
  payload: "",
  goroutine_count: 1,
  headers: [],
  transport_config: { disable_keep_alives: true },
};

const TestForm = (props: Props) => {
  const [test, setTest] = React.useState<Test>(initialTest);
  const [isValid, setIsValid] = React.useState<any>({
    request_count: true,
    url: false,
    method: true,
    goroutine_count: true,
  });

  React.useEffect(() => {
    props.test && setTest(props.test);
  }, [props.test]);

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const validate = () => {
    let valid = true;
    Object.keys(isValid).forEach(function (key) {
      if (!isValid[key]) {
        valid = false;
      }
    });
    return valid;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    props.setMessage && props.setMessage();

    if (!e.target && e.hasOwnProperty("value") && e.hasOwnProperty("label")) {
      if (e.value === "true" || e.value === "false") {
        setTest({
          ...test,
          transport_config: { disable_keep_alives: e.value === "true" },
        });
        return;
      }
      setTest({
        ...test,
        method: e.value,
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

  const onHeaderHandle = (header: Header) => (
    e: React.ChangeEvent<HTMLInputElement>,
  ) => {
    e.preventDefault();
    header[e.target.name] = e.target.value;
    setTest({
      ...test,
      headers: test.headers!.map((h: Header) => {
        if (h.id === header.id) {
          return header;
        }
        return h;
      }),
    });
  };

  const onAddHeader = (e: React.FormEvent) => {
    e.preventDefault();
    const header: Header = {
      id: Date.now(),
      key: "",
      value: "",
    };
    setTest({
      ...test,
      headers: [...test.headers!, header],
    });
  };

  console.log(test);

  const formContent = () => {
    return (
      <div css={container}>
        <div css={formDiv}>
          <h3 css={formTitle}>
            {props.test ? "Update the Test" : "Create New Test"}
          </h3>
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
            name="request_count"
            value={test.request_count}
            validate={{
              min: 1,
              validationFunction: validation("request_count"),
              message: "Request must be greather than 0.",
            }}
            isValid={isValid["request_count"]}
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
            name="expected_response_code"
            value={test.expected_response_code}
          />
          <TextInput
            onChange={handleChange}
            label="Expected Response Body"
            name="expected_response_body"
            value={test.expected_response_body}
          />

          <TextInput
            onChange={handleChange}
            label="Goroutine per worker (up to 10)"
            name="goroutine_count"
            value={test.goroutine_count}
            validate={{
              min: 1,
              max: 10,
              message: "Goroutine count must be less than or equal to 10",
              validationFunction: validation("goroutine_count"),
            }}
            isValid={isValid["goroutine_count"]}
          />
          <SelectBox
            name={"disable_keep_alives"}
            label={"Disable Keep-alive connections"}
            onChange={handleChange}
            options={[
              { value: "true", label: "True" },
              { value: "false", label: "False" },
            ]}
            value={test.transport_config.disable_keep_alives ? "true" : "false"}
          />
          {test.headers &&
            test.headers.map((header: Header) => {
              return (
                <div key={header.id!} css={flex}>
                  <div css={headerDiv(true)}>
                    <TextInput
                      label="Header key"
                      name="key"
                      value={header.key}
                      onChange={onHeaderHandle(header)}
                    />
                  </div>
                  <div css={headerDiv(false)}>
                    <TextInput
                      label="Header value"
                      name="value"
                      value={header.value}
                      onChange={onHeaderHandle(header)}
                    />
                  </div>
                </div>
              );
            })}
          <Button text="Add New Header" onClick={onAddHeader} />

          {props.test ? (
            <Button
              text="Update"
              onClick={() => {
                props.updateTest?.(test);
                setTest(initialTest);
              }}
              disabled={
                !validate() ||
                typeof props.testGroup === "undefined" ||
                props.testGroup.name.length === 0
              }
            />
          ) : (
            <Button
              text="Add New Test"
              onClick={() => {
                props.addTest?.(test);
                setTest(initialTest);
              }}
              disabled={
                !validate() ||
                typeof props.testGroup === "undefined" ||
                props.testGroup.name.length === 0
              }
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

const flex = css`
  display: flex;
`;

const headerDiv = (right?: boolean) => css`
  width: 50%;
  ${right && typeof right !== "undefined" ? "margin-right: 3rem;" : ""}
`;

export default TestForm;
