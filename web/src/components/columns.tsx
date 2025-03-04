import { ColumnDef } from "@tanstack/react-table";
import { Component } from "@/domain/entities/component";

export const ComponentColumns: ColumnDef<Component>[] = [
  {
    accessorKey: "type",
    header: "Type",
    enableSorting: true,
  },
  {
    accessorKey: "status",
    header: "Status",
    enableSorting: true,
  },
  {
    accessorKey: "category",
    header: "Category",
    enableSorting: false,
  },
];
