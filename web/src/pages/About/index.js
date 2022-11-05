import React from 'react';
import { Segment, Header } from 'semantic-ui-react';

const About = () => (
  <>
    <Segment>
      <Header as="h3">关于</Header>
      GitHub:{' '}
      <a href="https://github.com/songquanpeng/wechat-server">
        https://github.com/songquanpeng/wechat-server
      </a>
    </Segment>
  </>
);

export default About;
