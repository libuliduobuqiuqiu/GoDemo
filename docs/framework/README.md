# 框架主题

这里记录框架用法与项目结构。Gin 的基础路由、HTTP 代理及旧处理器已做本机验证；Beego 已核对依赖与聊天室代码，但聊天室尚无启动入口和端到端验收。

| 框架 | 文章 | 状态 | 相关代码 |
| --- | --- | --- | --- |
| Gin | [Gin 路由、旧处理器与代理](go_gin.md) | 路由、隔离旧处理器和代理已验证；SSH/终端未连接设备 | [`internal/goweb/gogin`](../../internal/goweb/gogin/) |
| Beego | [Beego 与 WebSocket 聊天室](go_beego.md) | 路由及依赖已静态复核；聊天室未运行验收 | [`internal/goweb/gowebsockets/chat`](../../internal/goweb/gowebsockets/chat/) |

[返回知识库首页](../README.md)
