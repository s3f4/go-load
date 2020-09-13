/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Colors, Sizes } from "../style";

interface SelectBoxData {
  value: string;
  text: string;
}

interface Props {
  label?: string;
  options: SelectBoxData[];
}

const SelectBox = (props: Props) => {
  return (
    <React.Fragment>
      {props.label ? <label css={label}>{props.label}</label> : ""}
      <select name="regions" css={selectBox}>
        {props.options.map((option) => (
          <option value={option.value}>{option.text}</option>
        ))}
      </select>
    </React.Fragment>
  );
};

const selectBox = css`
  height: ${Sizes.inputHeight};
  width: 100%;
  border-radius: ${Sizes.borderRadius1};
  font-size: ${Sizes.textInputFontSize};
`;

const label = css`
  font-size: ${Sizes.label};
  color: ${Colors.textPrimary};
  margin: 0.4rem;
`;

export default SelectBox;
