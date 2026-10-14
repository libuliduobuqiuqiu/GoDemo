# Redis 问答

状态：已按 Redis Open Source 官方文档校正概念；**仓库没有 Redis 客户端、服务端或自动测试示例**，本页是知识笔记与官方实践入口，不宣称已验证 Redis 运行行为。示例命令需在独立测试实例上执行，并核对所用 Redis 版本。

## 数据结构与持久化

**常用数据类型？** String、Hash、List、Set、Sorted Set 各提供不同命令语义；不要把 List 固定描述为双向链表，内部编码可变。[Redis 数据类型文档](https://redis.io/docs/latest/develop/data-types/)列出可用类型与命令。

**RDB 与 AOF？** RDB 是时间点快照，AOF 记录导致数据变化的操作；可单独或组合配置。AOF 每秒 `fsync` 是一种配置，不保证故障时“最多只丢一秒”，还受写入、操作系统和存储故障影响。复制是另一套机制，增量复制流不能等同于 AOF 文件。见 [持久化](https://redis.io/docs/latest/management/persistence/)与[复制](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/)文档。

**初次与断线后复制？** 首次或无法部分重同步时通常全量同步；断线重连可尝试根据复制 backlog 继续部分同步。具体流程依版本和配置变化，不能概括为“全量 RDB 后增量 AOF”。

## 内存、过期与缓存

**Redis 为什么常用于低延迟场景？** 内存访问、紧凑数据结构和事件驱动 I/O 都有帮助；不能简单说“Redis 只有一个线程”。后台持久化、I/O 和其他任务会用到额外线程或进程，具体性能需测量工作负载。

**过期与内存淘汰区别？** TTL 到期删除与达到 `maxmemory` 时按 `maxmemory-policy` 淘汰是两个机制。缓存场景可选近似 LRU、LFU 等策略；没有“Redis 默认一定按 LRU 淘汰”的规则。见 [Redis 淘汰策略](https://redis.io/docs/latest/develop/reference/eviction/)。

**缓存穿透、雪崩、并发竞争如何处理？** 穿透先校验请求、限制来源，必要时短期缓存空结果；大量键同时到期可分散 TTL 并保护回源容量。竞争同一键时先定义一致性要求：单键条件更新可用 `SET ... NX/XX` 等原子命令；跨键操作或分布式锁需要更严格的失败与过期处理。见 [SET 命令](https://redis.io/docs/latest/commands/set/)；仓库没有相应 Redis 测试，不能把这些策略标为已实施。

## 部署与有序集合

**主从、Sentinel、Cluster？** 复制提供副本，Sentinel 负责监测和故障转移，Cluster 还做数据分片。它们解决的容量与可用性问题不同，选型需考虑故障切换与数据丢失边界。入口：[Redis 复制](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/)、[Sentinel](https://redis.io/docs/latest/operate/oss_and_stack/management/sentinel/)、[Cluster](https://redis.io/docs/latest/operate/oss_and_stack/management/scaling/)。

**Sorted Set 一定是跳表吗？** 不应把内部编码当作 API 保证。对外只依赖按 score 排序及命令复杂度；具体命令查 [Sorted Set 文档](https://redis.io/docs/latest/develop/data-types/sorted-sets/)。

后续若要把本页升级为实践专题，应先添加受控的 Redis 测试环境、最小客户端代码和自动清理的数据，再逐项验证 TTL、条件写入和故障恢复。
