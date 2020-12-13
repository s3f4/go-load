/** @jsx jsx */
import React, {
  FormEvent,
  Fragment,
  ReactNode,
  useEffect,
  useState,
} from "react";
import { jsx, css } from "@emotion/core";
import { DisableSelect, MediaQuery } from "../style";
import { FiArrowDown, FiArrowUp } from "react-icons/fi";
import { ServerResponse } from "../../api/api";
import { Query } from "./query";
import Button, { ButtonColorType, ButtonType } from "./Button";
import RTableRow from "./RTableRow";

export interface TableTitle {
  header: string;
  // accessor is model column if there is one
  accessor?: string;
  sortable?: boolean;
  width?: string;
}

export interface IRTableRow {
  modal?: (show: boolean) => ReactNode;
  rowStyle?: any;
  allColumns?: IRTableColumn[];
  columns: IRTableColumn[];
}
export interface IRTableColumn {
  columnStyle?: any;
  content: ReactNode;
}

interface Props {
  title: TableTitle[];
  allTitles?: TableTitle[];
  builder: (data: any) => IRTableRow[];
  fetcher: (query?: Query) => Promise<ServerResponse>;
  setter?: (data: any[]) => void;
  limit?: number;
  trigger?: any;
}

const RTable: React.FC<Props> = (props: Props) => {
  const [increment, setIncrement] = useState<boolean>(false);
  const [orderedCol, setOrderCol] = useState<string>();
  const [selectedPage, setSelectedPage] = useState<number>(1);
  const [content, setContent] = useState<IRTableRow[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [query, setQuery] = useState<Query>({
    limit: props.limit ?? 10,
    offset: 0,
  });

  const { fetcher, setter, builder, trigger } = props;

  useEffect(() => {
    fetcher(query).then((response: ServerResponse) => {
      setTotal(response.data.total);
      setContent(builder(response.data.data));
      setter?.(response.data.data);
    });
    return () => {};
  }, [query, trigger]);

  const onOrder = (sortable: boolean, col: string) => (e: FormEvent) => {
    e.preventDefault();
    if (sortable) {
      setIncrement(!increment);
      setOrderCol(col);
      const incrementStr = !increment ? "i__" : "d__";
      setQuery({
        limit: props.limit ?? 10,
        offset: 0,
        order: incrementStr + col,
      });
    }
  };

  const onChangePage = (page: number) => (e: FormEvent) => {
    e.preventDefault();
    setSelectedPage(page);
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
      buttons.push(
        <div key={i} css={buttonDiv}>
          <Button
            colorType={
              i === selectedPage
                ? ButtonColorType.primary
                : ButtonColorType.secondary
            }
            type={ButtonType.small}
            text={i + ""}
            onClick={onChangePage(i)}
          />
        </div>,
      );
    }
    return buttons;
  };

  return (
    <Fragment>
      <div css={container}>
        <div css={row(true)}>
          {props.title.map((title: TableTitle, index) => (
            <div
              onClick={onOrder(title.sortable!, title.accessor!)}
              css={columnStyle(title.width, title.sortable)}
              key={index}
            >
              <b>{title.header}</b>{" "}
              {title.sortable && orderedCol === title.accessor ? (
                increment ? (
                  <FiArrowUp />
                ) : (
                  <FiArrowDown />
                )
              ) : (
                ""
              )}
            </div>
          ))}
        </div>

        {content.map((r, index) => {
          return (
            <RTableRow
              key={index}
              mobile={false}
              row={r}
              title={props.title}
              allTitles={props.allTitles}
            />
          );
        })}
      </div>

      <div css={mobileContainer}>
        {content.map((r, index) => {
          return (
            <RTableRow key={index} mobile={true} row={r} title={props.title} />
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

const row = (title?: boolean) => css`
  display: flex;
  justify-content: space-between;
  flex: 0 0 4.5rem;
  border-bottom: 1px solid black;
  background-color: ${title ? "#007d9c" : "none"};
  color: ${title ? "white" : "none"};
  ${title ? DisableSelect : ""}
  cursor:pointer;
`;

const columnStyle = (width?: string, sortable?: boolean) => css`
  flex: 0 1 ${width ? width : "20rem"};
  padding: 1rem 1rem 1rem 1rem;
  text-align: center;
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

const buttonDiv = css`
  display: inline-block;
  margin: 0.5rem 0.3rem;
`;

export default React.memo(RTable);
