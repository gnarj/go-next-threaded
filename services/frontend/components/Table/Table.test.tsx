import React from 'react';
import { render, screen } from '@testing-library/react';
import Table from './Table';

interface TodoItem {
  id: number;
  item: string;
}

describe('Table', () => {
  it('should render todos', () => {
    const todos: TodoItem[] = [
      { id: 1, item: 'Item1' },
      { id: 2, item: 'Item2' },
    ];
    render(<Table todos={todos} />);
    todos.forEach((todo) => {
      expect(screen.getByText(todo.item)).toBeInTheDocument();
    });
  });

  it('should handle empty todos prop', () => {
    render(<Table todos={[]} />);
    expect(screen.getByText(/No todos found/i)).toBeInTheDocument();
  });
});
