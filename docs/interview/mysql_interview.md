# MySQL 问答

状态：按 MySQL 8.4 官方文档复核概念；本仓库仅完成 sqlx/GORM 静态检查，未用真实 MySQL 验证查询计划或事务行为。实践入口：[数据库静态验收](../advanced/go_database.md)、[GORM 笔记](../advanced/go_gorm.md)。

## 索引与查询计划

**InnoDB 索引一定是 B+ 树吗？** 常规 InnoDB 聚簇和二级索引是 B-tree 类结构；全文和空间索引另有实现，不能把所有索引概括为一种。聚簇索引保存行数据，二级索引记录包含主键列，可通过主键定位行。见 [InnoDB 索引文档](https://dev.mysql.com/doc/refman/8.4/en/innodb-index-types.html)。

**联合索引的“最左匹配”怎么用？** 先看索引列顺序、筛选条件、排序及覆盖需求；范围条件、函数/表达式和优化器选择都可能改变可用部分。不要只凭口号推断，针对具体 SQL 使用 `EXPLAIN`。

**哪些条件“必然不走索引”？** 旧笔记把 `!=`、`IS NULL`、低选择性列和前导通配 `LIKE` 一律列为不命中，结论过强。优化器会按统计信息、代价和索引类型选路径。前导通配通常难以利用普通 B-tree 前缀查找；实际计划仍以目标数据上的 `EXPLAIN`/`EXPLAIN ANALYZE` 为准。

**`EXPLAIN` 看什么？** 先看访问方式、`possible_keys`、`key`、估计行数和过滤比例；`key` 是选择的索引，`possible_keys` 是候选。估计值不等于实际行数。MySQL 8.4 的默认格式可配置，不宜把旧版表格列当作唯一输出格式。见 [EXPLAIN 文档](https://dev.mysql.com/doc/refman/8.4/en/explain.html)。

## 事务与存储

**ACID 是什么？** 原子性、一致性、隔离性、持久性。一致性指数据库约束及应用不变量在事务前后得到维护，不能由隔离级别单独保证。持久性还取决于存储及刷盘配置。

**四种隔离级别和幻读？** SQL 常见级别为读未提交、读已提交、可重复读、串行化。InnoDB 默认可重复读；一致性读、锁定读及范围锁的行为不同，不能简单说“可重复读一定产生/一定消除幻读”。需要结合 SQL 语句、事务时间线和 [InnoDB 隔离级别文档](https://dev.mysql.com/doc/refman/8.4/en/innodb-transaction-isolation-levels.html)分析。仓库的事务示例见[数据库说明](../advanced/go_database.md)，但真实回滚未验收。

**MVCC、redo、undo、binlog 各负责什么？** MVCC 借助版本与读视图实现一致性读；undo 保存事务回滚及旧版本所需信息；redo 是 InnoDB 恢复已提交修改的重要日志；binlog 是 MySQL 服务器层的变更日志，常用于复制与时间点恢复。它们不是同一种“事务日志”。此处是概念入口，具体故障恢复需结合配置和官方手册验证。

**`utf8` 与 `utf8mb4`？** MySQL 的 `utf8` 是已弃用的 `utf8mb3` 别名，最多编码 3 字节字符；`utf8mb4` 支持 4 字节字符。新表应明确选择字符集和排序规则。见 [MySQL utf8mb3 文档](https://dev.mysql.com/doc/refman/8.4/en/charset-unicode-utf8mb3.html)。

## 架构问题的边界

主从复制、MHA/MMM、共享存储等是不同年代和场景的高可用方案，不能作为当前项目推荐架构。本仓库没有部署、故障切换或恢复演练；选择方案前需明确 MySQL 版本、RPO/RTO、读写模型和运维能力。

验证命令：`go test ./internal/gostorage/... ./test/gostorage` 只验证无需数据库的静态与模型行为；执行真实 `EXPLAIN` 和事务测试需按[数据库准备说明](../advanced/go_database.md)连接专用测试库。
