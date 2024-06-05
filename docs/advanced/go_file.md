## 文件操作

> Go中文件操作的基础类型是[]byte，字节切片；

### 文件打开

> Go内置的os包，主要通过Open和OpenFile方法对文件进行操作；(Open方法实际上也是调用OpenFile，OpenFile能做到更加精细的控制，以不同的模式打开文件)


Open打开文件

```go
```
