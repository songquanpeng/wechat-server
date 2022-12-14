import React, { useEffect, useState } from 'react';
import { Form, Grid } from 'semantic-ui-react';
import { API, showError } from '../helpers';

const SystemSetting = () => {
  let [inputs, setInputs] = useState({
    PasswordLoginEnabled: '',
    RegisterEnabled: '',
    EmailVerificationEnabled: '',
    GitHubOAuthEnabled: '',
    GitHubClientId: '',
    GitHubClientSecret: '',
    Notice: '',
    SMTPServer: '',
    SMTPAccount: '',
    SMTPToken: '',
    ServerAddress: '',
    FooterHTML: '',
  });
  let originInputs = {};
  let [loading, setLoading] = useState(false);

  const getOptions = async () => {
    const res = await API.get('/api/option');
    const { success, message, data } = res.data;
    if (success) {
      let newInputs = {};
      data.forEach((item) => {
        newInputs[item.key] = item.value;
      });
      setInputs(newInputs);
      originInputs = newInputs;
    } else {
      showError(message);
    }
  };

  useEffect(() => {
    getOptions().then();
  }, []);

  const updateOption = async (key, value) => {
    setLoading(true);
    switch (key) {
      case 'PasswordLoginEnabled':
      case 'RegisterEnabled':
      case 'EmailVerificationEnabled':
      case 'GitHubOAuthEnabled':
        value = inputs[key] === 'true' ? 'false' : 'true';
        break;
      default:
        break;
    }
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
    if (
      name === 'Notice' ||
      name.startsWith('SMTP') ||
      name === 'ServerAddress' ||
      name === 'GitHubClientId' ||
      name === 'GitHubClientSecret' ||
      name === 'FooterHTML'
    ) {
      setInputs((inputs) => ({ ...inputs, [name]: value }));
    } else {
      await updateOption(name, value);
    }
  };

  const submitNotice = async () => {
    await updateOption('Notice', inputs.Notice);
  };

  const submitServerAddress = async () => {
    let ServerAddress = inputs.ServerAddress;
    if (ServerAddress.endsWith('/')) {
      ServerAddress = ServerAddress.slice(0, ServerAddress.length - 1);
    }
    await updateOption('ServerAddress', ServerAddress);
  };

  const submitSMTP = async () => {
    if (originInputs['SMTPServer'] !== inputs.SMTPServer) {
      await updateOption('SMTPServer', inputs.SMTPServer);
    }
    if (originInputs['SMTPAccount'] !== inputs.SMTPAccount) {
      await updateOption('SMTPAccount', inputs.SMTPAccount);
    }
    if (
      originInputs['SMTPToken'] !== inputs.SMTPToken &&
      inputs.SMTPToken !== ''
    ) {
      await updateOption('SMTPToken', inputs.SMTPToken);
    }
  };

  const submitGitHubOAuth = async () => {
    if (originInputs['GitHubClientId'] !== inputs.GitHubClientId) {
      await updateOption('GitHubClientId', inputs.GitHubClientId);
    }
    if (
      originInputs['GitHubClientSecret'] !== inputs.GitHubClientSecret &&
      inputs.GitHubClientSecret !== ''
    ) {
      await updateOption('GitHubClientSecret', inputs.GitHubClientSecret);
    }
  };

  return (
    <Grid columns={1}>
      <Grid.Column>
        <Form loading={loading}>
          <Form.Group widths="equal">
            <Form.Input
              label="???????????????"
              placeholder="?????????https://yourdomain.com"
              value={inputs.ServerAddress}
              name="ServerAddress"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Button onClick={submitServerAddress}>
            ?????????????????????
          </Form.Button>
          <Form.Group widths="equal">
            <Form.Input
              label="?????? HTML"
              placeholder="???????????????????????????"
              value={inputs.FooterHTML}
              name="FooterHTML"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Button
            onClick={() => {
              updateOption('FooterHTML', inputs.FooterHTML).then();
            }}
          >
            ???????????? HTML
          </Form.Button>
          <Form.Group widths="equal">
            <Form.TextArea
              label="??????"
              placeholder="????????????????????????"
              value={inputs.Notice}
              name="Notice"
              onChange={handleInputChange}
              style={{ minHeight: 150, fontFamily: 'JetBrains Mono, Consolas' }}
            />
          </Form.Group>
          <Form.Button onClick={submitNotice}>????????????</Form.Button>
          <Form.Group inline>
            <Form.Checkbox
              checked={inputs.PasswordLoginEnabled === 'true'}
              label="??????????????????"
              name="PasswordLoginEnabled"
              onChange={handleInputChange}
            />
            <Form.Checkbox
              checked={inputs.RegisterEnabled === 'true'}
              label="?????????????????????"
              name="RegisterEnabled"
              onChange={handleInputChange}
            />
            <Form.Checkbox
              checked={inputs.EmailVerificationEnabled === 'true'}
              label="??????????????????"
              name="EmailVerificationEnabled"
              onChange={handleInputChange}
            />
            <Form.Checkbox
              checked={inputs.GitHubOAuthEnabled === 'true'}
              label="???????????? GitHub ????????????"
              name="GitHubOAuthEnabled"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group widths={3}>
            <Form.Input
              label="SMTP ???????????????"
              name="SMTPServer"
              onChange={handleInputChange}
              autoComplete="off"
              value={inputs.SMTPServer}
            />
            <Form.Input
              label="SMTP ??????"
              name="SMTPAccount"
              onChange={handleInputChange}
              autoComplete="off"
              value={inputs.SMTPAccount}
            />
            <Form.Input
              label="SMTP ????????????"
              name="SMTPToken"
              onChange={handleInputChange}
              type="password"
              autoComplete="off"
              value={inputs.SMTPToken}
            />
          </Form.Group>
          <Form.Button onClick={submitSMTP}>?????? SMTP ??????</Form.Button>
          <Form.Group widths={3}>
            <Form.Input
              label="GitHub Client ID"
              name="GitHubClientId"
              onChange={handleInputChange}
              autoComplete="off"
              value={inputs.GitHubClientId}
            />
            <Form.Input
              label="GitHub Client Secret"
              name="GitHubClientSecret"
              onChange={handleInputChange}
              type="password"
              autoComplete="off"
              value={inputs.GitHubClientSecret}
            />
          </Form.Group>
          <Form.Button onClick={submitGitHubOAuth}>
            ?????? GitHub OAuth ??????
          </Form.Button>
        </Form>
      </Grid.Column>
    </Grid>
  );
};

export default SystemSetting;
