export async function addTodo(todoText: string) {
  const response = await fetch('http://localhost:8000/add-todo', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ todoText }),
  });
  console.log(response);
  if (!response.ok) {
    throw new Error('Failed to add todo item');
  }
  const todo = await response.json();
  return todo;
}
