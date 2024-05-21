## Gin
> 有关Gin源码阅读的笔记

初始化：
New() -> init Engine -> init RouterGroup -> 保存当前Engine到RouterGroup中的engine地址

注册路由:
-> Group() 返回一个新的RouterGroup（保存Handlers、basePath、engine) 
-> 注册路由 
-> 计算路由绝对路径
-> 保存handlers（deep copy）
-> 保存到路由树?
-> addRoute?(node)
-> 统计路径参数（countParams）统计路径节点（countSections）
-> 注册路由结束

启动：
-> 解析监听地址，启动http服务监听地址(net.Listen)
-> 服务接收Listen上传入的连接
-> 启用service goroutine，处理requests，并且回调Handler进行回复
-> 调用Handler的ServeHTTP方法处理requests请求 
-> Gin引擎handleHTTPRequest处理请求，搜索路径树，获取到对应请求路径绑定的处理方法
-> 调用路径绑定的Handler的方法，返回响应

日志组件打印(中间件)：
-> 读取设置日志格式、日志输出Writer;
-> 生成一个控制器函数，函数中设置起始时间，继续Next调用Context绑定的其余Handlers；
-> 其余Handlers结束返回，根据结束时间计算耗时，根据日志格式，打印请求日志；
-> 最后将这个控制器函数返回，保存到RouterGroup上，后续基于这个RouterGroup上创建新的Group时都会复制这个Handler；
-> 后续当请求访问匹配到对应的路由时，就会执行路由上绑定的HandlerFunc；
