
## 索引

索引使用b+树算法？

索引类型：
- 唯一索引 UNIQUE INDEX
- 普通索引 NORMAL INDEX
- 主键索引 PRIMARY INDEX
- 全文索引 FULLTEXT INDEX
- 联合索引

联合索引的最左匹配原则?
聚簇索引和非聚簇索引的区别？
explain分析是否命中索引？
不会命中索引：
1. 索引设计不合理：bool字段做索引，sql选择器不命中索引；
2. where 不等于、对null的判断；
3. 模糊查询%like，需要回表的查询结果集过大；

## 事务
ACID：
- Atomicity: 原子性，事务是一个原子操作单元，对数据的修改要么全部执行，要么全部不执行；
- Consistency：一致性：事务开始前和结束后，数据库的完整性约束没有破坏；
- Isolation：隔离性，允许多个事务同个数据进行读写和修改的能力，但是隔离性可以防止多个事务之间交叉执行导致数据的不一致
- Duration：持久性，事务操作之后，对数据的操作是永久的

事务隔离性区分的四个级别：
- 读未提交：脏读
- 读已提交：避免脏堵、出现不可重复度
- 可重复读：避免不可重复读、出现幻读（可通过MVCC版本控制解决幻读）
- 串行

## 架构

高可用方案：
- 主从复制方案
- MMM/MHA高可用方案
- Heartbeat/SAN高可用方案
- Heartbeat/DRBD高可用方案

## 基础

mysql的utf8和utf8mb4有什么区别？
utf8最大字符长度为3字节，遇到4字节的宽字符(Emoji)就会插入异常

乐观锁和悲观锁
MVCC数据库并发控制
MVCC、Redolog、Undolog、Binlog有什么区别？

## 性能

explain命令

+----+-------------+-----------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
| id | select_type | table     | partitions | type  | possible_keys | key     | key_len | ref   | rows | filtered | Extra |
+----+-------------+-----------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
|  1 | SIMPLE      | syn_group | NULL       | const | PRIMARY       | PRIMARY | 192     | const |    1 |   100.00 | NULL  |
+----+-------------+-----------+------------+-------+---------------+---------+---------+-------+------+----------+-------+

select_type(查询的类型)
- SIMPLE （简单使用SELECT，不使用UNION或子查询）
- PRIMARY（查询中包含任何复杂的子部份，最外层的select被成为PRIMARY）
- UNION（UNION第二个或者后面Select）
- DEPENDENT UNION
- UNION RESULT（union的结果）
- SUBQUERY（子查询中的第一个SELECT）
- DEPENDENT SUBQUERY
- DERIVED（派生表的SELECT FROM 子句查询）
- UNCACHEABLE SUBQUERY

type(访问类型,从上到下，性能从差到好)
- ALL（全表查询）
- index（只遍历索引树）
- range（之检索给定范围的行）
- ref（表的连接匹配条件）
- eq_ref（多表连接使用唯一索引）
- const、system（查询转化为常量，例如查询的主键在where列表中，system是查询的表中仅有一行的情况）
- NULL

possible_keys: 可能使用的索引，并不是真正被查询使用；
key: 实际使用的索引
key_len：使用的索引的最大长度
ref：上述表的连接匹配条件
rows: mysql根据表统计信息及索引选用情况，估算找到所需记录需要读取的行数；


