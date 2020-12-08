/** @jsx jsx */
import React, { Fragment, useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import { listTests, runTest, Test } from "../../../api/entity/test";
import Loader from "../../basic/Loader";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { FiActivity } from "react-icons/fi";
import { useHistory } from "react-router-dom";
import { TestGroup } from "../../../api/entity/test_group";
import { Borders, Sizes } from "../../style";
import { setItem, getItem, search, removeItem } from "../../basic/localStorage";
import Message, { IMessage } from "../../basic/Message";
import { findInAOO } from "../../basic/helper";

interface Props {
  testGroup?: TestGroup;
  test?: Test;
}

interface RunConfig {
  test: Test;
  loading: boolean;
  passed: boolean;
  started: boolean;
  finished: boolean;
  error: any;
}

const RunTests: React.FC<Props> = (props: Props) => {
  const [runConfigs, setRunConfigs] = useState<RunConfig[]>([]);
  const [message, setMessage] = useState<IMessage>();
  const history = useHistory();

  useEffect(() => {
    const rc = getItem("run_configs");
    setRunConfigs(rc ?? []);

    if (props.test) {
      if (search("run_configs", [{ key: "test", value: props.test }]) === -1) {
        setItem("run_configs", [
          ...runConfigs,
          {
            test: props.test,
            loading: true,
            passed: true,
            started: false,
            error: null,
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
              passed: false,
              started: true,
              finished: false,
              error: null,
            });
          });
          setItem("run_configs", [
            ...runConfigs,
            {
              test: props.test,
              loading: true,
              passed: false,
              started: true,
              finished: false,
              error: null,
            },
          ]);
          setRunConfigs(getItem("run_configs"));
        })
        .catch((error) => {
          setMessage({
            type: "error",
            message: error,
          });
        });
    }

    runConfigs.map((runConfig: RunConfig) => {
      if (!runConfig.finished && !runConfig.error) {
        run(runConfig, runConfig.test);
      }
    });
  }, [props.test, props.testGroup]);

  const run = (runConfig: RunConfig, test: Test) => {
    runTest(test)
      .then((response) => {
        runConfig.loading = false;
        runConfig.started = true;
        runConfig.finished = true;
        runConfig.passed = response.data.passed;
        setRunConfigs(
          runConfigs.map((r: RunConfig) => {
            if (runConfig.test.id === r.test.id) {
              return runConfig;
            }
            return r;
          }),
        );
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const clear = () => {
    removeItem("run_configs");
    setRunConfigs([]);
  };

  const isLoading = () => {
    return findInAOO(runConfigs, "loading");
  };

  return (
    <Fragment>
      {message ? (
        <Message type={message.type} message={message.message} />
      ) : (
        <div css={container}>
          {runConfigs.map((runConfig: RunConfig) => {
            return (
              <div css={testLine} key={runConfig.test.id}>
                {runConfig.loading && (
                  <div css={item(5)}>
                    <Loader inlineLoading fill={"red"} />
                  </div>
                )}
                <div css={item(90)}>{runConfig.test.name}</div>
                {!runConfig.loading && (
                  <div css={item(10)}>
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
                )}
              </div>
            );
          })}
          {runConfigs.length > 0 && (
            <div css={clearButton}>
              <Button onClick={clear} text="Clear" disabled={isLoading()} />
            </div>
          )}
        </div>
      )}
    </Fragment>
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

const clearButton = css`
  margin-top: 0.5rem;
  width: 11.8%;
  align-self: flex-end;
`;

export default RunTests;
