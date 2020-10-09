/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../../basic/TextInput";
import Button from "../../basic/Button";
import { toNum } from "../../basic/helper";
import { BaseForm } from "../BaseForm";
import { Test, TransportConfig } from "../../../api/entity/test_config";
import SelectBox from "../../basic/SelectBox";

interface Props extends BaseForm {
  addNewTest: (test: Test) => (e: React.FormEvent) => void;
}

const CreateTest = (props: Props) => {
  const [requestCount, setRequestCount] = React.useState<number>(1);
  const [url, setUrl] = React.useState<string>("");
  const [method, setMethod] = React.useState<string>("");
  const [payload, setPayload] = React.useState<string>("");
  const [expectedResponseBody, setExpectedResponseBody] = React.useState<
    string
  >("");
  const [expectedResponseCode, setExpectedResponseCode] = React.useState<
    number
  >(0);
  const [goroutineCount, setGoroutineCount] = React.useState<number>(1);
  const [transportConfig, setTransportConfig] = React.useState<TransportConfig>(
    {
      DisableKeepAlives: true,
    },
  );
  const [isValid, setIsValid] = React.useState<any>({
    requestCount: true,
    url: true,
    method: true,
  });

  const validation = (name: string) => (value: boolean) =>
    setIsValid({
      ...isValid,
      [name]: value,
    });

  const test: Test = {
    requestCount,
    method,
    payload,
    goroutineCount,
    url,
    transportConfig,
    expectedResponseBody,
    expectedResponseCode,
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement> | any) => {
    if (!e.target && e.hasOwnProperty("value") && e.hasOwnProperty("label")) {
      setTransportConfig({
        ...transportConfig,
        DisableKeepAlives: e.value === "true",
      });
      return;
    }

    switch (e.target.name) {
      case "url":
        setUrl(e.target.value);
        break;
      case "requestCount":
        setRequestCount(toNum(e.target.value));
        break;
      case "method":
        setMethod(e.target.value);
        break;
      case "payload":
        setPayload(e.target.value);
        break;
      case "responseBody":
        setExpectedResponseBody(e.target.value);
        break;
      case "responseCode":
        setExpectedResponseCode(e.target.value);
        break;
      case "goroutineCount":
        let val = toNum(e.target.value);
        if (val <= 0) val = 1;
        setGoroutineCount(val);
        break;
    }
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
            value={url}
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
            value={requestCount}
            validate={{
              min: 1,
              validationFunction: validation("requestCount"),
              message: "Request must be greather than 0.",
            }}
            isValid={isValid["requestCount"]}
          />
          <TextInput
            onChange={handleChange}
            label="HTTP Method"
            name="method"
            value={method}
          />
          <TextInput
            onChange={handleChange}
            label="Request Payload"
            name="payload"
            value={payload}
          />
          <TextInput
            onChange={handleChange}
            label="Expected Response Code"
            name="responseCode"
            value={expectedResponseCode}
          />
          <TextInput
            onChange={handleChange}
            label="Expected Response Body"
            name="responseBody"
            value={expectedResponseBody}
          />

          <TextInput
            onChange={handleChange}
            label="Goroutine per worker (up to 10)"
            name="goroutineCount"
            value={goroutineCount}
          />
          <SelectBox
            name={"disableKeepAlives"}
            label={"Disable Keep-alive connections"}
            onChange={handleChange}
            options={[
              { value: "true", label: "True" },
              { value: "false", label: "False" },
            ]}
            value={transportConfig.DisableKeepAlives ? "true" : "false"}
          />
          <Button
            text="Add New Test"
            onClick={props.addNewTest(test)}
            disabled={!isValid["url"]}
          />
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

export default CreateTest;
