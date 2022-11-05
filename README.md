# 微信公众号服务器
> 用以微信公众号的后端，提供登录验证功能

## 功能
+ [ ] 登录验证
+ [ ] 自定义菜单
+ [ ] 自定义回复
+ [ ] Access Token 中控服务器

## 用法
1. 从 [GitHub Releases](https://github.com/songquanpeng/wechat-server/releases/latest) 下载可执行文件.
2. 运行: 
   1. `chmod u+x wechat-server`
   2. `./wechat-server --port 3000`
3. 初始账户用户名为 `root`，密码为 `123456`。

## 配置
系统本身开箱即用。

有一些环境遍历可供配置：
1. `REDIS_CONN_STRING`: 设置之后，将启用 Redis。
   + 例如：`REDIS_CONN_STRING=redis://default:redispw@localhost:49153`
2. `SESSION_SECRET`:设置之后，将使用给定会话密钥。
   + 例如：`SESSION_SECRET=random_string`
3. `SQL_DSN`: 设置之后，将使用目标数据库而非 SQLite。
   + 例如：`SQL_DSN=root:123456@tcp(localhost:3306)/gofile`