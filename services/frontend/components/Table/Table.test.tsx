import React from 'react';
import { render, screen } from '@testing-library/react';
import Table from './Table';

interface TodoItem {
  id: number;
  item: string;
  editMode: boolean;
}

describe('Table', () => {
  it('should render todos', () => {
    const todos: TodoItem[] = [
      { id: 1, item: 'Item1', editMode: false },
      { id: 2, item: 'Item2', editMode: false },
    ];
    const mockOnTodoUpdate = jest.fn();
    render(<Table todos={todos} onTodoUpdate={mockOnTodoUpdate} />);
    todos.forEach((todo) => {
      expect(screen.getByText(todo.item)).toBeInTheDocument();
    });
  });

  it('should handle empty todos prop', () => {
    const mockOnTodoUpdate = jest.fn();
    render(<Table todos={[]} onTodoUpdate={mockOnTodoUpdate} />);
    expect(screen.getByText(/No todos found/i)).toBeInTheDocument();
  });
});
