# MySQL：sqlx 与 GORM 往返验证

> 状态：已整理，待真实数据库验证；验证环境：Go 1.25、MySQL 8，仓库使用 `github.com/jmoiron/sqlx` v1.3.5 与 `gorm.io/gorm` v1.25.12。代码入口：[`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go)。

## 场景与范围

这个独立的 `integration` 测试通过 [`sqlxdemo.InitDB`](../../internal/gostorage/sqlxdemo/sqlx_example.go) 建立连接，用 sqlx 插入一条记录；随后将同一连接交给 [`gormdemo.InitDBByExistDB`](../../internal/gostorage/gormdemo/public.go)，用 GORM 读取、更新，再用 sqlx 核对结果。测试还验证 GORM 事务返回错误后的回滚。它验证两套库可以共用连接池。默认 `go test ./...` 不会连接数据库。

仓库中原有的 [`gorm_test.go`](../../test/gostorage/gorm_test.go) 和 [`mysql_test.go`](../../test/gostorage/mysql_test.go) 涉及 `users`、`company` 等已有业务表，部分会修改数据；它们不属于这条最小验证路径。不要一次运行整个 `integration` 测试集合。

## 准备测试数据库

在本机 MySQL 中创建专用数据库和用户，或使用已有的**专用测试数据库**。以下 SQL 中的密码只是占位符，实际值自行设置：

```sql
CREATE DATABASE go_demo CHARACTER SET utf8mb4;
CREATE USER 'go_demo'@'127.0.0.1' IDENTIFIED BY '替换为本机测试密码';
GRANT SELECT, INSERT, UPDATE, CREATE, DROP ON go_demo.* TO 'go_demo'@'127.0.0.1';
```

在仓库根目录复制 [`configs/conf_demo.json`](../../configs/conf_demo.json) 为已被 Git 忽略的 `configs/conf.json`，将 `mysql` 字段填为上述账号、密码、`host: "127.0.0.1"`、`port: 3306`、`prefix: "go_demo"`。`prefix` 是数据库名。然后运行：

```bash
export GO_DEMO_CONFIG="$PWD/configs/conf.json"
go test -tags=integration -run '^TestStorageRoundTrip$' -v ./test/gostorage
```

预期结果是 `PASS`。测试创建一个名为 `go_demo_music_<数字>` 的表，只读写这张表，正常结束时自动删除。若进程被强制终止，可能留下该表；先用 `SHOW TABLES LIKE 'go_demo_music_%';` 核对，再对确认属于本测试的表执行 `DROP TABLE 表名;`。

## 边界

连接失败时先检查 `GO_DEMO_CONFIG` 是否指向有效 JSON、数据库是否存在、账号权限及端口是否正确。该测试要求 MySQL 账号能在专用数据库中创建和删除表；它不验证旧业务表、GORM Gen 或所有 sqlx 演示。密码不要写入版本控制，也不要把这组演示权限授予生产数据库。
