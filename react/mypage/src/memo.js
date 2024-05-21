import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  EditOutlined,
  BgColorsOutlined,
  DeleteOutlined,
} from '@ant-design/icons';
import { Input, Layout, Card, Row, Col } from 'antd';
import { refreshToken } from './store/module/tokenStore';
import { useDispatch, useSelector } from 'react-redux';
import { fetchMemoList } from './store/module/memoStore';

const { Meta } = Card;
const { Header, Content } = Layout;

export default function App() {
  const [loginStatus, setLoginStatus] = useState(false);
  const navigate = useNavigate();
  const dispatch = useDispatch();
  useEffect(() => {
    const run = async () => {
      if (!loginStatus) {
        await refreshToken(
          () => setLoginStatus(true),
          () => navigate('/login')
        );
      }
      dispatch(fetchMemoList());
      return () => {};
    };

    run();
  }, [dispatch]);

  const memoList = useSelector((state) => state.memo.memoList);
  const onChange = (e) => {};
  return (
    <Layout>
      <Header
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <Input
          style={{ minWidth: '240px', width: '50%' }}
          showCount
          maxLength={20}
          onChange={onChange}
          placeholder='请输入要做的事？'
        />
      </Header>
      <Content style={{ padding: '0 48px', margin: '24px 16px' }}>
        <Row gutter={24}>
          {memoList.map((item, index) => {
            const key = `${item.id}`;
            return (
              <Col
                style={{ margin: '4px 8px 20px 8px' }}
                key={key}
                xs={{
                  flex: '100%',
                }}
                sm={{
                  flex: '50%',
                }}
                md={{
                  flex: '40%',
                }}
                lg={{
                  flex: '20%',
                }}
                xl={{
                  flex: '10%',
                }}
              >
                <Card
                  style={{
                    width: 240,
                    padding: '8 0',
                  }}
                  hoverable='true'
                  size='small'
                  cover={
                    <img
                      alt='example'
                      src='https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png'
                    />
                  }
                  actions={[
                    <DeleteOutlined key='delete' />,
                    <EditOutlined key='edit' />,
                    <BgColorsOutlined
                      key='pic'
                      onClick={() => navigate('/nav')}
                    />,
                  ]}
                >
                  <Meta title={item.content} />
                </Card>
              </Col>
            );
          })}
        </Row>
      </Content>
    </Layout>
  );
}
