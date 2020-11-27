/** @jsx jsx */
import React, { Fragment, useState } from "react";
import { jsx, css } from "@emotion/core";
import { listResponses, Response } from "../../../api/entity/stats";
import moment from "moment";
import { Line } from "react-chartjs-2";
import { defaultFormat, preciseFormat } from "../../basic/helper";
import { getTest, Test } from "../../../api/entity/test";
import { Borders, MediaQuery } from "../../style";
import { RunTest } from "../../../api/entity/runtest";
import { FiArrowRightCircle } from "react-icons/fi";
import RTable from "../../basic/RTable";
import { useParams } from "react-router-dom";

interface Props {}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [graphResponses, setGraphResponses] = useState<Response[]>([]);
  const [test, setTest] = useState<Test>();
  const [selectedRunTest, setSelectedRunTest] = useState<RunTest>();
  const { id }: any = useParams();

  React.useEffect(() => {
    getTest(id)
      .then((response) => {
        setTest(response.data);
      })
      .catch((error) => console.log(error));
  }, [id]);

  const byteSize = (str: string) => new Blob([str]).size;

  const chartData = React.useCallback((r) => {
    const datum: any[] = [];
    const labels: any[] = [];
    r.map((response: Response, index: number) => {
      datum.push(response.dns_time / 100000);
      labels.push("request_" + index);
      return null;
    });
    return {
      datum,
      labels,
    };
  }, []);

  const graph = React.useCallback(() => {
    if (!graphResponses || !selectedRunTest) {
      return;
    }

    const data = {
      datasets: [
        {
          data: chartData(graphResponses).datum,
          label: "Latency", // for legend
        },
      ],
      labels: chartData(graphResponses).labels,
    };
    return <Line data={data} />;
  }, [selectedRunTest, graphResponses]);

  const onSelectRunTest = (runTest: RunTest) => (e: React.FormEvent) => {
    e.preventDefault();
    setSelectedRunTest(runTest);
    listResponses(runTest.id!)().then((response) => {
      setGraphResponses(response.data.data);
    });
  };

  const testContent = (test: Test) => {
    return (
      <div css={container}>
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
      </div>
    );
  };

  const buildTable = (r: any) => {
    const content: any[][] = [];
    r.map((response: any) => {
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
      return null;
    });
    return content;
  };

  const responseTable = () => {
    if (!selectedRunTest) {
      return;
    }

    return (
      <Fragment>
        <RTable
          fetcher={listResponses(selectedRunTest.id!)}
          builder={buildTable}
          title={[
            {
              header: "FirstByte",
              accessor: "first_byte",
              sortable: true,
            },
            {
              header: "ConnectStart",
              accessor: "connect_start",
              sortable: true,
            },
            {
              header: "ConnectDone",
              accessor: "connect_done",
              sortable: true,
            },
            {
              header: "DNSStart",
              accessor: "dns_start",
              sortable: true,
            },
            {
              header: "DNSDone",
              accessor: "dns_done",
              sortable: true,
            },
            {
              header: "TLSStart",
              accessor: "tls_start",
              sortable: true,
            },
            {
              header: "TLSDone",
              accessor: "tls_done",
              sortable: true,
            },
            {
              header: "StatusCode",
              accessor: "status_code",
              sortable: true,
            },
            {
              header: "TotalTime(ms)",
              accessor: "TotalTime",
              sortable: true,
            },
            {
              header: "Body",
              sortable: false,
            },
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

const container = css`
  display: flex;
  width: 100%;
  flex-direction: column;
  ${MediaQuery[1]} {
    flex-direction: row;
  }
`;

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
