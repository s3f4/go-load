/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Colors, Sizes } from "../style";
import BasicProps from "./basicProps";
import Select from "react-select";
import { validate, Validate } from "./validate";

interface SelectBoxData {
  value: string;
  label: string;
}

interface Props extends BasicProps {
  options: SelectBoxData[];
  value: string;
  onChange?: (e: any) => void;
  validate?: Validate;
  isValid?: boolean;
}

const SelectBox = (props: Props) => {
  React.useEffect(() => {
    if (props.validate) validate(props.value, props.validate);
  }, [props.value]);

  return (
    <div css={container}>
      {props.label ? <label css={label}>{props.label}</label> : ""}
      <Select
        css={selectBox(props.isValid ?? true)}
        styles={{
          control: (base) => ({
            ...base,
            outline: "none !important",
            border:
              Borders.textInputBorder(props.isValid ?? true) + " !important",
            fontSize: Sizes.textInputFontSize + " !important",
            boxShadow: "none !important",
          }),
        }}
        onChange={props.onChange}
        name={props.name}
        options={props.options}
      />
      {!props.isValid && props.validate?.message ? (
        <span css={validateMessage}>{props.validate.message}</span>
      ) : (
        ""
      )}
    </div>
  );
};

const container = css`
  display: flex;
  flex-direction: column;
`;

const validateMessage = css`
  font-size: 1.3rem;
  color: red;
  text-align: right;
`;

const selectBox = (valid: boolean) => css`
  border: ${Borders.textInputBorder(valid)};
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
