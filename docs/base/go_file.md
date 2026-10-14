# 文件读写与路径

> 状态：已整理并验证（普通文件路径）；验证环境：Go 1.25.5、Linux。代码入口：[文件操作实现](../../internal/golib/osdemo/file_operation.go)、[临时目录测试](../../test/golib/filedemo_test.go)、[复制测试](../../test/golib/filecopy_test.go)。

## 打开与读取

`os.Open(path)` 以只读方式打开文件；需要指定模式时用 `os.OpenFile(path, flags, perm)`。创建文件时 `perm` 还会受到进程 umask 影响；打开现有文件时它不是修改权限的手段。持有 `*os.File` 后要关闭，通常在成功打开后立即 `defer f.Close()`。

小文件可用 `os.ReadFile` 一次读入；大文件或流式处理用 `io.Reader`、`bufio` 或 `io.Copy`，避免一次占用与文件大小相同的内存。`Read` 返回的 `n` 才是本次有效字节数，即使用了更大的缓冲区，也只能访问 `buf[:n]`。仓库的 `readFileData` 已按此规则修正，并向调用方传播读取错误；[短文件和空文件测试](../../internal/golib/osdemo/file_operation_test.go)检查了有效字节范围。

## 写入与复制

`os.WriteFile(path, data, perm)` 会创建文件，若文件已存在则截断内容；需要追加时用 `os.O_APPEND|os.O_CREATE|os.O_WRONLY`。`os.O_TRUNC` 与追加语义相冲突，不要在同一个示例中混用。打开后的文件可用 `Write` 或 `io.WriteString` 写入，写完要检查错误并关闭文件。

仓库有三种复制函数：

| 函数 | 行为 | 入口 |
| --- | --- | --- |
| `SimpleCopyFile` | 读入全部字节，再覆盖目标 | [实现](../../internal/golib/osdemo/file_operation.go) |
| `AdvancedCopyFile` | 用 `ReadFrom` 复制并追加到目标 | [追加断言](../../test/golib/filecopy_test.go) |
| `IOCopyFile` | 用 `io.Copy` 复制并覆盖目标 | [复制断言](../../test/golib/filecopy_test.go) |

`io.Copy` 优先利用源或目标支持的优化接口，具体调用路径由实现决定；业务代码应依赖“复制字节并返回数量/错误”的公开语义。创建多个文件时可用 `filepath.Join` 组合路径，不要手写平台分隔符。`os.Remove` 删除单个文件或空目录；`os.RemoveAll` 递归删除，务必限定目标路径。

## 验证与边界

```bash
go test ./test/golib -run '^(TestFileOperations|TestCopyFileInTempDir|TestAdvancedCopyAppends)$'
go test ./internal/golib/osdemo -run '^TestReadFileDataUsesOnlyReadBytes$'
```

测试全部使用 `t.TempDir()`，无需固定本机路径。它们验证普通文件的读写、覆盖、追加和删除。`SimpleOpenFile`/`AdvancedOpenFile` 只演示一次 `Read`，不能当作读取整个大文件的函数；需要完整内容时使用 `os.ReadFile` 或循环读取。符号链接、特殊文件、文件权限及磁盘故障未逐项模拟。更多 API 语义见[`os` 包文档](https://pkg.go.dev/os)和[`io` 包文档](https://pkg.go.dev/io)。
