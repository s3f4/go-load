/** @jsx jsx */
import { jsx, css } from "@emotion/core";
import React from "react";
import Loader from "../basic/Loader";
import {
  card,
  cardContainer,
  cardContent,
  cardItem,
  cardTitle,
} from "../style";
import { Service } from "../../api/entity/service";
import Message, { IMessage } from "../basic/Message";

interface Props {
  services?: Service[];
  loader: boolean;
  message?: IMessage;
}

const ServicesContent: React.FC<Props> = (props: Props) => {
  const servicesDiv = () => {
    if (props.message) {
      return (
        <Message message={props.message.message} type={props.message.type} />
      );
    }

    return props.services?.map((service: Service) => {
      return (
        <div css={card} key={service.Id}>
          <h1 css={cardTitle}>{service.Names[0].substr(1)}</h1>
          <div css={cardContent}>
            <div css={cardItem}>
              <b>ID:</b>
              <span>{service.Id.substr(0, 7)}</span>
            </div>
            <div css={cardItem}>
              <b>Status:</b>
              <span>{service.Status}</span>
            </div>
            <div css={cardItem}>
              <b>State:</b>
              <span>{service.State}</span>
            </div>
          </div>
        </div>
      );
    });
  };

  return (
    <div css={cardContainer}>
      {!props.loader ? (
        servicesDiv()
      ) : (
        <Loader message={"services list is loading..."} />
      )}
    </div>
  );
};

export default ServicesContent;
