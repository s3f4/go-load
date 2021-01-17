/** @jsx jsx */
import { jsx, css } from "@emotion/core";
import React, { useState } from "react";
import { showSwarmNodes } from "../../api/entity/instance";
import { Node } from "../../api/entity/nodes";
import {
  card,
  cardContainer,
  cardContent,
  cardItem,
  cardTitle,
  MediaQuery,
} from "../style";
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
      <div>
        <div css={title}>Swarm Node List</div>
        <div css={cardContainer}>
          {nodes &&
            nodes.map((node: Node) => {
              return (
                <div css={card} key={node.ID}>
                  <h1 css={cardTitle}>{node.Description.Hostname}</h1>
                  <div css={cardContent}>
                    <div css={cardItem}>
                      <b>Role:</b>
                      <span>{node.Spec.Role}</span>
                    </div>
                    <div css={cardItem}>
                      <b>Addr:</b>
                      <span>{node.Status.Addr}</span>
                    </div>
                    <div css={cardItem}>
                      <b>State:</b>
                      <span>{node.Status.State}</span>
                    </div>
                    <div css={cardItem}>
                      <b>Memory:</b>
                      <span>
                        {((node.Description.Resources.MemoryBytes /
                          (1024 * 1024)) as number).toFixed(2)}
                        MB
                      </span>
                    </div>
                    <div css={cardItem}>
                      <b>CPUs:</b>
                      <span>
                        {node.Description.Resources.NanoCPUs / Math.pow(10, 9)}
                      </span>
                    </div>
                    <div css={cardItem}>
                      <b>Availability:</b>
                      <span>{node.Spec.Availability}</span>
                    </div>
                    <div css={cardItem}>
                      <b>Created:</b>
                      <span>
                        {moment(node.CreatedAt).format(defaultFormat())}
                      </span>
                    </div>
                  </div>
                </div>
              );
            })}
        </div>
      </div>
    );
  };

  return loading ? (
    <Loader message="swarm config is loading..." />
  ) : (
    swarmContent()
  );
};

const title = css`
  width: 100%;
  text-align: center;
  margin: 1rem auto;
  padding: 1rem;
  background-color: #efefef;

  ${MediaQuery[1]} {
    height: 4rem;
  }
`;

export default SwarmContent;
