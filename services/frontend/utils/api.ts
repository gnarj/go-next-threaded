export async function getTodos() {
  const response = await fetch('http://localhost:8000/todos').then((x) =>
    x.json()
  );
  if (!response) {
    throw new Error('Failed to get todo items');
  }
  return response;
}
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

export async function updateTodo(
  id: number,
  item: string,
  duration: number,
  durationUnit: string
) {
  const response = await fetch('http://localhost:8000/update-todo', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id, item, duration, durationUnit }),
  });
  console.log(response);
  if (!response.ok) {
    throw new Error('Failed to update todo item');
  }
  return response;
}

export async function deleteTodo(id: number) {
  const response = await fetch('http://localhost:8000/delete-todo', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id }),
  });
  if (!response.ok) {
    throw new Error('Failed to delete todo item');
  }
  return response;
}
