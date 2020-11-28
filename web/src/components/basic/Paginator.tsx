/** @jsx jsx */
import React, { Fragment, useEffect, useState } from "react";
import { jsx } from "@emotion/core";
import { Query } from "./query";
import { ServerResponse } from "../../api/api";
import Button from "./Button";

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

  const { fetcher, setter } = props;

  useEffect(() => {
    let mount = true;
    fetcher(query).then((response: ServerResponse) => {
      if (mount) {
        setTotal(response.data.total);
        setter(response.data.data);
      }
    });
    return () => {
      mount = false;
    };
  }, [query]);

  const onChangePage = (page: number) => (e: React.FormEvent) => {
    e.preventDefault();
    setQuery({
      ...query,
      offset: (page - 1) * query.limit,
    });
  };

  const pages = () => {
    const buttons = [];
    const p = total / query.limit;
    const page = p > 1 ? Math.ceil(p) : p;
    for (let i = 1; i <= page; i++) {
      buttons.push(<Button key={i} text={i + ""} onClick={onChangePage(i)} />);
    }
    return buttons;
  };

  return <Fragment>{pages()}</Fragment>;
};

export default Paginator;
