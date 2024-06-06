## 文件操作

> Go中文件操作的基础类型是[]byte，字节切片；

### 文件打开

> Go内置的os包，主要通过Open和OpenFile方法对文件进行操作；(Open方法实际上也是调用OpenFile，OpenFile能做到更加精细的控制，以不同的模式打开文件)

#### 普通Open打开文件
```go
file, err := os.Open(filePath)
	if err != nil {
		return err
	}

```
> Open默认只读模式打开文件

#### OpenFile打开文件
```go
file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
```
> OpenFile函数可以控制以什么方式打开文件，上述代码是只写模式，所以无法读取到文件里面的内容；

具体控制的细节
```
const (
   // 只读，只写，读写 三种必须指定一个
   O_RDONLY int = syscall.O_RDONLY // 以只读的模式打开文件
   O_WRONLY int = syscall.O_WRONLY // 以只写的模式打开文件
   O_RDWR   int = syscall.O_RDWR   // 以读写的模式打开文件
   // 剩余的值用于控制行为
   O_APPEND int = syscall.O_APPEND // 当写入文件时，将数据添加到文件末尾
   O_CREATE int = syscall.O_CREAT  // 如果文件不存在则创建文件
   O_EXCL   int = syscall.O_EXCL   // 与O_CREATE一起使用, 文件必须不存在
   O_SYNC   int = syscall.O_SYNC   // 以同步IO的方式打开文件
   O_TRUNC  int = syscall.O_TRUNC  // 当打开的时候截断可写的文件（清空文件）
)
```

### 文件读取

> Go内置的os提供了ReadFile直接读取文件内容；也能直接调用io.ReadAll方法直接读取打开的文件；
同样也可以参考ReadFile的源码，动态扩容缓存，调用本身os.File类型提供的Read方法，将文件内容读取到切片；

#### 直接调用os.ReadFile方法读取文件
```go
data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

```
> 传入对应文件路径，返回字节切片

#### io.ReadAll读取打开的文件

```go
f, err := os.Open(filePath)
	if err != nil {
		return err
	}

  defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "读取文件数据失败")
	}
```

#### 打开文件后，通过os.File的Read方法读取文件内容

代码参考的是os.ReadFile源码
```go
func AdvancedReadFile() (data []byte, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer f.Close()

	data = make([]byte, 0, 512)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}

}
```
> 打开文件，创建一个缓存字节切片，通过无限循环不断读取文件数据，检查切片是否需要扩容，扩容后继续读取后续文件数据，最后读到文件关闭后，返回数据；


### 文件写入

> 可以直接通过OpenFile打开文件，模式需要设置只写模式或者读写模式，否则无法成功写入文件；os还有在前面基础上提供便捷函数,os.WriteFile和io.WriteString方式。

#### 普通打开文件，通过os.File的Write和WriteString方法写入

```go
  file, err := os.OpenFile(newFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	fmt.Println("文件打开成功： ", file.Name())
	for i := 0; i < 5; i++ {
		offset, err := file.WriteString("Hello,world\n")
		if err != nil {
			return errors.Wrap(err, "文件写入失败")
		}
		fmt.Println(offset)
	}
```
#### os.WriteFile直接将将数据写入文件

```go
	err := os.WriteFile(path, []byte("hello,world\nhello,world\n"), 0666)
	if err != nil {
		return err
	}
```
> WriteFile函数只是封装了打开文件的步骤，直接传入字符串即可

#### io.WriteString写入数据

```go
  file, err := os.OpenFile(newFilePath, os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	fmt.Println("文件打开成功： ", file.Name())
	for i := 0; i < 6; i++ {
		n, err := io.WriteString(file, "Hello,world\n")
		if err != nil {
			return errors.Wrap(err, "文件写入失败")
		}
		fmt.Println(n)
	}
	return nil

```
> io.WriteString 不仅可以写文件，只要实现了io.Writer接口的，都可以写入（net.Conn、os.Stdout、bytes.Buffer)，更加灵活；


