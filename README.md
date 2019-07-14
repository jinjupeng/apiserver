# 基于Go语言构建企业级的RESTful API服务

## 项目目录结构

```tree
├── admin.sh                     # 进程的start|stop|status|restart控制文件
├── conf                         # 配置文件统一存放目录
│   ├── config.yaml              # 配置文件
│   ├── server.crt               # TLS配置文件
│   └── server.key
├── config                       # 专门用来处理配置和配置文件的Go package
│   └── config.go
├── db.sql                       # 在部署新环境时，可以登录MySQL客户端，执行source db.sql创建数据库和表
├── docs                         # swagger文档，执行 swag init 生成的
│   ├── docs.go
│   └── swagger
│       ├── swagger.json
│       └── swagger.yaml
├── handler                      # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── handler.go
│   ├── sd                       # 健康检查handler
│   │   └── check.go 
│   └── v1                       # API接口版本
|       └── user                 # 核心：用户业务逻辑handler
│           ├── create.go        # 新增用户
│           ├── delete.go        # 删除用户
│           ├── get.go           # 获取指定的用户信息
│           ├── list.go          # 查询用户列表
│           ├── login.go         # 用户登录
│           ├── update.go        # 更新用户
│           └── user.go          # 存放用户handler公用的函数、结构体等
├── main.go                      # Go程序唯一入口
├── Makefile                     # Makefile文件，一般大型软件系统都是采用make来作为编译工具
├── model                        # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   ├── init.go                  # 初始化和连接数据库
│   ├── model.go                 # 存放一些公用的go struct
│   └── user.go                  # 用户相关的数据库CURD操作
├── pkg                          # 引用的包
│   ├── auth                     # 认证包
│   │   └── auth.go
│   ├── constvar                 # 常量统一存放位置
│   │   └── constvar.go
│   ├── errno                    # 错误码存放位置
│   │   ├── code.go
│   │   └── errno.go
│   ├── token
│   │   └── token.go
│   └── version                  # 版本包
│       ├── base.go
│       ├── doc.go
│       └── version.go
├── README.md                    # API目录README
├── router                       # 路由相关处理
│   ├── middleware               # API服务器用的是Gin Web框架，Gin中间件存放位置
│   │   ├── auth.go 
│   │   ├── header.go
│   │   ├── logging.go
│   │   └── requestid.go
│   └── router.go
├── service                      # 实际业务处理函数存放位置
│   └── service.go
├── util                         # 工具类函数存放目录
│   ├── util.go
│   └── util_test.go
└── vendor                         # vendor目录用来管理依赖包
    ├── github.com
    ├── golang.org
    ├── gopkg.in
    └── vendor.json
```

在Go API项目中，一般都会包括这些功能项：Makefile文件、配置文件目录、RESTful API服务器的handler目录、model目录、工具类目录、vendor目录，以及实际处理业务逻辑函数所存放的service目录。

这些都在上面的代码结构中列出，新加功能时将代码放入对应功能的目录/文件中以使整个项目代码结构更加清晰，非常有利于后期的查找和维护。

## 注意

1. 在golang1.12中，go tool vet已经转为go vet

```bash
go tool vet main.go
# 错误：vet: invoking "go tool vet" directly is unsupported; use "go vet"
```
