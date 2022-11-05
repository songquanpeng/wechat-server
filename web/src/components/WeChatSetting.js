import React, { useEffect, useState } from 'react';
import { Form, Grid } from 'semantic-ui-react';
import { API, showError } from '../helpers';
import { Link } from 'react-router-dom';

const WeChatSetting = () => {
  let [inputs, setInputs] = useState({
    WeChatToken: '',
    WeChatAppID: '',
    WeChatAppSecret: '',
    WeChatEncodingAESKey: '',
    WeChatOwnerID: '',
    WeChatMenu: '',
  });
  let [loading, setLoading] = useState(false);

  const getOptions = async () => {
    const res = await API.get('/api/option');
    const { success, message, data } = res.data;
    if (success) {
      let newInputs = {};
      data.forEach((item) => {
        if (item.key.startsWith('WeChat')) {
          newInputs[item.key] = item.value;
        }
      });
      setInputs(newInputs);
    } else {
      showError(message);
    }
  };

  useEffect(() => {
    getOptions().then();
  }, []);

  const updateOption = async (key, value) => {
    setLoading(true);
    const res = await API.put('/api/option', {
      key,
      value,
    });
    const { success, message } = res.data;
    if (success) {
      setInputs((inputs) => ({ ...inputs, [key]: value }));
    } else {
      showError(message);
    }
    setLoading(false);
  };

  const handleInputChange = async (e, { name, value }) => {
    if (name === 'WeChatMenu') {
      setInputs((inputs) => ({ ...inputs, [name]: value }));
    } else {
      await updateOption(name, value);
    }
  };

  const submitWeChatMenu = async () => {
    await updateOption('WeChatMenu', inputs.WeChatMenu);
  };

  return (
    <Grid columns={1}>
      <Grid.Column>
        <Form loading={loading}>
          <Form.Group widths="equal">
            <Form.Input
              label="令牌（Token）"
              placeholder=""
              value={inputs.WeChatToken}
              name="WeChatToken"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Input
              label="开发者 ID（AppID）"
              placeholder=""
              value={inputs.WeChatAppID}
              name="WeChatAppID"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Input
              label="开发者密码（AppSecret）"
              placeholder=""
              value={inputs.WeChatAppSecret}
              name="WeChatAppSecret"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Input
              label="消息加解密密钥（EncodingAESKey）"
              placeholder=""
              value={inputs.WeChatEncodingAESKey}
              name="WeChatEncodingAESKey"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Input
              label="Root 用户微信 ID"
              placeholder=""
              value={inputs.WeChatOwnerID}
              name="WeChatOwnerID"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.TextArea
              label={
                <p>
                  公众号菜单（
                  <a
                    target="_blank"
                    href="https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html"
                  >
                    格式请参考此处
                  </a>
                  ）
                </p>
              }
              placeholder="JSON 格式"
              value={inputs.WeChatMenu}
              name="WeChatMenu"
              onChange={handleInputChange}
              style={{ minHeight: 150, fontFamily: 'JetBrains Mono, Consolas' }}
            />
          </Form.Group>
          <Form.Button onClick={submitWeChatMenu}>更新公众号菜单</Form.Button>
        </Form>
      </Grid.Column>
    </Grid>
  );
};

export default WeChatSetting;
