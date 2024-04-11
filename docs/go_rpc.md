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
