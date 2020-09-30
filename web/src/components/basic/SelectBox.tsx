/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Colors, Sizes } from "../style";
import BasicProps from "./basicProps";
import Select from "react-select";

interface SelectBoxData {
  value: string;
  label: string;
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
      <Select
        css={selectBox}
        onChange={props.onChange}
        name={props.name}
        options={props.options}
      />
    </React.Fragment>
  );
};

const selectBox = css`
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
