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
