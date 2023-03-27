import * as React from 'react';
import styles from '../../styles/Table.module.css';
import { DataGrid, GridColDef } from '@mui/x-data-grid';

interface TodoItem {
  id: number;
  item: string;
}

interface Props {
  todos: TodoItem[];
}

const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'item', headerName: 'Item', width: 140 },
];

export default function BasicTable({ todos }: Props): JSX.Element {
  return (
    <div className={styles.table}>
      <DataGrid rows={todos ? todos : []} columns={columns} checkboxSelection />
    </div>
  );
}
