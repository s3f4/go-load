/** @jsx jsx */
import React, { Fragment, useState } from "react";
import { jsx, css } from "@emotion/core";
import { stats, Response } from "../../api/entity/stats";
import moment from "moment";
import { Line } from "react-chartjs-2";
import { defaultFormat, preciseFormat } from "../basic/helper";
import { getTest, Test } from "../../api/entity/test";
import { Borders, MediaQuery } from "../style";
import { RunTest } from "../../api/entity/runtest";
import { FiArrowRightCircle } from "react-icons/fi";
import RTable from "../basic/RTable";

interface Props {
  testID: number;
}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [responses, setResponses] = useState<Response[]>([]);
  const [test, setTest] = useState<Test>();
  const [selectedRunTest, setSelectedRunTest] = useState<RunTest>();

  React.useEffect(() => {
    getTest(props.testID)
      .then((response) => {
        setTest(response.data);
      })
      .catch((error) => console.log(error));
  }, [props.testID]);

  const listResponses = () => {
    if (selectedRunTest) {
      stats(selectedRunTest.id!)
        .then((response) => {
          setResponses(response.data);
        })
        .catch((error) => console.log(error));
    }
  };

  const byteSize = (str: string) => new Blob([str]).size;

  const chartData = () => {
    const datum: any[] = [];
    const labels: any[] = [];
    responses.map((response: Response, index: number) => {
      datum.push(response.dns_time / 100000);
      labels.push("request_" + index);
      return null;
    });
    return {
      datum,
      labels,
    };
  };

  const graph = () => {
    if (!responses || !selectedRunTest) {
      return;
    }

    const data = {
      datasets: [
        {
          data: chartData().datum,
          label: "Latency", // for legend
        },
      ],
      labels: chartData().labels,
    };
    return <Line data={data} />;
  };

  const onSelectRunTest = (runTest: RunTest) => (e: React.FormEvent) => {
    e.preventDefault();
    setSelectedRunTest(runTest);
    listResponses();
  };

  const testContent = (test: Test) => {
    return (
      <React.Fragment>
        <div css={testDiv}>
          <div css={title}>{test.test_group?.name}</div>
          Test URL: {test.url} <br />
          Method: {test.method} <br />
        </div>
        <div css={testDiv}>
          <div css={title}>Run Tests</div>
          {test.run_tests &&
            test.run_tests.map((runTest: RunTest) => {
              return (
                <div
                  css={runTestDiv}
                  onClick={onSelectRunTest(runTest)}
                  key={runTest.id}
                >
                  <div css={runTestRow}>
                    <div css={runTestRowLeft}>
                      <FiArrowRightCircle
                        size="2.1rem"
                        color={runTest.passed ? "green" : "red"}
                      />
                    </div>
                    <div>
                      {runTest.id}-Start Time:{" "}
                      {moment(runTest.start_time).format(defaultFormat())} - End
                      Time:
                      {moment(runTest.end_time).format(defaultFormat())} -
                      Passed: {runTest.passed}
                    </div>
                  </div>
                </div>
              );
            })}
        </div>
      </React.Fragment>
    );
  };

  const responseTable = () => {
    if (!responses || !selectedRunTest) {
      return;
    }
    const content: any[][] = [];
    responses.map((response) => {
      content.push([
        moment(response.first_byte).format(preciseFormat()),
        moment(response.connect_start).format(preciseFormat()),
        moment(response.connect_done).format(preciseFormat()),
        moment(response.dns_start).format(preciseFormat()),
        moment(response.dns_done).format(preciseFormat()),
        moment(response.tls_start).format(preciseFormat()),
        moment(response.tls_done).format(preciseFormat()),
        response.status_code,
        response.total_time / 1000000,
        byteSize(response.body),
      ]);
    });

    return (
      <Fragment>
        <RTable
          content={content}
          title={[
            "FirstByte",
            "ConnectStart",
            "ConnectDone",
            "DNSStart",
            "DNSDone",
            "TLSStart",
            "TLSDone",
            "StatusCode",
            "TotalTime(ms)",
            "Body",
          ]}
        />
      </Fragment>
    );
  };

  return (
    <div>
      <div css={testContainer}>{test && testContent(test)}</div>
      {graph()}
      {responseTable()}
    </div>
  );
};

const testContainer = css`
  margin: 1rem 0 1rem 0;
  width: 100%;
  min-height: 20rem;
`;

const testDiv = css`
  width: 100%;
  ${MediaQuery[1]} {
    width: 100%;
  }
  margin: 1rem auto;
  padding: 3rem 2rem 3rem 2rem;
  background-color: #efefef;
  border-bottom: ${Borders.border1};
`;

const table = css`
  width: 100%;
`;

const runTestDiv = css`
  display: block;
  cursor: pointer;
  & :hover {
    background-color: lightgray;
  }
`;

const title = css`
  border-bottom: ${Borders.border1};
  font-weight: bold;
  font-size: 2.2rem;
  padding: 0 0 0.5rem 1rem;
  margin-bottom: 2rem;
`;

const runTestRow = css`
  display: flex;
`;

const runTestRowLeft = css`
  margin-right: 1rem;
`;

export default StatsContent;
