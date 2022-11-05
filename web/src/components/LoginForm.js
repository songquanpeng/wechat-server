import React, { useContext, useEffect, useState } from 'react';
import {
  Button,
  Divider,
  Form,
  Grid,
  Header,
  Image,
  Message,
  Segment,
} from 'semantic-ui-react';
import { Link, useNavigate } from 'react-router-dom';
import { UserContext } from '../context/User';
import { API, showError, showSuccess } from '../helpers';

const LoginForm = () => {
  const [inputs, setInputs] = useState({
    username: '',
    password: '',
  });
  const [submitted, setSubmitted] = useState(false);
  const { username, password } = inputs;
  const [userState, userDispatch] = useContext(UserContext);
  let navigate = useNavigate();

  const [status, setStatus] = useState({});

  useEffect(() => {
    let status = localStorage.getItem('status');
    if (status) {
      status = JSON.parse(status);
      setStatus(status);
    }
  }, []);

  const onGitHubOAuthClicked = () => {
    window.open(
      `https://github.com/login/oauth/authorize?client_id=${status.github_client_id}&scope=user:email`
    );
  };

  function handleChange(e) {
    const { name, value } = e.target;
    setInputs((inputs) => ({ ...inputs, [name]: value }));
  }

  async function handleSubmit(e) {
    setSubmitted(true);
    if (username && password) {
      const res = await API.post('/api/user/login', {
        username,
        password,
      });
      const { success, message, data } = res.data;
      if (success) {
        userDispatch({ type: 'login', payload: data });
        localStorage.setItem('user', JSON.stringify(data));
        navigate('/');
        showSuccess('登录成功！');
      } else {
        showError(message);
      }
    }
  }

  return (
    <Grid textAlign="center" style={{ marginTop: '48px' }}>
      <Grid.Column style={{ maxWidth: 450 }}>
        <Header as="h2" textAlign="center">
          <Image src="/logo.png" /> 用户登录
        </Header>
        <Form size="large">
          <Segment>
            <Form.Input
              fluid
              icon="user"
              iconPosition="left"
              placeholder="用户名"
              name="username"
              value={username}
              onChange={handleChange}
            />
            <Form.Input
              fluid
              icon="lock"
              iconPosition="left"
              placeholder="密码"
              name="password"
              type="password"
              value={password}
              onChange={handleChange}
            />
            <Button fluid size="large" onClick={handleSubmit}>
              登录
            </Button>
          </Segment>
        </Form>
        <Message>
          忘记密码？
          <Link to="/reset" className="btn btn-link">
            点击重置
          </Link>
          ； 没有账户？
          <Link to="/register" className="btn btn-link">
            点击注册
          </Link>
        </Message>
        {status.github_oauth ? (
          <>
            <Divider horizontal>Or</Divider>
            <Button
              circular
              color="black"
              icon="github"
              onClick={onGitHubOAuthClicked}
            />
          </>
        ) : (
          <></>
        )}
      </Grid.Column>
    </Grid>
  );
};

export default LoginForm;
