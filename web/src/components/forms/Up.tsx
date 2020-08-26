/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";

const Up = () => {
  return (
    <div css={formDiv}>
      <TextInput label={"Request Count"} type="text " name="requestCount" />

      <TextInput label={"Instance Count"} type="text " name="requestCount" />

      <TextInput label={"Region"} type="text " name="regions" id="region" />

      <TextInput
        label={"Max working period(minutes)"}
        type="text "
        name="maxWorkingPeriod"
      />
      <Button text="Up" />
    </div>
  );
};

const formDiv = css`
  margin: 0 auto;
  width: 50%;
`;

export default Up;
