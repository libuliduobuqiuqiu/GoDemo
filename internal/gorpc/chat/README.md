# gRPC chat：尚未实现

这里已保存 [`message.proto`](protobuf/message.proto) 和生成的 Go 协议代码，但 `client/` 与 `server/` 只有占位包，没有可运行的聊天服务或客户端。`go build ./...` 可以编译这些包，不能用 `go run` 启动它们。

可运行的最小 RPC 示例见 [Hello World](../helloworld/)；启动和验证步骤见 [RPC 笔记](../../../docs/advanced/go_rpc.md)。
