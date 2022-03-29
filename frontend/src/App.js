import { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';
import { Table, Tabs } from 'antd';
import 'antd/dist/antd.min.css';
import Layout from 'antd/lib/layout/layout';
import {faker} from '@faker-js/faker'

const { TabPane } = Tabs;

const columns = [
  {
    title: 'Ip Address',
    dataIndex: 'IpAddress',
    key: 'IpAddress',
    render: text => <a>{text}</a>,
  },
  {
    title: 'Visit',
    dataIndex: 'Visit',
    key: 'Visit',
  },
];

const columns2 = [
  {
    title: 'Ip Address',
    dataIndex: 'IpAddress',
    key: 'IpAddress',
    render: text => <a>{text}</a>,
  },
  {
    title: 'Timezone',
    dataIndex: 'Timezone',
    key: 'Timezone',
  },
  {
    title: 'Email',
    dataIndex: 'Email',
    key: 'Email',
  },
  {
    title: 'Latitude',
    dataIndex: 'Latitude',
    key: 'Latitude',
  },
  {
    title: 'Longitude',
    dataIndex: 'Longitude',
    key: 'Longitude',
  },
];

function App() {
  const [latestVistors, setLatestVisitor] = useState([]);
  const [mostVisitUser, setMostVisitUser] = useState([]);

  useEffect(() => {
    navigator.geolocation.getCurrentPosition(({coords}) => {
      const {latitude, longitude} = coords;
      const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
      axios.post('/api/events', {
        "Email": faker.internet.email(),
        "Timezone": timezone,
        "Longitude": longitude,
        "Latitude": latitude,
      });
    })
    axios.get('/api/stats').then((res) => {
      const {LatestVisit, MostVisit} = res.data
      setMostVisitUser(MostVisit);
      setLatestVisitor(LatestVisit);
    });
  }, []);

  return (
    <Layout style={{ padding: '0 24px 24px' }}>
      <Tabs defaultActiveKey="1">
        <TabPane tab="Most Visit" key="1">
        <Table columns={columns} dataSource={mostVisitUser} />
        </TabPane>
        <TabPane tab="Latest Visit" key="2">
        <Table columns={columns2} dataSource={latestVistors} />
        </TabPane>
      </Tabs>
    </Layout>
  );
}

export default App;
