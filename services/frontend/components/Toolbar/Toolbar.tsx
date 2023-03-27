import * as React from 'react';
import { useState } from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Unstable_Grid2';
import { addTodo } from '../../utils/api';

export default function TodoInput() {
  const [inputValue, setInputValue] = useState('');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleAddTodo = async () => {
    const todo = await addTodo(inputValue);
    console.log('New todo:', todo);
    setInputValue('');
  };

  return (
    <Box
      component='form'
      sx={{
        '& > :not(style)': { m: 1, width: '25ch' },
      }}
      noValidate
      autoComplete='off'
    >
      <Grid container spacing={2} alignItems='center'>
        <Grid xs={8}>
          <TextField
            id='outlined-basic'
            label='Todo Item'
            variant='outlined'
            value={inputValue}
            onChange={handleInputChange}
          />
        </Grid>
        <Grid xs={4}>
          <Button onClick={handleAddTodo} variant='text'>
            Add
          </Button>
        </Grid>
      </Grid>
    </Box>
  );
}
