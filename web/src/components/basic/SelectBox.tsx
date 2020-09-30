/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Colors, Sizes } from "../style";
import BasicProps from "./basicProps";
import Loader from "./Loader";

interface SelectBoxData {
  value: string;
  text: string;
}

interface Props extends BasicProps {
  options: SelectBoxData[];
  value: string;
  onChange?: (e: any) => void;
}

const SelectBox = (props: Props) => {
  return (
    <React.Fragment>
      {props.label ? <label css={label}>{props.label}</label> : ""}
      <select
        css={selectBox}
        onChange={props.onChange}
        name={props.name}
        value={props.value}
      >
        {props.options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.text}
          </option>
        ))}
      </select>
    </React.Fragment>
  );
};

const selectBox = css`
  height: ${Sizes.inputHeight};
  width: 100%;
  border: ${Borders.border1};
  border-radius: ${Sizes.borderRadius1};
  font-size: ${Sizes.textInputFontSize};
`;

const label = css`
  font-size: ${Sizes.label};
  color: ${Colors.textPrimary};
  margin: 0.4rem;
`;

export default SelectBox;
