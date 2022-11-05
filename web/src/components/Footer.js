import React, { useEffect, useState } from 'react';

import { Container, Segment } from 'semantic-ui-react';

const Footer = () => {
  const [footerHTML, setFooterHTML] = useState('');
  useEffect(() => {
    let savedFooterHTML = localStorage.getItem('footer_html');
    if (!savedFooterHTML) savedFooterHTML = '';
    setFooterHTML(savedFooterHTML);
  });

  return (
    <Segment vertical>
      <Container textAlign="center">
        {footerHTML === '' ? (
          <div className="custom-footer">
            <a
              href="https://github.com/songquanpeng/wechat-server"
              target="_blank"
            >
              微信服务器 {process.env.REACT_APP_VERSION}{' '}
            </a>
            由{' '}
            <a href="https://github.com/songquanpeng" target="_blank">
              JustSong
            </a>{' '}
            构建，源代码遵循{' '}
            <a href="https://opensource.org/licenses/mit-license.php">
              MIT 协议
            </a>
          </div>
        ) : (
          <div
            className="custom-footer"
            dangerouslySetInnerHTML={{ __html: footerHTML }}
          ></div>
        )}
      </Container>
    </Segment>
  );
};

export default Footer;
