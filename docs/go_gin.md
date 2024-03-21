## Gin

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
