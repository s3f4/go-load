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
import RTable, { IRTableRow } from "../../basic/RTable";
import { useParams } from "react-router-dom";
import Paginator from "../../basic/Paginator";
import Loader from "../../basic/Loader";
import { nanoToMilli } from "../../basic/helper";
interface Props {}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [responses, setResponses] = useState<Response[]>([]);
  const [test, setTest] = useState<Test>();
  const [selectedRunTest, setSelectedRunTest] = useState<RunTest>();
  const [runTests, setRunTests] = useState<RunTest[]>([]);
  const [loader, setLoader] = useState<boolean>(false);

  const { id }: any = useParams();

  React.useEffect(() => {
    getTest(id)
      .then((response) => {
        setTest(response.data);
      })
      .catch((error) => console.log(error));
  }, [id]);

  const chartData = React.useCallback((r) => {
    const labels: any[] = [];

    const dns_time: number[] = [];
    const connect_time: number[] = [];
    const tls_time: number[] = [];
    const first_byte_time: number[] = [];
    const total_time: number[] = [];

    r.map((response: Response, index: number) => {
      first_byte_time.push(nanoToMilli(response.first_byte_time));
      connect_time.push(nanoToMilli(response.connect_time));
      dns_time.push(nanoToMilli(response.dns_time));
      tls_time.push(nanoToMilli(response.tls_time));
      total_time.push(nanoToMilli(response.total_time));

      labels.push("request_" + index);
      return null;
    });
    const data: any = {
      first_byte_time,
      connect_time,
      dns_time,
      tls_time,
      total_time,
    };
    return {
      data,
      labels,
    };
  }, []);

  const graph = React.useCallback(() => {
    if (!responses || !selectedRunTest) {
      return;
    }

    if (loader) {
      return <Loader message={"Response Graph loading..."} />;
    }

    const graphData = chartData(responses);

    const data = {
      datasets: [
        {
          data: graphData.data.first_byte_time,
          label: "First Byte Time(ms)",
        },
        {
          data: graphData.data.connect_time,
          label: "Connect Time(ms)",
        },
        {
          data: graphData.data.dns_time,
          label: "DNS Time(ms)",
        },
        {
          data: graphData.data.tls_time,
          label: "TLS Time(ms)",
        },
        {
          data: graphData.data.total_time,
          label: "Total Time(ms)",
        },
      ],
      labels: graphData.labels,
    };
    return <Line data={data} />;
  }, [selectedRunTest, responses]);

  const onSelectRunTest = (runTest: RunTest) => (e: React.FormEvent) => {
    e.preventDefault();
    setLoader(true);
    setSelectedRunTest(runTest);
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
                  return (
                    <div
                      css={runTestRow(runTest.id === selectedRunTest?.id)}
                      onClick={onSelectRunTest(runTest)}
                      key={runTest.id}
                    >
                      <div css={runTestRowLeft}>
                        <FiArrowRightCircle
                          size="2.1rem"
                          color={runTest.passed ? "#87b666" : "#ff6961"}
                        />
                      </div>
                      <div>
                        Start:{" "}
                        {moment(runTest.start_time).format(defaultFormat())}
                      </div>
                      <div>
                        End: {moment(runTest.end_time).format(defaultFormat())}{" "}
                      </div>
                      <div>Passed: {runTest.passed.toString()}</div>
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
    [runTests, selectedRunTest],
  );

  const buildTable = React.useCallback((r: Response[]) => {
    const rows: IRTableRow[] = [];
    r.forEach((response: Response) => {
      const reasons = response.reasons.split("\n");

      const row: IRTableRow = {
        rowStyle: response.passed
          ? "background-color:#87b666;"
          : "background-color:#ff6961;",
        allColumns: [
          { content: moment(response.start_time).format(preciseFormat()) },
          { content: moment(response.first_byte).format(preciseFormat()) },
          { content: moment(response.connect_start).format(preciseFormat()) },
          { content: moment(response.connect_done).format(preciseFormat()) },
          { content: moment(response.dns_start).format(preciseFormat()) },
          { content: moment(response.dns_done).format(preciseFormat()) },
          { content: moment(response.tls_start).format(preciseFormat()) },
          { content: moment(response.tls_done).format(preciseFormat()) },
          { content: response.status_code },
          { content: response.total_time },
          { content: response.body },
          { content: response.passed.toString() },
          {
            content: (
              <div>
                {reasons.map((reason, index) => {
                  return <div key={index}>{reason}</div>;
                })}
              </div>
            ),
          },
        ],
        columns: [
          { content: moment(response.start_time).format(preciseFormat()) },
          { content: moment(response.first_byte).format(preciseFormat()) },
          { content: nanoToMilli(response.first_byte_time) },
          { content: nanoToMilli(response.connect_time) },
          { content: nanoToMilli(response.dns_time) },
          { content: nanoToMilli(response.tls_time) },
          { content: response.status_code },
          { content: nanoToMilli(response.total_time) },
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
        setLoader={setLoader}
        loader={loader}
        limit={50}
        setter={setResponses}
        fetcher={listResponses(selectedRunTest?.id!)}
        builder={buildTable}
        trigger={selectedRunTest}
        allTitles={[
          { header: "StartTime" },
          { header: "FirstByte" },
          { header: "ConnectStart" },
          { header: "ConnectDone" },
          { header: "DNSStart" },
          { header: "DNSDone" },
          { header: "TLSStart" },
          { header: "TLSDone" },
          { header: "StatusCode" },
          { header: "TotalTime(ms)" },
          { header: "Body" },
          { header: "Passed" },
          { header: "Reasons" },
        ]}
        title={[
          {
            header: "Start",
            accessor: "start_time",
            sortable: true,
            width: "18%",
          },
          {
            header: "FirstByte",
            accessor: "first_byte",
            sortable: true,
            width: "18%",
          },
          {
            header: "FirstByteTime",
            accessor: "first_byte_time",
            sortable: true,
            width: "12%",
          },
          {
            header: "Connect",
            accessor: "connect_time",
            sortable: true,
            width: "10%",
          },
          {
            header: "DNS",
            accessor: "DNS_time",
            sortable: true,
            width: "10%",
          },
          {
            header: "TLS",
            accessor: "TLS_time",
            sortable: true,
            width: "10%",
          },
          {
            header: "Status",
            accessor: "status_code",
            sortable: true,
            width: "10%",
          },
          {
            header: "Total(ms)",
            accessor: "total_time",
            sortable: true,
            width: "12%",
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
  width: 50%;
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

const runTestRow = (selected: boolean) => {
  return css`
    display: flex;
    border: ${Borders.border1};
    ${Box.boxShadow1}
    border-radius: ${Sizes.borderRadius1};
    justify-content: space-between;
    margin-bottom: 0.5rem;
    padding: 0.7rem;
    ${selected ? "background-color:#FFA042;" : ""};
  `;
};

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
