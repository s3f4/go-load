/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Borders, Colors, Sizes } from "../style";
import BasicProps from "./basicProps";
import Select from "react-select";
import { ValidationResult, validate } from "./validate";

interface SelectBoxData {
  value: string;
  label: string;
}

interface Props extends BasicProps {
  options: SelectBoxData[];
  value: string;
  onChange?: (e: any) => void;
  validate?: string;
  validation?: (name: string, value: boolean) => void;
}

const SelectBox = (props: Props) => {
  const [isValid, setIsValid] = useState<ValidationResult>();

  React.useEffect(() => {
    if (!props.validate) {
      setIsValid({ isValid: true });
    }
    if (props.validate && props.value !== undefined) {
      const validObj = validate(props.value, props.validate);
      setIsValid(validObj);
      props.validation?.(props.name, validObj.isValid);
    }
  }, [props.value, setIsValid]);

  return (
    <div css={container}>
      {props.label ? <label css={label}>{props.label}</label> : ""}
      <Select
        css={selectBox(isValid?.isValid!)}
        value={props.options.filter((option) => option.value === props.value)}
        styles={{
          control: (base) => ({
            ...base,
            outline: "none !important",
            border: Borders.textInputBorder(isValid?.isValid!) + " !important",
            fontSize: Sizes.textInputFontSize + " !important",
            boxShadow: "none !important",
          }),
        }}
        onChange={props.onChange}
        name={props.name}
        options={props.options}
      />
      {!isValid?.isValid && isValid?.message ? (
        <span css={validateMessage}>{isValid.message}</span>
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
