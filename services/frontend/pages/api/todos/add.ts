import { NextApiRequest, NextApiResponse } from 'next';
import { addTodo } from '../../../utils/api';

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method === 'POST') {
    const { todoText } = req.body;
    const todo = await addTodo(todoText);
    res.status(200).json(todo);
  } else {
    res.status(400).json({ message: 'Invalid request method' });
  }
}
