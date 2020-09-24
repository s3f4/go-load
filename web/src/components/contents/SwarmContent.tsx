/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { showSwarmNodes } from "../../api/entity/instance";

interface Props {}

const SwarmContent: React.FC<Props> = (props: Props) => {
  React.useEffect(() => {
    showSwarmNodes()
      .then((response) => console.log(response))
      .catch((error) => console.log(error));
  }, []);

  return <div css={swarmContainer}></div>;
};

const swarmContainer = css``;

export default SwarmContent;
