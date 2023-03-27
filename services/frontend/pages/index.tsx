import React, { useState } from 'react';
import Head from 'next/head';
import styles from '../styles/Home.module.css';
import { GetServerSideProps } from 'next';
import dynamic from 'next/dynamic';

const Table = dynamic(() => import('../components/Table/Table'), {
  ssr: false,
});
const Toolbar = dynamic(() => import('../components/Toolbar/Toolbar'), {
  ssr: false,
});

interface Props {
  status: string;
  todos: string[];
  username: string;
}

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  const { status } = await fetch('http://localhost:8000/status').then((x) =>
    x.json()
  );
  const { username } = await fetch('http://localhost:8000/username').then((x) =>
    x.json()
  );
  const todos = await fetch('http://localhost:8000/todos').then((x) =>
    x.json()
  );
  return {
    props: {
      status: status,
      todos: todos,
      username: username,
    },
  };
};

export default function Home({ status, todos, username }: Props): JSX.Element {
  const [todoList, setTodoList] = useState<string[]>(todos);

  const handleTodoUpdate = async () => {
    const newTodos = await fetch('http://localhost:8000/todos').then((x) =>
      x.json()
    );
    setTodoList(newTodos);
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <a href='https://nextjs.org'>Next.js!</a>
        </h1>
        <Toolbar onTodoUpdate={handleTodoUpdate} />
        {todoList ? <Table todos={todoList} /> : <h3>Loading....</h3>}
        <div>
          Status is: {status}, your username is: {username}
        </div>
      </main>
    </div>
  );
}
