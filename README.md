Tools 项目，使用go语言编写，使用gorm作为数据库驱动，使用gin作为web框架，使用gormgen作为数据库模型生成工具，使用gormgen作为数据库模型生成工具，使用gormgen作为数据库模型生成工具。

config目录下存放配置文件，develop.yml为应用本地开发配置，app.yml为正式应用配置文件，create_db.sh为创建数据库脚本文件，start-app.sh为更新镜像最新版本脚本文件。

在启动项目运行时之前，先执行create_db.sh脚本文件，创建数据库。

创建数据库后，执行start-app.sh脚本文件，更新镜像最新版本。

# 项目目录
- [1.正式项目目录运行时结构](#1正式项目目录运行时结构)
- [2.快速使用](#2快速使用)
    - [1. 服务器部署](#1-服务器部署)
    - [2. 上传app.yml到config目录下：](#2-上传appyml到config目录下)
    - [3. 创建create_db.sh文件：](#3-创建create_dbsh文件)
    - [4. 编写对应自动跟新镜像版本脚本start.sh文件：](#4-编写对应自动跟新镜像版本脚本startsh文件)
- [3.打包对应的架构镜像以及运行镜像](#3打包对应的架构镜像以及运行镜像)
    - [3.1 在本地打包适合x86_64架构的镜像](#31-在本地打包适合x86_64架构的镜像)
    - [3.2 保存镜像到本地](#32-保存镜像到本地)
    - [3.3 上传镜像到服务器](#33-上传镜像到服务器)
    - [3.4 登录到服务器加载镜像](#34-登录到服务器加载镜像)
    - [3.5 更新最新版本镜像并运行](#35-更新最新版本镜像并运行)


## 1.正式项目目录运行时结构

```
tools/
├── config/
│   └── app.yml
├── script/
│   └── python/   # python脚本文件目录
├── create_db.sh  # 创建数据库脚本文件
└── start-app.sh  # 更新镜像最新版本脚本文件
```

## 2.快速使用
## Commands
```shell
生成SQL ORM
go run main.go generate

数据库建表
go run main.go migrate up

生成API文档
go run main.go swag init
```