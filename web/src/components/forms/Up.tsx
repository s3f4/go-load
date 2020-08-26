/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import TextInput from "../basic/TextInput";
import Button from "../basic/Button";

const Up = () => {
  return (
    <div>
      <label>Request Count</label>
      <TextInput type="text " name="requestCount" />
      <label>Instance Count</label>
      <TextInput type="text " name="requestCount" />
      <Button text="Up" />
    </div>
  );
};

export default Up;
