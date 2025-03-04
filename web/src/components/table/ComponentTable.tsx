import { flexRender } from "@tanstack/react-table";
import { useComponentTable } from "@/hooks/useComponentTable";
import { Component } from "@/domain/entities/component";
import PrimaryButton from "../ui/button/primaryButton";

interface UserTableProps {
  data: Component[];
}

export const ComponentTable: React.FC<UserTableProps> = ({ data }) => {
  const table = useComponentTable(data);

  return (
    <div className="overflow-hidden rounded-md border border-gray-200 shadow-md shadow-gray-100">
      <table className="w-full table-fixed border-collapse">
        <thead>
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id} className="bg-gradient-to-r from-blue-400 via-green-300 to-green-400 text-white">
              {headerGroup.headers.map((header) => (
                <th
                  key={header.id}
                  className="px-4 py-3 text-left font-semibold uppercase tracking-wider cursor-pointer hover:bg-green-500 transition-all duration-200 whitespace-nowrap w-1/5"
                  onClick={header.column.getToggleSortingHandler()}
                >
                  {flexRender(header.column.columnDef.header, header.getContext())}
                  {header.column.getIsSorted() === "asc" ? " 🔼" : header.column.getIsSorted() === "desc" ? " 🔽" : ""}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody className="divide-y divide-gray-200 bg-white">
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id} className="hover:bg-gray-100 transition-colors">
              {row.getVisibleCells().map((cell) => (
                <td
                  key={cell.id}
                  className="px-4 py-3 truncate whitespace-nowrap w-1/5"
                >
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>

      <div className="flex justify-between items-center px-4 py-3 bg-gray-100 border-t border-gray-200">
        <span>
          <PrimaryButton
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
            text="Anterior"
          />
        </span>
        <span>
          Página <strong>{table.getState().pagination.pageIndex + 1}</strong> de{" "}
          <strong>{table.getPageCount()}</strong>
        </span>
        <span>
          <PrimaryButton
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
            text="Siguiente"
          />
        </span>
      </div>
    </div>
  );
};
