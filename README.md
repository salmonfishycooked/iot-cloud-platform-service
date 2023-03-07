# 物联网技能赛后端

#### 介绍

物联网技能赛 - 后端



#### 开发环境

Go 1.19.2



#### 使用框架

Gin - Web 后端服务框架

参考文档：https://gin-gonic.com/

Gorm - 数据库操作框架

参考文档：https://gorm.io/



#### 软件架构

- **config** 存放 App 总配置项

- **dao** 数据库持久层

- **controller** 控制层

- **param** 参数层（前端传入的参数结构体映射）

- **model** 模型层（对应数据库里面的表的映射）

- **router** 路由层

- **service** 业务层，由 controller 层分发业务函数去处理数据

- **tcp** 与单片机进行 TCP 连接，实现数据收发功能

- **util** 工具包



#### 运行

1. 下载项目所依赖的第三方模块

```
go get
```

2. 输入 `go run main.go` 运行项目即可（确保已经开启 MySQL 服务）
