/** @jsx jsx */
import React, { Fragment, useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import { MediaQuery } from "../style";
import { FiArrowDown } from "react-icons/fi";
import { ServerResponse } from "../../api/api";
import { Query } from "./query";
import Button, { ButtonType } from "./Button";

export interface TableTitle {
  header: string;
  accessor?: string;
  sortable: boolean;
}
interface Props {
  title: TableTitle[];
  builder: (data: any) => any[][];
  fetcher: (query?: Query) => Promise<any>;
  setter: (val: any) => any;
  limit?: number;
}

const RTable: React.FC<Props> = (props: Props) => {
  const [content, setContent] = useState<any[][]>([]);
  const [total, setTotal] = useState<number>(0);
  const [query, setQuery] = useState<Query>({
    limit: props.limit ?? 10,
    offset: 0,
  });

  useEffect(() => {
    props.fetcher(query).then((response: ServerResponse) => {
      setTotal(response.data.total);
      props.setter(response.data.data);
      setContent(props.builder(response.data.data));
    });
    return () => {};
  }, [query]);

  const onChangePage = (page: number) => (e: React.FormEvent) => {
    e.preventDefault();
    setQuery({
      ...query,
      ["offset"]: (page - 1) * query.limit,
    });
  };

  const pages = () => {
    const buttons = [];
    const p = total / query.limit;
    const page = p > 1 ? Math.ceil(p) : p;
    for (let i = 1; i <= page; i++) {
      buttons.push(
        <Button
          type={ButtonType.small}
          text={i + ""}
          onClick={onChangePage(i)}
          key={i}
        />,
      );
    }
    return buttons;
  };

  return (
    <Fragment>
      <div css={container}>
        <div css={row(true)}>
          {props.title.map((title: TableTitle, index) => (
            <div css={columnStyle(true)} key={index}>
              <b>{title.header}</b> <FiArrowDown />
            </div>
          ))}
        </div>

        {content.map((rows, index) => {
          return (
            <div key={index} css={row(false)}>
              {rows.map((column, colIndex) => (
                <div css={columnStyle()} key={colIndex}>
                  {column}
                </div>
              ))}
            </div>
          );
        })}
      </div>

      <div css={mobileContainer}>
        {content.map((rows, index) => {
          return (
            <div key={index} css={mobileRow}>
              {rows.map((column, colIndex) => (
                <div css={mobileFlex} key={colIndex}>
                  <b>{props.title[colIndex].header}</b>
                  {column}
                </div>
              ))}
            </div>
          );
        })}
      </div>
      {pages()}
    </Fragment>
  );
};

const mobileContainer = css`
  display: block;
  ${MediaQuery[2]} {
    display: none;
  }
  width: 100%;
  border: 1px solid #e1e1e1;
  border-radius: 0.5rem;
  text-align: left;
  padding: 1rem 1rem 1rem 1rem;
`;

const mobileFlex = css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 5rem;
  min-height: 4rem;
`;

const mobileRow = css`
  margin-top: 1rem;
  padding: 2rem;
  background-color: #e1e1e1;
`;

const row = (title?: boolean) => css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 4.5rem;
  border-bottom: 1px solid black;
  background-color: ${title ? "#007d9c" : "none"};
  color: ${title ? "white" : "none"};
`;

const columnStyle = (sortable?: boolean) => css`
  flex: 0 1 20rem;
  padding-left: 1rem;
  padding-top: 1rem;
  text-align: center;
  width: 7rem;
  ${sortable ? "cursor:pointer;" : ""}
`;

const container = css`
  display: none;
  ${MediaQuery[2]} {
    display: flex;
    flex-direction: column;
    width: 100%;
    border: 1px solid #e1e1e1;
    border-radius: 0.5rem;
    background-color: #e1e1e1;
    text-align: left;
    padding: 1rem 1rem 1rem 1rem;
  }
`;

export default RTable;
