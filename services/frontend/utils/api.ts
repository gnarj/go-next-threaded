export async function addTodo(todoText: string) {
  const response = await fetch('http://localhost:8000/add-todo', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ todoText }),
  });
  if (!response.ok) {
    throw new Error('Failed to add todo item');
  }
  return response;
}

export async function updateTodo(id: number, item: string) {
  const response = await fetch('http://localhost:8000/update-todo', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id, item }),
  });
  if (!response.ok) {
    throw new Error('Failed to update todo item');
  }
  return response;
}

export async function deleteTodo(todoId: number) {
  const response = await fetch('http://localhost:8000/delete-todo', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ todoId }),
  });
  if (!response.ok) {
    throw new Error('Failed to delete todo item');
  }
  return response;
}
