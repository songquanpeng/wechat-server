import React, { useEffect, useState } from 'react';
import { Form, Grid } from 'semantic-ui-react';
import { API, showError } from '../helpers';

const WeChatSetting = () => {
  let [inputs, setInputs] = useState({
    WeChatToken: '',
    WeChatAppID: '',
    WeChatAppSecret: '',
    WeChatEncodingAESKey: '',
    WeChatOwnerID: '',
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
    await updateOption(name, value);
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
        </Form>
      </Grid.Column>
    </Grid>
  );
};

export default WeChatSetting;
