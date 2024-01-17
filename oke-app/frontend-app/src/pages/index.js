import { Table } from 'antd';

const columns = [
  {
    title: 'Name',
    dataIndex: 'Name',
    key: 'Name',
  },
  {
    title: 'Date',
    dataIndex: 'Date',
    key: 'Date',
  },
  {
    title: 'Topics',
    key: 'Topics',
    dataIndex: 'Topics',
  },
  {
    title: 'Presenters',
    key: 'Presenters',
    dataIndex: 'Presenters',
  }
]

export async function getServerSideProps() {
  const res = await fetch("http://" + process.env.API_URL + "/items")
  const data = await res.json()
  return {
    props: {
      data,
    },
  }
}


export default function Home({ data }) {
  return (
    <>
      <h1>OCHaCafe Season8 fast-start golang Demo App</h1>
      <Table dataSource={data} columns={columns} />
    </>
  )
}
