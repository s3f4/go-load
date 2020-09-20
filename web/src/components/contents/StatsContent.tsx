/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { stats, Response } from "../../api/entity/stats";
import moment from "moment";

interface Props {}

const StatsContent: React.FC<Props> = (props: Props) => {
  const [responses, setResponses] = React.useState<Response[]>([]);

  React.useEffect(() => {
    listResponses();
  }, []);

  const listResponses = () => {
    stats()
      .then((response) => {
        setResponses(response.data.data);
        console.log(response.data);
      })
      .catch((error) => console.log(error));
  };

  const format = (): string => {
    return "hh:mm:ss SSSS";
  };

  const byteSize = (str: string) => new Blob([str]).size;

  return (
    <div css={statsContainer}>
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
                <tr>
                  <td>{moment(response.FirstByte).format(format())}</td>
                  <td>{moment(response.ConnectStart).format(format())}</td>
                  <td>{moment(response.ConnectDone).format(format())}</td>
                  <td>{moment(response.DNSStart).format(format())}</td>
                  <td>{moment(response.DNSDone).format(format())}</td>
                  <td>{moment(response.TLSStart).format(format())}</td>
                  <td>{moment(response.TLSDone).format(format())}</td>
                  <td>{moment(response.StatusCode).format(format())}</td>
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

const table = css`
  width: 100%;
`;

export default StatsContent;
