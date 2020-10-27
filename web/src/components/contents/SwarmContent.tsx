import React from "react";
import { css } from "@emotion/core";
import { showSwarmNodes } from "../../api/entity/instance";
import { Node } from "../../api/entity/nodes";
import { card, cardContainer, cardTitle } from "../style";
import moment from "moment";
import { defaultFormat } from "../basic/helper";

const SwarmContent: React.FC = () => {
  const [nodes, setNodes] = React.useState<Node[]>([]);
  React.useEffect(() => {
    showSwarmNodes()
      .then((response) => setNodes(response.data))
      .catch((error) => console.log(error));
  }, []);

  return (
    <div css={cardContainer}>
      {nodes &&
        nodes.map((node: Node) => {
          return (
            <div css={card} key={node.ID}>
              <h1 css={cardTitle}>{node.Description.Hostname}</h1>
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

export default SwarmContent;
