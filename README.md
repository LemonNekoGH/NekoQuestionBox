# NekoQuestionBox
这是一个可以自己部署的问题箱，他人可以通过这个问题箱向你发起提问（将来可能允许使用第三方社交账号登录，可以配置是否允许删除提问）。  
使用 PostgreSQL 存储数据（将来可能支持更多数据库类型）。

## 功能
- 接收问题
- 将问题发送给指定的 TG 聊天，并将回复作为答案
## 部署前端
暂无文档

## 部署后端
### 配置文件编写
配置使用 YAML 格式编写。示例如下：
```yaml
port: 5000 # 后端将会使用的端口
database:
  host: localhost # 数据库主机
  port: 15432 # 数据库端口
  username: postgres # 数据库用户名
  password: 123456 # 数据库密码
  database: postgres # 连接时默认连接到的数据库
  sslMode: disable # 禁用 SSL 连接
telegram:
  enabled: true # 启用 TG Bot
  apiToken: <your-token>
  chatId: <your-chat-id>
  sendErrors: true # 遇到错误时发送给目标聊天 id
```
字段解释：
| 字段名 | 字段类型 | 值的范围 | 解释 |
|:---:|:---:|:---:|:---:|
| port | number | 0~65535 | 将会使用的端口 |
| database | Database | - | 数据库配置 |
| telegram | Telegram | - | TG Bot 配置 |

Database 对象：
| 字段名 | 字段类型 | 值的范围 | 解释 |
|:---:|:---:|:---:|:---:|
| host | string | - | 数据库主机地址 |
| username | string | - | 数据库用户名 |
| password | string | - | 数据库密码 |
| port | number | 0~65535 | 数据库端口号 |
| sslMode | string | [disable, required, allow, prefer, verify-ca, verify-full] | SSL 认证方式，[参考](https://www.postgresql.org/docs/current/libpq-ssl.html) |
| database | string | - | 启动时默认连接到的数据库名称 |

Telegram 对象：
| 字段名 | 字段类型 | 值的范围 | 解释 |
|:---:|:---:|:---:|:---:|
| apiToken | string | - | 要使用的 TG Bot 令牌 |
| enabled | boolean | [true, false] | 是否启用 TG Bot |
| chatId | number | - | 新问题要发送到的聊天，[如何获取聊天 id](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id) |
| sendErrors | boolean | [true, false] | 在设置回答时遇到错误时是否将错误发送给目标聊天 |

### 部署
拉取镜像到本地
```shell
$ docker pull ghcr.io/lemonnekogh/qboxb:latest
```
运行
```shell
$ docker run -it -d --name qboxb \
  -v <config-dir>:/etc/qboxb/config \
  -p <port>:<port> \
  ghcr.io/lemonnekogh/qboxb:latest
```

## 感谢名单
- [朵琳（Dolyn157）](https://github.com/Dolyn157)
