/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { showSwarmNodes } from "../../api/entity/instance";
import { Node } from "../../api/entity/nodes";
import { Box, Sizes } from "../style";
import moment from "moment";
import { defaultFormat } from "../basic/helper";

interface Props {}

const SwarmContent: React.FC<Props> = (props: Props) => {
  const [nodes, setNodes] = React.useState<Node[]>([]);
  React.useEffect(() => {
    showSwarmNodes()
      .then((response) => setNodes(response.data))
      .catch((error) => console.log(error));
  }, []);

  return (
    <div css={swarmContainer}>
      {nodes &&
        nodes.map((node: Node) => {
          return (
            <div css={nodeCard} key={node.ID}>
              <h1 css={nodeTitle}>{node.Description.Hostname}</h1>
              Role: {node.Spec.Role} <br />
              Addr: {node.Status.Addr} <br />
              State: {node.Status.State}
              <br />
              Memory :{" "}
              {((node.Description.Resources.MemoryBytes /
                (1024 * 1024)) as number).toFixed(2)}{" "}
              MB
              <br />
              CPUs : {node.Description.Resources.NanoCPUs / Math.pow(10, 9)}
              <br />
              Availability: {node.Spec.Availability}
              <br />
              Created: {moment(node.CreatedAt).format(defaultFormat())}
              <br />
            </div>
          );
        })}
    </div>
  );
};

const swarmContainer = css`
  display: flex;
  flex-wrap: wrap;
  height: 100%;
`;

const nodeCard = css`
  width: 28rem;
  height: 25rem;
  margin: 1rem 1rem;
  border: 1px solid black;
  text-align: center;
  ${Box.boxShadow1}
  border-radius: ${Sizes.borderRadius1}
`;

const nodeTitle = css`
  background-color: #007d9c;
  color: white;
  width: 100%;
  height: 100;
  padding: 0.5rem;
`;

export default SwarmContent;
