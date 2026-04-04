import type { ReactNode } from 'react';

interface Column<T> {
  key: string;
  header: string;
  render: (row: T) => ReactNode;
  className?: string;
}

interface DataTableProps<T> {
  columns: Column<T>[];
  data: T[];
  keyExtractor: (row: T) => string;
}

export function DataTable<T>({ columns, data, keyExtractor }: DataTableProps<T>) {
  return (
    <div className="overflow-x-auto rounded-2xl bg-surface-container-lowest shadow-xl border border-outline-variant/10">
      <table className="w-full text-left text-sm">
        <thead className="border-b border-outline-variant/20 bg-surface-container-low text-[11px] uppercase tracking-wider text-outline">
          <tr>
            {columns.map((col) => (
              <th key={col.key} className={`px-6 py-4 font-bold ${col.className ?? ''}`}>
                {col.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="divide-y divide-outline-variant/10">
          {data.map((row, index) => (
            <tr
              key={keyExtractor(row)}
              className={`hover:bg-primary/5 transition-colors ${index % 2 !== 0 ? 'bg-surface-container-low/30' : ''}`}
            >
              {columns.map((col) => (
                <td key={col.key} className={`px-6 py-5 ${col.className ?? ''}`}>
                  {col.render(row)}
                </td>
              ))}
            </tr>
          ))}
          {data.length === 0 && (
            <tr>
              <td
                colSpan={columns.length}
                className="px-6 py-12 text-center text-on-surface-variant"
              >
                No data available
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
}
