
初始化：
New() -> init Engine -> init RouterGroup -> 保存当前Engine到RouterGroup中的engine地址

RouterGroup注册路由
Group() -> 
返回一个新的RouterGroup（保存Handlers、basePath、engine) 
-> 注册路由 
-> 计算路由绝对路径
-> 保存handlers（deep copy）
-> 保存到路由树?
-> 


