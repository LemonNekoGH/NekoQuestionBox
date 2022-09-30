# NekoQuestionBox
这是一个可以自己部署的问题箱，他人可以通过这个问题箱向你发起提问（将来可能允许使用第三方社交账号登录，可以配置是否允许删除提问）。  
使用 PostgreSQL 存储数据（将来可能支持更多数据库类型）。

## 开始部署
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
  sslMode: disable # SSL 认证方式
```
字段解释：
| 字段名 | 字段类型 | 值的范围 | 解释 |
|:---:|:---:|:---:|:---:|
| port | number | 0~65535 | 将会使用的端口 |
| database | Database | - | 数据库配置 |

Database 对象：
| 字段名 | 字段类型 | 值的范围 | 解释 |
|:---:|:---:|:---:|:---:|
| host | string | - | 数据库主机地址 |
| username | string | - | 数据库用户名 |
| password | string | - | 数据库密码 |
| port | number | 0~65535 | 数据库端口号 |
| sslMode | string | [disable, required, allow, prefer, verify-ca, verify-full] | SSL 认证方式，[参考](https://www.postgresql.org/docs/current/libpq-ssl.html) |
| database | string | - | 启动时默认连接到的数据库名称 |

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
