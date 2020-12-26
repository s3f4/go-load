/** @jsx jsx */
import React, { Fragment, useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import { listTestsOfTestGroup, runTest, Test } from "../../../api/entity/test";
import Loader from "../../basic/Loader";
import Button, { ButtonColorType, ButtonType } from "../../basic/Button";
import { FiActivity, FiArrowRightCircle } from "react-icons/fi";
import { useHistory } from "react-router-dom";
import { TestGroup } from "../../../api/entity/test_group";
import { Borders, Colors, Sizes } from "../../style";
import {
  setItems,
  getItems,
  search,
  removeAll,
} from "../../basic/localStorage";
import Message, { IMessage } from "../../basic/Message";
import { findInAOO } from "../../basic/helper";
import { RunConfig } from "./ShowTests";
import { Query } from "../../basic/query";

interface Props {
  test?: Test;
  testGroup?: TestGroup;
  runConfigs: RunConfig[];
  setRunConfigs: (data: any) => void;
  clear: () => void;
}

const RunTests: React.FC<Props> = (props: Props) => {
  const [message, setMessage] = useState<IMessage>();
  const history = useHistory();

  useEffect(() => {
    props.setRunConfigs(getItems("run_configs") ?? []);

    if (!props.testGroup && props.test) {
      if (search("run_configs", [{ key: "test", value: props.test }]) === -1) {
        setItems("run_configs", [
          ...props.runConfigs,
          {
            test: props.test,
            loading: true,
            passed: false,
            started: true,
            finished: false,
            error: null,
          },
        ]);
        props.setRunConfigs(getItems("run_configs"));
        runWithConditions().catch((error) => {
          setItems("run_configs", []);
          setMessage({
            type: "error",
            message: error.message,
          });
        });
      }
    }
    if (props.testGroup && !props.test) {
      runTestGroup(props.testGroup.id!);
    }
  }, [props.test, props.testGroup]);

  const runWithConditions = async (): Promise<any> => {
    return new Promise((resolve, reject) => {
      const configs = getItems("run_configs") || [];
      configs.map((runConfig: RunConfig) => {
        if (
          runConfig.started &&
          runConfig.loading &&
          !runConfig.finished &&
          !runConfig.error
        ) {
          run(runConfig, runConfig.test)
            .then(() => {
              resolve(undefined);
            })
            .catch((error) => {
              reject(error);
            });
        }
      });
    });
  };

  const runTestGroup = (id: number, query?: Query, page?: number) => {
    let q: Query = {
        limit: 5,
        offset: 0,
        order: "d__id",
      },
      p: number = 1,
      total: number = 0;

    if (query) {
      q = query;
    }

    if (page) {
      p = page;
    }

    listTestsOfTestGroup(id)(q).then((response) => {
      total = response.data.total;
      response.data.data.map((test: Test) => {
        const oldItems = getItems("run_configs") || [];
        setItems("run_configs", [
          ...oldItems,
          {
            test,
            loading: true,
            passed: false,
            started: true,
            finished: false,
            error: null,
          },
        ]);
      });

      props.setRunConfigs(getItems("run_configs"));
      runWithConditions()
        .then(() => {
          if (p * q.limit <= total) {
            q = {
              ...q,
              offset: p * q.limit,
            };
            runTestGroup(id, q, ++p);
          }
        })
        .catch((error) => {
          setItems("run_configs", []);
          setMessage({
            type: "error",
            message: error.message,
          });
        });
    });
  };

  const run = async (runConfig: RunConfig, test: Test): Promise<any> => {
    try {
      const response = await runTest(test);
      runConfig.loading = false;
      runConfig.started = true;
      runConfig.finished = true;
      runConfig.passed = response.data.passed;
      const nRc = getItems("run_configs").map((r: RunConfig) => {
        if (runConfig.test.id === r.test.id) {
          return runConfig;
        }
        return r;
      });
      setItems("run_configs", nRc);
      props.setRunConfigs(nRc);
    } catch (error) {
      throw error;
    }
  };

  const clear = () => {
    removeAll("run_configs");
    props.clear();
    setMessage(undefined);
    history.push("/tests");
  };

  const isLoading = () => {
    return findInAOO(props.runConfigs, "loading");
  };

  return (
    <Fragment>
      <div css={container}>
        {message ? (
          <Message type={message.type} message={message.message} />
        ) : (
          props.runConfigs.map((runConfig: RunConfig) => {
            return (
              <div css={testLine} key={runConfig.test.id}>
                {runConfig.loading && (
                  <div css={item(5)}>
                    <Loader inlineLoading fill={Colors.passed} />
                  </div>
                )}
                {!runConfig.loading && (
                  <div css={item(5)}>
                    {runConfig.passed}
                    <FiArrowRightCircle
                      size="2.1rem"
                      color={
                        runConfig.passed ? Colors.passed : Colors.notPassed
                      }
                    />
                  </div>
                )}
                <div css={item(87)}>{runConfig.test.name}</div>
                {!runConfig.loading && (
                  <div css={item(7)}>
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
          })
        )}
        {props.runConfigs.length > 0 && (
          <div css={clearDiv}>
            <Button
              onClick={clear}
              text="Clear"
              disabled={!message && isLoading()}
            />
          </div>
        )}
      </div>
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

const clearDiv = css`
  margin: 0.5rem 0.5rem 0 0;
  width: 10rem;
  align-self: flex-end;
`;

export default RunTests;
