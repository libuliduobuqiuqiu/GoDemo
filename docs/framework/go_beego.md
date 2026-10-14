# Beego 与仓库中的 WebSocket 聊天室

状态：已复核仓库依赖与代码入口；聊天室仅做静态检查，未完成端到端运行验收。仓库使用 `github.com/beego/beego/v2 v2.1.6`。

## Beego 在这个仓库里做什么

[聊天室路由](../../internal/goweb/gowebsockets/chat/chatrouter.go)在包初始化时调用 `web.CtrlGet("/chat/join", (*ChatController).Join)`，将 GET 请求映射到控制器方法。`Join` 读取用户名并尝试升级 WebSocket；[聊天室状态](../../internal/goweb/gowebsockets/chat/chatroom.go)通过 channel 驱动加入、离开、消息广播，用户表使用 `sync.Map`。代码里没有使用双向链表，原笔记“为什么需要双向链表”的问题与当前实现无关。

这个包只有路由注册和客户端函数，没有调用 `web.Run` 的项目启动入口。因此它**不是可直接运行的 Beego 聊天室示例**。编译可核对 API 仍与依赖版本兼容，但不能证明连接、广播和退出行为正确。Beego 控制器路由的用法参见 [Beego v2.1 路由文档](https://beegodoc.com/en-US/v2.1.x/web/router/ctrl_style/)。

## 当前代码的运行边界

- [客户端](../../internal/goweb/gowebsockets/chat/chatclient.go)固定连接 `127.0.0.1:8090` 并从标准输入读取；需先有对应服务器入口。
- 服务器在包初始化时启动永久运行的聊天室 goroutine；目前没有显式生命周期和退出控制。`Join` 在升级连接后才检查重名；断开后仍尝试通过 HTTP 输出错误。广播写入失败未处理；用户名检查与注册并非同一步完成。这些行为需要在后续聊天室任务中修正与测试。
- WebSocket 升级器没有显式配置来源策略；当前 gorilla/websocket 版本会在请求带 `Origin` 时默认要求其主机与请求主机相同。代码没有身份认证，也没有服务器启动与部署配置，不能把此演示作为对外服务部署说明。见 [Upgrader 文档](https://pkg.go.dev/github.com/gorilla/websocket#Upgrader)。
- 仓库的可运行 Web 路由和受控代理入口见 [Gin 路由与代理](go_gin.md)。Beego 聊天室与其是独立示例。

## 验证

在仓库根目录执行 `go build ./internal/goweb/gowebsockets/chat`；该命令只做编译。聊天室的入口、连接流程、广播和安全边界仍列在[待办](../../TODO.md)中，未宣称已验收。
