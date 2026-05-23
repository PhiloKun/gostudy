# Go Study

<div align="center">

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-008ECF?logo=gin&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

**从零开始的 Go 语言学习之旅 —— 基础语法 -> 网络编程 -> Web 框架**

</div>

## 关于本项目

本项目是一套完整的 Go 语言学习笔记与实战代码集合，面向 Go 语言初学者和进阶者。内容从最基础的变量定义起步，逐步深入到协程并发、网络编程，再到流行的 Gin Web 框架实战，覆盖了 Go 后端开发的核心知识体系。

### 适合人群

- **Go 初学者** -- 从零开始系统学习 Go 语言语法和编程范式
- **Web 开发者** -- 希望快速掌握 Gin 框架进行 Web 开发
- **后端工程师** -- 需要 TCP/HTTP 网络编程的实践参考

### 内容特色

- **循序渐进** -- 三个阶段层层递进，从语言基础到实际框架应用
- **代码即笔记** -- 每个知识点配有可直接运行的示例代码
- **注释详尽** -- 代码中包含详细的中文注释，便于理解

## 目录结构

```
├── go基础/              # Go 语言基础知识
│   ├── 变量定义.go           # 变量声明与初始化
│   ├── 基本数据类型.go        # 基本数据类型
│   ├── 数组.go               # 数组
│   ├── 切片.go               # 切片
│   ├── map.go                # 字典
│   ├── 函数.go               # 函数定义
│   ├── 结构体.go             # 结构体
│   ├── 结构体指针.go         # 结构体指针
│   ├── 结构体tag.go          # 结构体标签
│   ├── 接口.go               # 接口
│   ├── 闭包.go               # 闭包
│   ├── 协程.go               # goroutine 协程
│   ├── defer函数.go          # defer 延迟调用
│   ├── if.go                 # if 条件判断
│   ├── switch.go             # switch 分支
│   ├── for 循环.go           # for 循环
│   ├── 自定义类型和类型别名.go  # 自定义类型与类型别名
│   ├── 输入.go               # 输入
│   ├── 输出.go               # 输出
│   ├── init函数.go           # init 初始化函数
│   └── version/              # 版本相关
│
├── gin_study/           # Gin 框架学习
│   ├── 初始gin.go            # Gin 入门
│   ├── 原生http服务器/        # Go 原生 HTTP 服务器
│   ├── 请求/                 # 请求处理
│   │   ├── 1.查询参数.go     # URL 查询参数
│   │   ├── 2.动态参数.go     # 动态路由参数
│   │   ├── 3.form表单.go     # 表单提交
│   │   ├── 4.文件上传.go     # 单文件上传
│   │   ├── 5.多文件上传.go   # 多文件上传
│   │   └── 6.原始内容.go     # 原始请求体
│   └── 响应/                 # 响应处理
│       ├── 1.响应json.go     # JSON 响应
│       ├── 2.响应html.go     # HTML 模板渲染
│       ├── 3.响应文件.go     # 文件响应
│       └── 4.静态文件.go     # 静态文件服务
│
└── 网络编程/            # 网络编程
    ├── tcp/
    │   ├── tcp服务端.go      # TCP 服务端
    │   └── tcp客户端.go      # TCP 客户端
    └── http_study/
        ├── 服务端.go         # HTTP 服务端
        └── 客户端.go         # HTTP 客户端
```

## 快速开始

### 环境要求

- Go 1.21+
- Gin 框架（gin_study 部分需要）

### 运行示例

```bash
# 运行 Go 基础示例
cd "go基础"
go run 变量定义.go

# 运行 Gin 示例
cd gin_study
go run 初始gin.go

# 运行 TCP 示例
cd "网络编程/tcp"
# 先启动服务端
go run tcp服务端.go
# 再启动客户端
go run tcp客户端.go
```

## 学习路线

1. **Go 基础** -- 从变量定义到协程并发，打好语言基础
2. **网络编程** -- 理解 TCP/HTTP 协议基础
3. **Gin 框架** -- 掌握路由、请求处理、响应渲染、文件处理等 Web 开发技能

## License

MIT
