/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { stats, Response } from "../../api/entity/stats";
import moment from "moment";
import { Line } from "react-chartjs-2";
import { preciseFormat } from "../basic/helper";
import { getTest, Test } from "../../api/entity/test";
import { Borders, MediaQuery } from "../style";
import { RunTest } from "../../api/entity/runtest";

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
    if (!responses) {
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
      <div css={testDiv}>
        Test URL: {test.url} <br />
        Method: {test.method} <br />
        <hr />
        Run Tests
        {test.run_tests &&
          test.run_tests.map((runTest: RunTest) => {
            return (
              <div
                css={runTestDiv}
                onClick={onSelectRunTest(runTest)}
                key={runTest.id}
              >
                Start Time: {runTest.start_time} - End Time: {runTest.end_time}{" "}
                - Passed: {runTest.passed}
                <br />
              </div>
            );
          })}
      </div>
    );
  };

  const responseTable = () => (
    <table css={table}>
      <thead>
        <tr>
          <th>FirstByte</th>
          <th>ConnectStart</th>
          <th>ConnectDone</th>
          <th>DNSStart</th>
          <th>DNSDone</th>
          <th>TLSStart</th>
          <th>TLSDone</th>
          <th>StatusCode</th>
          <th>TotalTime</th>
          <th>Body</th>
        </tr>
      </thead>
      <tbody>
        {responses &&
          responses.map((response: Response, key: number) => {
            console.log(JSON.stringify(response));
            return (
              <tr key={key}>
                <td>{moment(response.first_byte).format(preciseFormat())}</td>
                <td>
                  {moment(response.connect_start).format(preciseFormat())}
                </td>
                <td>{moment(response.connect_done).format(preciseFormat())}</td>
                <td>{moment(response.dns_start).format(preciseFormat())}</td>
                <td>{moment(response.dns_done).format(preciseFormat())}</td>
                <td>{moment(response.tls_start).format(preciseFormat())}</td>
                <td>{moment(response.tls_done).format(preciseFormat())}</td>
                <td>{response.status_code}</td>
                <td>{response.total_time / 1000000}</td>
                <td>{byteSize(response.body)}</td>
              </tr>
            );
          })}
      </tbody>
    </table>
  );

  return (
    <div css={statsContainer}>
      <div css={testContainer}>{test && testContent(test)}</div>
      {graph()}
      {responseTable()}
    </div>
  );
};

const statsContainer = css``;

const testContainer = css`
  margin: 1rem 0 1rem 0;
  width: 100%;
  min-height: 20rem;
`;

const testDiv = css`
  width: 100%;
  ${MediaQuery[1]} {
    width: 90%;
  }
  margin: 0 auto;
  padding: 3rem 2rem 3rem 2rem;
  background-color: #efefef;
  border: ${Borders.border1};
`;

const table = css`
  width: 100%;
`;

const runTestDiv = css`
  display: block;
  cursor: pointer;
`;

export default StatsContent;
