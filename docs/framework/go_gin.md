# Gin：本机路由与代理示例

> 状态：基础路由、HTTP 代理及隔离的旧表单/上传/外部 HTTP 处理器已做本机测试；SSH、终端和设备网络演示未连接真实设备。验证环境：Go 1.25.5、Gin v1.9.1。代码入口：[基础路由](../../internal/goweb/gogin/app/second_demo.go)、[代理路由](../../internal/goweb/gogin/proxy/routers/router.go)、[命令入口](../../cmd/gin.go)。

## 场景与结论

基础示例展示 Gin 的路由注册、分组与 HTML/JSON 响应。代理示例只转发到启动时明确允许的 HTTP(S) 源站；未设置 `--proxy-target` 时所有转发请求都返回 403。两项命令默认只监听 `127.0.0.1`。

## 基础 Gin 路由

从**仓库根目录**运行，模板路径以该目录为基准：

```bash
go run . gin
curl -i http://127.0.0.1:8999/index
curl -i http://127.0.0.1:8999/json-group/print-json
```

预期 `/index` 返回 200 和 `Main WebSite` 页面，`/json-group/print-json` 返回 200 和 JSON。`go run . gin --port 9000` 可换端口。基础路由实现见 [`app.NewRouter`](../../internal/goweb/gogin/app/second_demo.go)，本地无网络服务的断言见[路由测试](../../internal/goweb/gogin/app/second_demo_test.go)。

| 路由 | 用途 | 前提 |
| --- | --- | --- |
| `GET /index` | HTML 模板 | 仓库根目录下有 `web/templates/index.html` |
| `GET /json-group/print-json` 等 | JSON 响应格式演示 | 无外部服务 |
| `POST /json-group/json` | JSON 请求绑定演示 | 已注册；无效 JSON 返回 400，成功返回 200 |

`formGroupInit`、`othersGroupInit` 和单独的文件上传服务仍未接入 `go run . gin`。其处理器已在隔离路由中复核：登录表单修正 `password` 标签，缺字段返回 400、凭据不匹配返回 401；单文件和多文件上传只返回文件元信息或数量，**不会写入工作目录**，请求体限制为 8 MiB；[独立上传处理器](../../internal/goweb/gogin/app/file_upload_demo.go)同样只解析和验证文件，不写磁盘，默认只监听本机。演示账号为固定样例值，隔离路由不应作为真实认证服务。[处理器测试](../../internal/goweb/gogin/app/legacy_handlers_test.go)覆盖成功、缺字段、超限和错误响应。

`othersGroupInit` 的图片代理接收调用方配置的图片 URL 与 HTTP 客户端；默认客户端超时为 5 秒，请求沿用入站上下文。非 200 上游响应返回 503，成功和失败时都关闭上游响应体。它没有接入主入口，也不允许由请求参数指定任意目标；测试通过内存传输替身验证响应和资源关闭，未访问公网。若要在其他程序中注册它，需自行限定可信目标与访问权限。

## HTTP 代理

先准备一个自有的本地 HTTP 服务，例如另开终端运行 `python3 -m http.server 9001 --bind 127.0.0.1`，再从仓库根目录启动：

```bash
go run . proxy --proxy-target http://127.0.0.1:9001
curl -i http://127.0.0.1:8090/healthz
curl -i 'http://127.0.0.1:8090/netac/base/proxy?proxy_pass=http%3A%2F%2F127.0.0.1%3A9001%2F'
```

`/healthz` 返回 `{"status":"ok"}`；代理请求返回源站页面。`--proxy-target` 只接受完整的 HTTP(S) **源站**（协议、主机、可选端口），可重复传入；请求中的 `proxy_pass` 必须是该源站的完整 URL。格式错误返回 400，未允许的源站返回 403，上游请求失败返回 502。代码使用标准 TLS 校验，不跳过证书验证；自签名证书需先加入系统信任链，或在测试中预期 502。转发行为、目标限制和 TLS 校验见[代理测试](../../internal/goweb/gogin/proxy/routers/router_test.go)。

代理入口见 [`godemo proxy`](../../cmd/proxy.go)，路由见[路由表](../../internal/goweb/gogin/proxy/routers/router.go)。默认还注册 `GET /netac/base/view_stat`（返回模拟数据）、`GET /netac/base/ping`、`GET /netac/base/traceroute` 和 `POST /netac/base/upload`。上传端点只解析最多 8 MiB 的 multipart 请求，不保存或记录字段内容：格式错误返回 400，超限返回 413，合法请求返回成功 JSON；见[处理器测试](../../internal/goweb/gogin/proxy/handlers/file_upload_test.go)。ping/traceroute 仍依赖操作系统网络工具，未做设备级运行验收。当前服务没有客户端身份认证；即使显式改用 `--address 0.0.0.0`，也只应在隔离的可信网络中运行，并自行加访问控制。源站允许列表代表信任该源站及其可解析的地址，不能把用户任意提供的地址加入列表。

## SSH 与终端的状态

`POST /netac/base/ssh` 和 WebSocket `GET /netac/base/terminal` 默认不注册。仅在明确指定 `--enable-ssh` 且 `GO_DEMO_SSH_KNOWN_HOSTS` 指向可读取的 OpenSSH `known_hosts` 文件时才注册；缺少或损坏该文件会导致启动失败。应先从可信渠道核对主机密钥，再把已核对的条目写入该文件。SSH 连接使用该文件验证主机密钥，终端升级只接受同源浏览器请求。此功能仍需 SSH 设备、凭据和人工交互，未验证真实连接；它也没有客户端身份认证，不适合暴露到公共网络。

## Gin v1.9.1 内部调用路径

当前依赖版本的 `RouterGroup` 在注册时合并组中间件与处理器，再按 HTTP 方法和路径存入路由树；`Engine.ServeHTTP` 从池中取得 `Context`、重置请求状态、查找路由并执行处理链，完成后归还池。中间件使用 `c.Next()` 推进处理链，因此同一中间件可在调用前后处理请求与响应。此处只解释 [Gin v1.9.1 源码](https://github.com/gin-gonic/gin/blob/v1.9.1/gin.go)中的调用路径，不把池或树的布局当成稳定 API。

仓库的绑定分支采用 `ShouldBind`/`ShouldBindJSON`，由处理器自己返回 400 并立即结束；`Bind` 系列在失败时会写入 400 响应，再写其他状态可能引发重复写响应的问题。上传时 `MaxMultipartMemory` 只是 multipart 解析的内存阈值，不能代替请求体大小限制，所以处理器另用 `http.MaxBytesReader`。参见 [Gin 绑定文档](https://gin-gonic.com/en/docs/binding/)与[上传文档](https://gin-gonic.com/en/docs/routing/upload-file/)。

## 验证与边界

```bash
go build ./...
go test -timeout 20s ./internal/goweb/gogin/...
go test -tags=integration -run '^$' ./test/gogin
```

第二条测试基础路由、隔离旧处理器、代理允许列表、HTTP 转发、上传错误状态和 TLS 证书校验；代理测试会在本机监听临时端口。第三条只编译手动演示，**不会**运行会持续监听的 `TestGinStart`。该手动测试使用测试目录到模板的相对路径。真实 SSH/终端会话需要专用设备和已核对的主机密钥，属于可选手动验收；本阶段未执行，不据此宣称设备功能已验证。
