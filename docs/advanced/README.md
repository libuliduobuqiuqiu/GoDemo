# 进阶主题

这里收纳测试方法、调试工具、数据库和网络服务实践。涉及数据库或设备的示例先看[测试运行说明](../testing.md)。

| 主题 | 文章 | 状态 | 相关代码 |
| --- | --- | --- | --- |
| 测试 | [Go 测试笔记](go_testing.md) | 已整理并验证 | [`test/manual`](../../test/manual/)、[`test/golib`](../../test/golib/) |
| 反射 | [反射笔记](go_reflect.md) | 原始笔记 | [`internal/gomanual/reflectdemo`](../../internal/gomanual/reflectdemo/) |
| 数据库实践 | [sqlx 与 GORM 往返验证](go_database.md) | 已整理，待真实数据库验证 | [`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go) |
| GORM | [GORM 笔记](go_gorm.md) | 核心 API 已复核，待真实数据库验证 | [`internal/gostorage/gormdemo`](../../internal/gostorage/gormdemo/)、[`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go) |
| RPC | [gRPC Hello World](go_rpc.md) | 已整理并验证（chat 未实现） | [`Hello World`](../../internal/gorpc/helloworld/)、[`chat 状态`](../../internal/gorpc/chat/README.md) |
| 性能分析 | [pprof 笔记](go_pprof.md) | 原始笔记 | [`internal/gotool/pprofdemo`](../../internal/gotool/pprofdemo/) |

[返回知识库首页](../README.md)
