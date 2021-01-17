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
  MediaQuery,
} from "../style";
import { Service } from "../../api/entity/service";
import Message, { IMessage } from "../basic/Message";
import moment from "moment";
import { defaultFormat, preciseFormat } from "../basic/helper";
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

    return (
      <React.Fragment>
        <div css={title}>Swarm Service List</div>
        <div css={cardContainer}>
          {props.services?.map((service: Service) => {
            return (
              <div css={card} key={service.ID}>
                <h1 css={cardTitle}>{service.Spec.Name}</h1>
                <div css={cardContent}>
                  <div css={cardItem}>
                    <b>Name:</b>
                    <span>{service.Spec.Name}</span>
                  </div>
                  <div css={cardItem}>
                    <b>Replicas:</b>
                    <span>{service.Spec.Mode.Replicated.Replicas}</span>
                  </div>
                  <div css={cardItem}>
                    <b>CreatedAt:</b>
                    <span>
                      {moment(service.CreatedAt).format(defaultFormat())}
                    </span>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </React.Fragment>
    );
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
export default ServicesContent;
