import * as React from 'react';
import styles from '../../styles/Table.module.css';
import Grid from '@mui/material/Unstable_Grid2';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/DeleteOutlined';
import SaveIcon from '@mui/icons-material/Save';
import CancelIcon from '@mui/icons-material/Close';
import {
  DataGrid,
  GridColDef,
  GridRowSelectionModel,
  GridRowModesModel,
  GridRowModes,
  GridRowId,
} from '@mui/x-data-grid';
import dynamic from 'next/dynamic';
import { Cancel } from '@mui/icons-material';

const TableActions = dynamic(
  () => import('../../components/TableActions/TableActions'),
  {
    ssr: false,
  }
);

interface TodoItem {
  id: number;
  item: string;
  editMode: boolean;
}

interface Props {
  todos: TodoItem[];
  onTodoUpdate: () => void;
}

export default function BasicTable({
  todos,
  onTodoUpdate,
}: Props): JSX.Element {
  const [rowModesModel, setRowModesModel] = React.useState<GridRowModesModel>(
    {}
  );
  const [rowSelectionModel, setRowSelectionModel] =
    React.useState<GridRowSelectionModel>([]);

  // let enableDeleteButton = rowSelectionModel.length > 0;
  // let enableUpdateButton = false;
  const handleEditClick = (id: GridRowId) => () => {
    setRowModesModel({ ...rowModesModel, [id]: { mode: GridRowModes.Edit } });
  };

  const handleCancelClick = (id: GridRowId) => () => {
    setRowModesModel({
      ...rowModesModel,
      [id]: { mode: GridRowModes.View, ignoreModifications: true },
    });
  };

  const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 100 },
    { field: 'item', headerName: 'Item', flex: 1, minWidth: 150 },
    {
      field: 'actions',
      headerName: 'Actions',
      width: 120,
      renderCell: ({ id }) => {
        const isInEditMode = rowModesModel[id]?.mode === GridRowModes.Edit;
        if (isInEditMode) {
          return (
            <div>
              <IconButton aria-label='edit'>
                <SaveIcon />
              </IconButton>
              <IconButton aria-label='delete'>
                <CancelIcon onClick={handleCancelClick(id)} />
              </IconButton>
              ,
            </div>
          );
        }
        return (
          <div>
            <IconButton aria-label='edit'>
              <EditIcon onClick={handleEditClick(id)} />
            </IconButton>
            <IconButton aria-label='delete'>
              <DeleteIcon />
            </IconButton>
          </div>
        );
      },
    },
  ];

  return (
    <Grid container className={styles.tableContainer}>
      {/* <TableActions
        onTodoUpdate={onTodoUpdate}
        enableUpdateButton={enableUpdateButton}
        enableDeleteButton={enableDeleteButton}
      /> */}
      <div className={styles.table}>
        {todos && todos.length > 0 ? (
          <DataGrid
            rows={todos}
            columns={columns}
            checkboxSelection
            onRowSelectionModelChange={(newRowSelectionModel) => {
              setRowSelectionModel(newRowSelectionModel);
            }}
            rowSelectionModel={rowSelectionModel}
          />
        ) : (
          <span>No todos found</span>
        )}
      </div>
    </Grid>
  );
}
