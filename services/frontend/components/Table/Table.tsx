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
  GridEventListener,
  GridRowParams,
  GridRowModesModel,
  GridRowModes,
  GridRowId,
  GridRowsProp,
  GridRowModel,
  MuiEvent,
} from '@mui/x-data-grid';
import { deleteTodo, updateTodo } from '../../utils/api';

interface TodoItem {
  id: number;
  item: string;
}

interface Props {
  todos: TodoItem[];
  handleGetTodos: () => void;
}

export default function Table({ todos, handleGetTodos }: Props): JSX.Element {
  const [rows, setRows] = React.useState<GridRowsProp>(todos);
  const [rowModesModel, setRowModesModel] = React.useState<GridRowModesModel>(
    {}
  );

  React.useEffect(() => {
    if (todos) {
      setRows(todos);
    }
  }, [todos]);

  const handleDeleteClick = async (id: GridRowId) => {
    await deleteTodo(Number(id));
    handleGetTodos();
  };

  const handleEditClick = (id: GridRowId) => () => {
    setRowModesModel({ ...rowModesModel, [id]: { mode: GridRowModes.Edit } });
  };

  const handleCancelClick = (id: GridRowId) => () => {
    setRowModesModel({
      ...rowModesModel,
      [id]: { mode: GridRowModes.View, ignoreModifications: true },
    });
  };

  const handleSaveClick = async (id: GridRowId) => {
    setRowModesModel({ ...rowModesModel, [id]: { mode: GridRowModes.View } });
  };

  const handleRowEditStart = (
    params: GridRowParams,
    event: MuiEvent<React.SyntheticEvent>
  ) => {
    event.defaultMuiPrevented = true;
  };

  const handleRowEditStop: GridEventListener<'rowEditStop'> = (
    params,
    event
  ) => {
    event.defaultMuiPrevented = true;
  };

  const handleRowModesModelChange = (newRowModesModel: GridRowModesModel) => {
    setRowModesModel(newRowModesModel);
  };

  const processRowUpdate = (newRow: GridRowModel) => {
    const { id, item } = newRow;
    const updatedRow = { ...newRow, isNew: false };
    const updatedRows = rows.map((row) =>
      row.id === newRow.id ? updatedRow : row
    );
    setRows(updatedRows);
    updateTodo(id, item);
    return updatedRow;
  };

  const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', type: 'number', width: 40 },
    {
      field: 'item',
      headerName: 'Item',
      editable: true,
      flex: 1,
      minWidth: 150,
    },
    {
      field: 'actions',
      headerName: 'Actions',
      width: 120,
      renderCell: ({ id }) => {
        const isInEditMode = rowModesModel[id]?.mode === GridRowModes.Edit;
        if (isInEditMode) {
          return (
            <div>
              <IconButton onClick={() => handleSaveClick(id)} aria-label='save'>
                <SaveIcon />
              </IconButton>
              <IconButton onClick={handleCancelClick(id)} aria-label='delete'>
                <CancelIcon />
              </IconButton>
            </div>
          );
        }
        return (
          <div>
            <IconButton onClick={handleEditClick(id)} aria-label='edit'>
              <EditIcon />
            </IconButton>
            <IconButton
              onClick={() => handleDeleteClick(id)}
              aria-label='delete'
            >
              <DeleteIcon />
            </IconButton>
          </div>
        );
      },
    },
  ];

  return (
    <Grid container className={styles.tableContainer}>
      <div className={styles.table}>
        {rows && rows.length > 0 ? (
          <DataGrid
            editMode='row'
            rows={rows}
            columns={columns}
            onRowEditStart={handleRowEditStart}
            onRowEditStop={handleRowEditStop}
            rowModesModel={rowModesModel}
            onRowModesModelChange={handleRowModesModelChange}
            processRowUpdate={processRowUpdate}
          />
        ) : (
          <span>No todos found</span>
        )}
      </div>
    </Grid>
  );
}
