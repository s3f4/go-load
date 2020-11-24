/** @jsx jsx */
import { jsx, css } from "@emotion/core";
import React, { useState } from "react";
import { showSwarmNodes } from "../../api/entity/instance";
import { Node } from "../../api/entity/nodes";
import { card, cardContainer, cardTitle } from "../style";
import moment from "moment";
import { defaultFormat } from "../basic/helper";
import Message, { IMessage } from "../basic/Message";
import Loader from "../basic/Loader";

const SwarmContent: React.FC = () => {
  const [nodes, setNodes] = useState<Node[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [message, setMessage] = useState<IMessage>();

  React.useEffect(() => {
    setLoading(true);
    showSwarmNodes()
      .then((response) => {
        setNodes(response.data);
        setLoading(false);
      })
      .catch((error) => {
        setLoading(false);
        setMessage({ message: error.message, type: "error" });
      });
  }, []);

  const swarmContent = () => {
    if (message) {
      return <Message message={message.message} type={message.type} />;
    }

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
                  (1024 * 1024)) as number).toFixed(2)}
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

  console.log(message);

  return loading ? (
    <Loader message="swarm config is loading..." />
  ) : (
    swarmContent()
  );
};

export default SwarmContent;
