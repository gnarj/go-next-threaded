import * as React from 'react';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Unstable_Grid2';
import DeleteIcon from '@mui/icons-material/Delete';
import SendIcon from '@mui/icons-material/Send';
import { updateTodo, deleteTodo } from '../../utils/api';

interface Props {
  onTodoUpdate: () => void;
  enableDeleteButton: boolean;
  enableUpdateButton: boolean;
}

export default function TodoInput({
  onTodoUpdate,
  enableDeleteButton,
  enableUpdateButton,
}: Props): JSX.Element {
  const handleUpdateTodo = async () => {
    // await updateTodo();
    onTodoUpdate();
  };

  const handleDeleteTodo = async () => {
    // await deleteTodo();
    onTodoUpdate();
  };

  return (
    <Grid container spacing={2} alignItems='center' justifyContent='flex-start'>
      <Grid xs={6}>
        <Button
          disabled={!enableUpdateButton}
          variant='outlined'
          onClick={handleUpdateTodo}
          startIcon={<SendIcon />}
        >
          Update
        </Button>
      </Grid>
      <Grid xs={6}>
        <Button
          disabled={!enableDeleteButton}
          variant='outlined'
          onClick={handleDeleteTodo}
          startIcon={<DeleteIcon />}
          color='error'
        >
          Delete
        </Button>
      </Grid>
    </Grid>
  );
}
