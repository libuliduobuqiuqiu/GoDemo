# go-demo
![GitHub Created At](https://img.shields.io/github/created-at/libuliduobuqiuqiu/GoDemo)
![GitHub last commit](https://img.shields.io/github/last-commit/libuliduobuqiuqiu/GoDemo)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/libuliduobuqiuqiu/GoDemo)
![GitHub License](https://img.shields.io/github/license/libuliduobuqiuqiu/GoDemo)


## 项目简介
GoDemo 是个人 Go 知识库：按主题整理笔记，用代码和测试验证重点知识，并保留数据库、Web、RPC 等场景的实践示例。使用时可以先从知识库目录找结论，再沿文章中的链接查看代码和运行方式。

## 知识库导航

从[知识库首页](docs/README.md)按基础、进阶、框架、面试四个方向查找专题；每个目录都有文章索引和对应代码入口。新增笔记按[收纳与写作约定](docs/knowledge-guide.md)登记。

后续整理顺序与完成标准见[知识库待办](TODO.md)。

## 目录结构

- cmd/: 执行程序的入口以及自动生成代码。
- configs/: 配置文件目录，如数据库,服务器配置。
- docs/: 知识库入口与按主题整理的笔记。
- internal/：存放各种库的使用示例、算法实现、以及实际业务逻辑等。
- pkg/: 工具函数包，可重复使用。
- test/: 自动测试与需单独运行的手动演示。
- web/: Web框架的应用示例和模板。

## 环境要求
运行此项目需要：
- Go 1.25（以当前 `go.mod` 为准）
- 对Go语言有基础了解

## 快速开始
### 克隆代码
```bash
git clone https://github.com/libuliduobuqiuqiu/GoDemo.git
cd GoDemo
```
### 运行示例
查看命令行入口：
```bash
go run . --help
```

Gin 本机路由与代理的启动方式、路由和运行边界见[Gin 路由与代理](docs/framework/go_gin.md)。
### 执行测试
在项目根目录执行默认测试，无需数据库、设备或私有配置：
```bash
go build ./...
go test ./...
```

需要外部服务、固定文件或交互的演示使用 `integration` 构建标签，详见[测试与演示说明](docs/testing.md)。

数据库、SSH 等演示需要配置时，复制 `configs/conf_demo.json` 到已被 Git 忽略的 `configs/conf.json`，填入本机连接信息，然后设置 `GO_DEMO_CONFIG` 指向该文件。也可以在代码中显式传入配置路径；显式路径优先。默认构建和测试无需配置文件。

项目整理过程与验证结果见[整理记录](docs/maintenance-log.md)。

## 贡献指南

欢迎贡献代码！您可以通过提交 Issue 或 Pull Request 为项目做出贡献。

## 许可证
Released under the [MIT License](LICENSE).
