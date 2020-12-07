/** @jsx jsx */
import React, { useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import { listTests, runTest, Test } from "../../../api/entity/test";
import Loader from "../../basic/Loader";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { FiActivity } from "react-icons/fi";
import { useHistory } from "react-router-dom";
import { TestGroup } from "../../../api/entity/test_group";
import { Borders, Sizes } from "../../style";
import { setItem, getItem, search, removeItem } from "../../basic/localStorage";

interface Props {
  testGroup?: TestGroup;
  test?: Test;
}

interface RunConfig {
  test: Test;
  loading: boolean;
  passed: boolean;
}

const RunTests: React.FC<Props> = (props: Props) => {
  const [runConfigs, setRunConfigs] = useState<RunConfig[]>([]);
  const history = useHistory();

  useEffect(() => {
    const rc = getItem("run_configs");
    setRunConfigs(rc ?? []);

    if (props.test) {
      if (
        search("run_configs", [
          { key: "test", value: props.test },
          { key: "loading", value: true },
        ]) === -1
      ) {
        setItem("run_configs", [
          ...runConfigs,
          {
            test: props.test,
            loading: true,
            passed: true,
          },
        ]);
        setRunConfigs(getItem("run_configs"));
      }
    }
    if (props.testGroup) {
      removeItem("run_configs");
      const runConfigsList: RunConfig[] = [];
      listTests()
        .then((response) => {
          response.data.data.map((test: Test) => {
            runConfigsList.push({
              test,
              loading: true,
              passed: true,
            });
          });
          setItem("run_configs", [
            ...runConfigs,
            {
              test: props.test,
              loading: true,
              passed: true,
            },
          ]);
          setRunConfigs(getItem("run_configs"));
        })
        .catch((error) => {
          console.log(error);
        });
    }

    runConfigs.map((runConfig: RunConfig) => {
      console.log("runconfig run:", runConfig);
      //   runTest(runConfig.test)
      //     .then((response) => {
      //       console.log(response);
      //     })
      //     .catch((error) => {
      //       console.log(error);
      //     });
    });
  }, [props.test, props.testGroup]);

  return (
    <div css={container}>
      {runConfigs.map((runConfig: RunConfig) => {
        return (
          <div css={testLine} key={runConfig.test.id}>
            <div css={item(5)}>
              <Loader inlineLoading fill={"red"} />
            </div>
            <div css={item(90)}>{runConfig.test.name}</div>
            <div css={item(5)}>
              <Button
                colorType={ButtonColorType.info}
                type={ButtonType.iconButton}
                icon={<FiActivity />}
                onClick={(e: React.FormEvent) => {
                  e.preventDefault();
                  history.push(`/stats/${runConfig.test.id}`);
                }}
              />
            </div>
          </div>
        );
      })}
    </div>
  );
};

const container = css`
  display: flex;
  flex-direction: column;
  margin-bottom: 2rem;
`;

const testLine = css`
  display: flex;
  width: 100%;
  margin: 0.2rem auto;
  background-color: #e3e3e3;
  border: 1px solid black;
  min-height: 3rem;
  padding: 1rem;
  font-size: 1.7rem;
  border: ${Borders.border1};
  border-radius: ${Sizes.borderRadius1};
`;

const item = (width: number) => css`
  width: ${width}%;
`;

export default RunTests;
