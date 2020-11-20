/** @jsx jsx */
import { jsx, css } from "@emotion/core";
import React from "react";
import Loader from "../basic/Loader";
import { card, cardContainer, cardTitle } from "../style";
import { Service } from "../../api/entity/service";

interface Props {
  services?: Service[];
  loader: boolean;
}

const ServicesContent: React.FC<Props> = (props: Props) => {
  const servicesDiv = () =>
    props.services?.map((service: Service) => {
      return (
        <div css={card} key={service.Id}>
          <h1 css={cardTitle}>{service.Names[0].substr(1)}</h1>
          <br />
          {service.Id.substr(0, 7)} <br />
          {service.Status} <br />
          {service.State}
          <br />
        </div>
      );
    });

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