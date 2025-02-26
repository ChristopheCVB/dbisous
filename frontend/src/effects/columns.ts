import { Effect } from "effect";
import { client } from "_/go/models";
import { cell } from "@/components/database/table/table";
import { TableColumn, TableData } from "@nuxt/ui/dist/module";

function formatColumns(columns: client.ColumnMetadata[]) {
  const formatted = columns.map(
    ({ name, type, default_value: defaultValue, nullable }) =>
      ({
        accessorKey: name,
        header: name,
        cell: cell({ type, defaultValue, nullable, disabled: false }),
      }) as TableColumn<TableData>,
  );

  formatted.push({
    accessorKey: "action",
    header: "Actions",
  });

  return formatted;
}

export function formatQueryResult(result: client.QueryResult) {
  return Effect.succeed({
    ...result,
    columns: formatColumns(result.columns),
  });
}
