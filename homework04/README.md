# homework04

本次作业要求你使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

## 接口清单

- ✅ 注册
- ✅ 登录
- ✅ 获取用户信息
- ✅ 修改用户信息
- ✅ 创建文章
- ✅ 查询文章列表
- ✅ 查询文章详情
- ✅ 修改文章
- ✅ 删除文章
- ✅ 创建评论
- ✅ 根据postID查询评论

## 项目结构

```
project/
├── main.go              # 程序入口
├── config/              # 配置管理
│   └── config.go
|   └── config.yaml
├── handler/             # 处理器（Controller）
│   └── user_handler.go
|   └── post_handler.go
|   └── comment_handler.go
├── middleware/          # 中间件
│   ├── auth.go
│   ├── logger.go
│   └── cors.go
├── models/              # 数据模型
│   └── user.go
│   └── post.go
│   └── comment.go
├── services/            # 业务逻辑层
│   └── user_service.go
│   └── post_service.go
│   └── comment_service.go
└── utils/               # 工具函数
    ├── db.go
    ├── jwt.go
    └── response.go
```

## 快速开始

### 1. 安装依赖

```bash
cd D:/projects/blockchain/homework/blockchain_learning/homework04
go mod tidy
```

### 2. 运行项目

```bash
go run main.go
```

服务器将在 `http://0.0.0.0:8888` 启动。

### 3. 测试 API

#### 3.1.1注册

```cmd
curl -X POST http://localhost:8888/register ^
-H "Content-Type: application/json" ^
-d "{ \"userName\": \"tom\", \"passWord\": \"123456\", \"email\": \"tom@qq.com\" }"
```

#### 3.1.2登录

```cmd
curl -X POST http://localhost:8888/login ^
-H "Content-Type: application/json" ^
-d "{ \"userName\": \"tom\", \"passWord\": \"123456\"}"
```

#### 3.2.1获取用户信息（需要 Token）

登录后会返回 token，将 `YOUR_TOKEN` 替换为实际的 token：

```cmd
curl -X GET http://localhost:8888/user/1 ^
-H "Authorization: Bearer YOUR_TOKEN"
```

#### 3.2.2更新用户信息（需要 Token）

```cmd
curl -X PUT http://localhost:8888/user ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer YOUR_TOKEN" ^
-d "{ \"id\": 1,\"userName\": \"jack\", \"passWord\": \"666666\",\"email\": \"jack@qq.com\"}"
```

#### 3.3.1创建文章（需要 Token）

```cmd
curl -X POST http://localhost:8888/post ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer YOUR_TOKEN" ^
-d "{ \"title\": \"标题1\",\"content\": \"内容1\"}"
```

#### 3.3.2查询文章列表（需要 Token）

```cmd
curl -X GET http://localhost:8888/post/title/%E6%96%87%E7%AB%A0%21%40%23%24%2F/current/1/size/1 ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 3.3.3查询文章详情（需要 Token）

```cmd
curl -X GET http://localhost:8888/post/1 ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 3.3.4修改文章（需要 Token）

```cmd
curl -X PUT http://localhost:8888/post ^
-H "Content-Type: application/json" ^
-H "Authorization: Bearer YOUR_TOKEN" ^
-d "{ \"id\": 1,\"title\": \"标题修改\", \"content\": \"内容修改\"}"
```

#### 3.3.5删除文章（需要 Token）

```cmd
curl -X DELETE http://localhost:8888/post ^
-H "Content-Type: application/json" ^
-H "Authorization: Bearer YOUR_TOKEN" ^
-d "{ \"id\": 1}"
```

#### 3.4.1创建评论（需要 Token）

```cmd
curl -X POST http://localhost:8888/comment ^
-H "Content-Type: application/json" ^
-H "Authorization: Bearer YOUR_TOKEN" ^
-d "{ \"postID\": 1, \"content\": \"对文章1的评论\" }"
```

#### 3.4.2根据postID查询评论（需要 Token）

```cmd
curl -X GET http://localhost:8888/comment/1 ^
-H "Authorization: Bearer YOUR_TOKEN"
```

#### 健康检查

```cmd
curl -X GET http://localhost:8888/health
```

## API 端点

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/register` | 注册 | 否 |
| POST | `/login` | 登录 | 否 |
| GET | `/user/:id` | 获取用户信息 | 是 |
| PUT | `/user` | 修改用户信息 | 是 |
| POST | `/post` | 创建文章 | 是 |
| GET | `/post/title/:title/current/:current/size/:size` | 查询文章列表 | 是 |
| GET | `/post/:id` | 查询文章详情 | 是 |
| PUT | `/post` | 修改文章 | 是 |
| DELETE | `/post` | 删除文章 | 是 |
| POST | `/comment` | 创建评论 | 是 |
| GET | `/comment` | 根据postID查询评论 | 是 |
| GET | `/health` | 健康检查 | 否 |

## 技术栈

- **Web 框架**: Gin v1.10.0
- **ORM**: GORM v1.25.12
- **数据库**: MySQL
- **认证**: JWT (github.com/golang-jwt/jwt/v5)
- **密码加密**: bcrypt (golang.org/x/crypto)

