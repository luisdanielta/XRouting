import {
  useReactTable,
  getCoreRowModel,
  getSortedRowModel,
  SortingState,
  getPaginationRowModel,
} from "@tanstack/react-table";
import { useState } from "react";
import { ComponentColumns } from "@/components/columns";
import { Component } from "@/domain/entities/component";

export const useComponentTable = (data: Component[]) => {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [pageIndex, setPageIndex] = useState(0);
  const pageSize = 10;

  const table = useReactTable({
    data,
    columns: ComponentColumns,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    state: { sorting, pagination: { pageIndex, pageSize } },
    onSortingChange: setSorting,
    onPaginationChange: (updater) => {
      if (typeof updater === "function") {
        setPageIndex(updater({ pageIndex, pageSize }).pageIndex);
      } else {
        setPageIndex(updater.pageIndex);
      }
    },
  });

  return table;
};
