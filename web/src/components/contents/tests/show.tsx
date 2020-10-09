/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { Link } from "react-router-dom";
import { listTests, runTests, Test, TestConfig } from "../../../api/entity/test_config";

interface Props {
  testConfg?: TestConfig;
}

const Show: React.FC<Props> = (props: Props) => {
  const [configs,setConfigs] = useState<TestConfig[]>();
  
  React.useEffect(()=>{
    listTests().then(response=>{
    setConfigs(response.data)
  }).catch(error=>console.log(error));
  },[])


  const run = (e: React.FormEvent) => {
    e.preventDefault();

    runTests(props.testConfg!)
      .then(() => {})
      .catch(() => {});
  };

  return (
    <div css={container}>
      <div css={leftColumn}>
        {configs?.map((config:TestConfig)=>
        <div key={config.id}>{config.name}adas</div>)}
        <hr />
        <Link to="/tests/create"> New Test Group</Link>
      </div>
      <div css={rightColumn}>
        {configs && configs[0].tests.map((test: Test) => {
          return (
            <div css={configCss} key={test.url}>
              URL : {test.url} - Method: {test.method} - Request Count:{" "}
              {test.requestCount}
            </div>
          );
        })}
      </div>
    </div>
  );
};

const container = css`
  display: flex;
  width: 100%;
  flex-direction: row;
`;

const leftColumn = css`
  background-color: #e3e3e3;
  width: 30%;
  padding: 2rem;
`;

const rightColumn = css`
  width: 70%;
  padding: 2rem;
`;

const configCss = css`
  width: 100%;
  height: 5rem;
  padding: 2rem 0;
  border-bottom: 1px solid black;
  text-align: left;
`;

export default Show;
