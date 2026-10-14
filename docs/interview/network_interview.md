# 网络问答

状态：已核对基本协议语义；仓库只有本机 RPC、HTTP/WebSocket 演示，未做抓包或公网故障实验。实践入口：[gRPC 示例](../advanced/go_rpc.md)、[Gin 路由与代理](../framework/go_gin.md)。

## TCP、UDP 与连接

**TCP 和 UDP 的主要差别？** TCP 提供有序可靠的字节流和连接状态；UDP 保留数据报边界，不承诺交付、顺序或去重。应用层可在 UDP 之上另行实现可靠性。TCP 的“粘包/拆包”是应用一次 `Write` 与接收端一次 `Read` 没有消息边界对应关系；用长度前缀或分隔协议划分消息，不按 TCP 分段猜测消息长度。见 [TCP 规范 RFC 9293](https://www.rfc-editor.org/rfc/rfc9293)。

**三次握手、关闭与 TIME_WAIT？** 建连通过 SYN、SYN-ACK、ACK 协商初始序列空间；关闭通常涉及双方各自关闭发送方向，但实际报文可合并，不能保证总是“四个包”。TIME_WAIT 用于处理旧报文及最后 ACK 的重传；持续时间和端口复用由协议及操作系统实现决定，不能把 2MSL 写成所有机器固定时长。

**长连接与 keep-alive？** TCP keepalive 探测与 HTTP 持久连接是不同层面的机制。Go 客户端应复用 `http.Client`/`Transport`，设置合理的超时和空闲连接参数；不能仅靠心跳解决服务端、代理和网络中的全部超时。见 [net/http 文档](https://pkg.go.dev/net/http)。

## HTTP、TLS 与 RPC

**HTTP、HTTPS、HTTP/2 和 gRPC 的关系？** HTTPS 是 HTTP 经 TLS 保护的常见部署形式，不是与 HTTP 独立的应用协议。HTTP/2 用二进制帧和流多路复用；服务端推送是可选机制，不能作为当前浏览器场景必备优势。gRPC 定义服务与消息语义，常使用 HTTP/2 承载。见 [RFC 9113](https://www.rfc-editor.org/rfc/rfc9113)、[gRPC 核心概念](https://grpc.io/docs/what-is-grpc/core-concepts/)及仓库的[本机 gRPC 示例](../advanced/go_rpc.md)。

**Cookie 与 Session？** Cookie 是客户端携带的 HTTP 状态机制；Session 是服务端管理会话的一种应用设计，通常用 Cookie 保存会话标识，也可用其他传递方式。Cookie 属性、会话存储和过期机制决定安全边界，不能笼统说两者“都是保存状态”。

**ASCII、Unicode、UTF-8？** ASCII 是早期字符编码；Unicode 定义字符及码点，UTF-8 是 Unicode 码点的一种可变长度编码。字符串和字节序列的区别见[Go 语言总览](../base/go_manual.md)。

## 性能指标

带宽是链路可提供的数据传输能力，吞吐量是实际完成量，延迟是一次操作的耗时，PPS 是每秒包数。诊断网络性能时还要看丢包、重传、连接数及测量范围；单一指标无法解释应用响应时间。

验证：`go test ./internal/goweb/gogin/app ./internal/goweb/gogin/proxy/routers` 与 `go build ./internal/gorpc/helloworld/...` 只验证仓库示例，不构成协议抓包证明。
