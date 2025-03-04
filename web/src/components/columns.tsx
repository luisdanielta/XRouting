import { ColumnDef } from "@tanstack/react-table";
import { Component } from "@/domain/entities/component";
import { Maintenance } from "@/domain/entities/maintenance";

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

export const MaintenanceColumns: ColumnDef<Maintenance>[] = [
  {
    accessorKey: "description",
    header: "Description",
    enableSorting: true,
  },
  {
    accessorKey: "cost",
    header: "Cost ($)",
    enableSorting: true,
  },
  {
    accessorKey: "date",
    header: "Date",
    enableSorting: true,
  },
]