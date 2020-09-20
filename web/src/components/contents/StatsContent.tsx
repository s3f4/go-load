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
        console.log(responses);
      })
      .catch((error) => console.log(error));
  };

  return (
    <div css={statsContainer}>
      {responses &&
        responses.map((response: Response, key: number) => {
          return (
            <div key={key}>
              ConnectStart: {response.ConnectStart} <br />
              ConnectDone:{response.ConnectDone}
              <br />
              DNSStart: {response.DNSStart}
              <br />
              DNSDone: {response.DNSDone}
              <br />
              TLSStart: {response.TLSStart}
              <br />
              TLSDone:{response.TLSDone}
              <br />
              StatusCode:{response.StatusCode}
              <br />
              Total Time:{response.TotalTime / 1000000}
              <br />
            </div>
          );
        })}
      {responses.length}
    </div>
  );
};

const statsContainer = css``;

export default StatsContent;
