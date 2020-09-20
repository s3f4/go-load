/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { stats, Response } from "../../api/entity/stats";

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

  return (
    <div css={statsContainer}>
      <table>
        <thead>
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
        </thead>
        <tbody>
          {responses &&
            responses.map((response: Response, key: number) => {
              return (
                <tr>
                  <td>{response.FirstByte}</td>
                  <td>{response.ConnectStart}</td>
                  <td>{response.ConnectDone}</td>
                  <td>{response.DNSStart}</td>
                  <td>{response.DNSDone}</td>
                  <td>{response.TLSStart}</td>
                  <td>{response.TLSDone}</td>
                  <td>{response.StatusCode}</td>
                  <td>{response.TotalTime / 1000000}</td>
                  <td>{response.Body}</td>
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

export default StatsContent;
