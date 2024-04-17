## 基础

TCP和UDP的区别?
- TCP面向连接，TCP需要建立连接才能传输、UDP无连接
- TCP提供可靠的服务，UDP尽可能交付，不保证可靠，可能出现丢包
- TCP面向字节流，UDP面向报文
- 每一条TCP连接都是点对点，UDP支持一对一、一对多、多对多的交互通信
- TCP的逻辑信道是全双工的可靠信道，UDP则是不可靠信道

TCP的三次握手和四次挥手？
https协议和http协议有什么区别？
grpc的优缺点？
ASCII、UNICODE、UTF-8区别？
- Ascii: 定义了英文字符和二进制之间的对应关系
- Unicode: 字符集，包含了所有符号的编码；
- UTF-8：编码规则，unicode只规定了符号的二进制代码，但是没有规定二进制代码是如何存储的，utf-8可以理解为unicode的实现之一

Cookie和Session区别？
- Cookie和Session都是为了解决HTTP无状态，用来保存客户端状态信息的机制；
- Cookie是保存到客户端，Session是保存到服务端

Client如何实现长连接？
- tcp中的心跳包机制（HeartBeat），tcp的选项：SO_KEEPALIVE
- http的keepalive能够解决每次请求都需要重新建立连接释放连接，但是需要正确设置超时时间，保证tcp连接导致系统资源不会被无效占用；

http1,http2,grpc之间的区别？
- http协议是用于应用层数据传输；
- 影响http请求主要的问题：网络和延迟；（延迟：浏览器阻塞、DNS查询、建立连接）
- http2相对于http1解决的就是性能问题，主要的方式：二进制格式、多路复用、Header压缩、服务端推送；
- grpc协议是基于http2

TCP的拆包和粘包？
- 粘包：当通讯一端一次性发送多条数据时，TCP会将多条数据打包成一条tcp报文发送出去
- 拆包：当通讯一端一次性发送的数据大于TCP报文一次传输的最大长度，就会将发送数据进行拆分成多个TCP长度的tcp报文发送

TIME_WAIT作用？
- 主动关闭的Socket段就会进入TIME_WAIT状态，进入TIME_WAIT状态，并且持续2MSL时间长度

网络性能指标？
- 带宽
- 延迟
- 吞吐率
- PPS(Packe Per Second)

