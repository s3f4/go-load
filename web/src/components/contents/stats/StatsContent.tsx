/** @jsx jsx */
import React, { useState } from "react";
import { jsx, css } from "@emotion/core";
import { listResponses, Response } from "../../../api/entity/response";
import moment from "moment";
import { Line } from "react-chartjs-2";
import { defaultFormat, preciseFormat } from "../../basic/helper";
import { getTest, Test } from "../../../api/entity/test";
import { listRunTestsOfTest } from "../../../api/entity/runtest";
import { Borders, MediaQuery, Box, Sizes } from "../../style";
import { RunTest } from "../../../api/entity/runtest";
import { FiArrowRightCircle } from "react-icons/fi";
import RTable, { RTableRow } from "../../basic/RTable";
import { useParams } from "react-router-dom";
import Paginator from "../../basic/Paginator";

interface Props {}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [responses, setResponses] = useState<Response[]>([]);
  const [test, setTest] = useState<Test>();
  const [selectedRunTest, setSelectedRunTest] = useState<RunTest>();
  const [runTests, setRunTests] = useState<RunTest[]>([]);

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
    if (!responses || !selectedRunTest) {
      return;
    }

    const data = {
      datasets: [
        {
          data: chartData(responses).datum,
          label: "Latency", // for legend
        },
      ],
      labels: chartData(responses).labels,
    };
    return <Line data={data} />;
  }, [selectedRunTest, responses]);

  const onSelectRunTest = (runTest: RunTest) => (e: React.FormEvent) => {
    e.preventDefault();
    setSelectedRunTest(runTest);
    console.log(runTest);
  };

  const testContent = React.useCallback(
    (test: Test) => {
      return (
        <div css={container}>
          <div css={testDiv}>
            <div css={title}>{test.name}</div>
            {Object.keys(test).map((key) => {
              if (
                [
                  "id",
                  "transport_config",
                  "test_group",
                  "test_group_id",
                  "run_tests",
                ].includes(key)
              ) {
                return;
              }
              return (
                <div key={key} css={testPropRow}>
                  <div css={testProp}>{key}</div>
                  <div>{JSON.stringify(test[key])}</div>
                </div>
              );
            })}
          </div>
          <div css={testDiv}>
            <div css={title}>Finished Tests</div>
            <div css={runTestDiv}>
              {runTests &&
                runTests.map((runTest: RunTest) => {
                  console.log(runTest);
                  return (
                    <div
                      css={runTestRow}
                      onClick={onSelectRunTest(runTest)}
                      key={runTest.id}
                    >
                      <div css={runTestRowLeft}>
                        <FiArrowRightCircle
                          size="2.1rem"
                          color={runTest.passed ? "green" : "red"}
                        />
                      </div>
                      <div>
                        Start:{" "}
                        {moment(runTest.start_time).format(defaultFormat())}
                      </div>
                      <div>
                        End: {moment(runTest.end_time).format(defaultFormat())}{" "}
                      </div>
                      <div>Passed: {runTest.passed}</div>
                    </div>
                  );
                })}
            </div>
            <Paginator
              limit={5}
              fetcher={listRunTestsOfTest(id)}
              setter={setRunTests}
            />
          </div>
        </div>
      );
    },
    [runTests],
  );

  const buildTable = React.useCallback((r: Response[]) => {
    const rows: RTableRow[] = [];
    r.forEach((response: Response) => {
      const row: RTableRow = {
        columns: [
          { content: moment(response.first_byte).format(preciseFormat()) },
          { content: moment(response.connect_start).format(preciseFormat()) },
          { content: moment(response.connect_done).format(preciseFormat()) },
          { content: moment(response.dns_start).format(preciseFormat()) },
          { content: moment(response.dns_done).format(preciseFormat()) },
          { content: moment(response.tls_start).format(preciseFormat()) },
          { content: moment(response.tls_done).format(preciseFormat()) },
          { content: response.status_code },
          { content: response.total_time / 1000000 },
          { content: byteSize(response.body) },
        ],
      };
      rows.push(row);
    });
    return rows;
  }, []);

  const responseTable = () => {
    if (!selectedRunTest) {
      return;
    }

    return (
      <RTable
        limit={50}
        setter={setResponses}
        fetcher={listResponses(selectedRunTest?.id!)}
        builder={buildTable}
        trigger={selectedRunTest}
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
  display: flex;
  flex-direction: column;
  justify-content: space-around;
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
  border: ${Borders.border1};
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1};
  justify-content: space-between;
  margin-bottom: 0.5rem;
  padding: 0.7rem;
`;

const runTestRowLeft = css`
  margin-right: 1rem;
`;

const testPropRow = css`
  display: flex;
  justify-content: space-between;
`;

const testProp = css`
  font-weight: bold;
`;

export default StatsContent;
