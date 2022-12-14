import React, { useEffect, useState } from 'react';
import {
  Button,
  Form,
  Grid,
  Header,
  Image,
  Message,
  Segment,
} from 'semantic-ui-react';
import { Link, useNavigate } from 'react-router-dom';
import { API, showError, showInfo, showSuccess } from '../helpers';

const RegisterForm = () => {
  const [inputs, setInputs] = useState({
    username: '',
    password: '',
    password2: '',
    email: '',
    verification_code: '',
  });
  const { username, password, password2 } = inputs;

  const [showEmailVerification, setShowEmailVerification] = useState(false);

  useEffect(() => {
    let status = localStorage.getItem('status');
    if (status) {
      status = JSON.parse(status);
      setShowEmailVerification(status.email_verification);
    }
  });

  let navigate = useNavigate();

  function handleChange(e) {
    const { name, value } = e.target;
    console.log(name, value);
    setInputs((inputs) => ({ ...inputs, [name]: value }));
  }

  async function handleSubmit(e) {
    if (password !== password2) {
      showInfo('两次输入的密码不一致');
      return;
    }
    if (username && password) {
      const res = await API.post('/api/user/register', inputs);
      const { success, message } = res.data;
      if (success) {
        navigate('/login');
        showSuccess('注册成功！');
      } else {
        showError(message);
      }
    }
  }

  const sendVerificationCode = async () => {
    if (inputs.email === '') return;
    const res = await API.get(`/api/verification?email=${inputs.email}`);
    const { success, message } = res.data;
    if (success) {
      showSuccess('验证码发送成功，请检查你的邮箱！');
    } else {
      showError(message);
    }
  };

  return (
    <Grid textAlign="center" style={{ marginTop: '48px' }}>
      <Grid.Column style={{ maxWidth: 450 }}>
        <Header as="h2"  textAlign="center">
          <Image src="/logo.png" /> 新用户注册
        </Header>
        <Form size="large">
          <Segment>
            <Form.Input
              fluid
              icon="user"
              iconPosition="left"
              placeholder="输入用户名"
              onChange={handleChange}
              name="username"
            />
            <Form.Input
              fluid
              icon="lock"
              iconPosition="left"
              placeholder="输入密码"
              onChange={handleChange}
              name="password"
              type="password"
            />
            <Form.Input
              fluid
              icon="lock"
              iconPosition="left"
              placeholder="再次输入密码"
              onChange={handleChange}
              name="password2"
              type="password"
            />
            {showEmailVerification ? (
              <>
                <Form.Input
                  fluid
                  icon="mail"
                  iconPosition="left"
                  placeholder="输入邮箱地址"
                  onChange={handleChange}
                  name="email"
                  type="email"
                  action={
                    <Button onClick={sendVerificationCode}>获取验证码</Button>
                  }
                />
                <Form.Input
                  fluid
                  icon="lock"
                  iconPosition="left"
                  placeholder="输入验证码"
                  onChange={handleChange}
                  name="verification_code"
                />
              </>
            ) : (
              <></>
            )}
            <Button  fluid size="large" onClick={handleSubmit}>
              注册
            </Button>
          </Segment>
        </Form>
        <Message>
          已有账户？
          <Link to="/login" className="btn btn-link">
            点击登录
          </Link>
        </Message>
      </Grid.Column>
    </Grid>
  );
};

export default RegisterForm;
