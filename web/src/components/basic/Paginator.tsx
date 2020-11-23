/** @jsx jsx */
import React, { Fragment, ReactNode, useEffect, useState } from "react";
import { jsx, css } from "@emotion/core";
import { Query } from "./query";
import { ServerResponse } from "../../api/api";

interface Props {
  // children: ReactNode;
  fetcher: (query: Query) => Promise<any>;
  setter: (val: any) => any;
  limit?: number;
}

const Paginator: React.FC<Props> = (props: Props) => {
  const [total, setTotal] = useState<number>(0);
  const [query, setQuery] = useState<Query>({
    limit: props.limit ?? 10,
    offset: 0,
  });

  useEffect(() => {
    props.fetcher(query).then((response: ServerResponse) => {
      setTotal(response.data.total);
      props.setter(response.data.data);
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
    for (let i = 1; i <= Math.ceil(total / query.limit); i++) {
      buttons.push(
        <button onClick={onChangePage(i)} key={i}>
          {i}
        </button>,
      );
    }
    return buttons;
  };

  return <Fragment>{pages()}</Fragment>;
};

export default Paginator;
