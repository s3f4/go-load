/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { stats, Response } from "../../api/entity/stats";
import moment from "moment";
import { Line } from "react-chartjs-2";
import { preciseFormat } from "../basic/helper";
import { getTest, runTest, Test } from "../../api/entity/test";
import { Borders } from "../style";
import { RunTest } from "../../api/entity/runtest";

interface Props {
  testID: number;
}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [responses, setResponses] = useState<Response[]>([]);
  const [test, setTest] = useState<Test>();

  React.useEffect(() => {
    getTest(props.testID)
      .then((response) => {
        setTest(response.data);
      })
      .catch((error) => console.log(error));
  }, [props.testID]);

  const listResponses = () => {
    if (test) {
      stats(test.id!)
        .then((response) => {
          setResponses(response.data.data);
          console.log(response.data);
        })
        .catch((error) => console.log(error));
    }
  };

  const byteSize = (str: string) => new Blob([str]).size;

  const chartData = () => {
    const datum: any[] = [];
    const labels: any[] = [];
    responses.map((response: Response, index: number) => {
      datum.push(response.DNSTime);
      labels.push("request_" + index);
      return null;
    });
    return {
      datum,
      labels,
    };
  };

  const data = {
    datasets: [
      {
        data: chartData().datum,
        label: "Latency", // for legend
      },
    ],
    labels: chartData().labels,
  };

  const graph = () => {
    return <Line data={data} />;
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
              <div>
                Start Time: {runTest.start_time}
                <br />
                End Time: {runTest.end_time}
                <br />
                Passed: {runTest.passed}
                <br />
              </div>
            );
          })}
      </div>
    );
  };

  return (
    <div css={statsContainer}>
      <div css={testContainer}>{test && testContent(test)}</div>
      {graph()}
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
              return (
                <tr key={key}>
                  <td>{moment(response.FirstByte).format(preciseFormat())}</td>
                  <td>
                    {moment(response.ConnectStart).format(preciseFormat())}
                  </td>
                  <td>
                    {moment(response.ConnectDone).format(preciseFormat())}
                  </td>
                  <td>{moment(response.DNSStart).format(preciseFormat())}</td>
                  <td>{moment(response.DNSDone).format(preciseFormat())}</td>
                  <td>{moment(response.TLSStart).format(preciseFormat())}</td>
                  <td>{moment(response.TLSDone).format(preciseFormat())}</td>
                  <td>{response.StatusCode}</td>
                  <td>{response.TotalTime / 1000000}</td>
                  <td>{byteSize(response.Body)}</td>
                </tr>
              );
            })}
        </tbody>
      </table>
      {responses.length}
    </div>
  );
};

const statsContainer = css``;

const testContainer = css`
  margin: 1rem 0 1rem 0;
  width: 100%;
  height: 12rem;
`;

const testDiv = css`
  width: 80%;
  margin: 0 auto;
  padding: 3rem 2rem 3rem 2rem;
  background-color: #efefef;
  border: ${Borders.border1};
`;

const table = css`
  width: 100%;
`;

export default StatsContent;
