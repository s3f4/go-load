/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { Box, Sizes } from "../style";
import Button from "../basic/Button";

interface Props {
  configs: any;
  add: (config: any) => (e: any) => any;
  remove: (config: any) => (e: any) => any;
}

const InstanceConfigCards: React.FC<Props> = (props: Props) => {
  return (
    <div css={container}>
      {props.configs.map((config: any) => {
        return (
          <div css={configCss} key={config.region}>
            <div css={instanceTitle}>Region: {config.region}</div>
            <br />
            Instance Count: <b>{config.instanceCount}</b>
            <Button type={1} text="+" onClick={props.add(config)} />
            <Button type={1} text="-" onClick={props.remove(config)} />
          </div>
        );
      })}
    </div>
  );
};

const container = css`
  display: flex;
  width: 100%;
  flex-wrap: wrap;
`;

const configCss = css`
  background-color: #efefef;
  width: 15rem;
  height: 15rem;
  margin: 1rem 1rem;
  border: 1px solid black;
  text-align: center;
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1}
`;

const instanceTitle = css`
  background-color: #007d9c;
  color: white;
  width: 100%;
  height: 100;
  padding: 0.5rem;
  font-weight: bold;
`;

export default InstanceConfigCards;
