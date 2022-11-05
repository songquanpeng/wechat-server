# 微信公众号服务器
> 用以微信公众号的后端，提供登录验证功能

## 功能
+ [x] Access Token 自动刷新 & 提供外部访问接口
+ [ ] 自定义菜单
+ [ ] 登录验证
+ [ ] 自定义回复

## 配置
1. 从 [GitHub Releases](https://github.com/songquanpeng/wechat-server/releases/latest) 下载可执行文件。
2. 系统本身开箱即用，有一些环境变量可供配置：
   1. `REDIS_CONN_STRING`: 设置之后，将启用 Redis。
      + 例如：`REDIS_CONN_STRING=redis://default:redispw@localhost:49153`
   2. `SESSION_SECRET`:设置之后，将使用给定会话密钥。
      + 例如：`SESSION_SECRET=random_string`
   3. `SQL_DSN`: 设置之后，将使用目标数据库而非 SQLite。
      + 例如：`SQL_DSN=root:123456@tcp(localhost:3306)/gofile`
3. 运行: 
   1. `chmod u+x wechat-server`
   2. `./wechat-server --port 3000`
4. 初始账户用户名为 `root`，密码为 `123456`，记得登录后立刻修改密码。
5. 前往[微信公众号配置页面 -> 设置与开发 -> 基本配置](https://mp.weixin.qq.com/)获取 AppID 和 AppSecret，并在我们的配置页面填入上述信息，另外还需要配置 IP 白名单，按照页面上的提示完成即可。
6. 前往[微信公众号配置页面 -> 设置与开发 -> 基本配置](https://mp.weixin.qq.com/)填写以下配置：
   1. `URL` 填：`https://<your.domain>/api/wechat_verification`
   2. `Token` 首先在我们的配置页面随便填写一个 Token，然后在微信公众号的配置页面填入同一个 Token 即可。
   3. `EncodingAESKey` 点随机生成，然后在我们的配置页面填入该值。
   4. 消息加解密方式选择明文模式。
7. 之后保存设置并启用设置。

## API
### 获取 Access Token
1. 请求方法：`GET`
2. URL：`/api/access_token`
3. 无参数，但是需要设置 HTTP 头部：`Authorization: your token`