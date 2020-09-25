/** @jsx jsx */
import React from "react";
import { jsx, css } from "@emotion/core";
import { showSwarmNodes } from "../../api/entity/instance";
import { Node } from "../../api/entity/nodes";

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
          return <div key={node.ID}>{node.Description.Hostname}</div>;
        })}
    </div>
  );
};

const swarmContainer = css``;

export default SwarmContent;
