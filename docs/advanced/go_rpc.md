# gRPC：本机 Hello World 示例

> 状态：已整理并验证；验证环境：Go 1.25，仓库使用 `google.golang.org/grpc` v1.41.0。关联代码：[`internal/gorpc/helloworld`](../../internal/gorpc/helloworld/)。

## 场景与结论

这个示例展示一次最小的 gRPC 一元调用：[`hello.proto`](../../internal/gorpc/helloworld/protobuf/hello.proto) 定义请求、响应和 `SayHello.Hello` 方法；[服务端](../../internal/gorpc/helloworld/server/main.go) 注册实现；[客户端](../../internal/gorpc/helloworld/client/main.go) 建立连接并发送请求。

客户端用同一个 5 秒截止时间约束连接和 RPC，连接失败与调用失败分别报告错误，并在退出时关闭连接。服务端仅监听 `127.0.0.1:8090`；客户端使用无 TLS 的连接凭据，因此此示例只用于本机演示。

## 运行与验证

在仓库根目录打开两个终端。先启动服务端：

```bash
go run ./internal/gorpc/helloworld/server
```

再运行客户端：

```bash
go run ./internal/gorpc/helloworld/client
```

客户端应输出 `hello, world zhangsan`；服务端会记录收到的请求。运行完毕后在服务端终端按 Ctrl+C。若服务端未启动，客户端应在约 5 秒内报告连接超时，而不会一直等待。默认 `go build ./...` 和 `go test ./...` 无需启动服务端，也无需安装 `protoc`：生成的 Go 文件已保存在仓库中。

## 范围与后续

[`internal/gorpc/chat`](../../internal/gorpc/chat/) 目前只有协议定义和生成代码；客户端、服务端仍是占位包，详见[目录说明](../../internal/gorpc/chat/README.md)。这个示例尚不能运行。

Hello World 演示不涵盖 TLS、认证、流式调用或部署配置。需要从 `.proto` 重新生成代码时，再安装与仓库配置匹配的 Protocol Buffers 工具链；运行现有示例不需要重新生成代码。
