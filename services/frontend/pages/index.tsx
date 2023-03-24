import Head from 'next/head';
import styles from '../styles/Home.module.css';
import { GetServerSideProps } from 'next';

interface Props {
  status: string;
  username: string;
}

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  const { status } = await fetch('http://localhost:8000/status').then((x) =>
    x.json()
  );
  const { username } = await fetch('http://localhost:8000/username').then((x) =>
    x.json()
  );

  return {
    props: {
      status: status,
      username: username,
    },
  };
};

export default function Home({ status, username }: Props): JSX.Element {
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
        <div>
          Status is: {status}, your username is: {username}
        </div>
      </main>
    </div>
  );
}
