## 基础

grpc协议基于http2协议，相对于http1协议的优点：
- 传输的是二进制数据，性能好、效率高；
- 序列化和反序列化直接对应程序中的类，不需要解析后进行映射；
- 支持多种语言；

grpc使用：
- 定义protobuf文件，具体包含接口和数据结构;
- 使用protoc工具编译protobuf文件，生成对应的客户端和服务端的代码；
- 在服务端实现对应业务逻辑代码；
- 在客户端建立gRPC连接，自动生成代码调用函数；

## 安装
1. 安装probuf编译器
```shell
 apt install -y protobuf-compiler
```

2. 安装代码生成器，将对应proto文件生成对应语言的代码文件
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
3. 修改环境变量(.zshrc)
```
export PATH=$PATH:/root/go/bin
```

grpc 参考文档：
> https://golang.halfiisland.com/community/mirco/gprc.html


