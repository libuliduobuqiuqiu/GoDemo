# MySQL：sqlx 与 GORM 静态检查

> 状态：已完成静态检查与无数据库测试；未验证真实数据库运行。验证环境：Go 1.25，仓库使用 `github.com/jmoiron/sqlx` v1.3.5 与 `gorm.io/gorm` v1.25.12。代码入口：[`sqlxdemo`](../../internal/gostorage/sqlxdemo/)、[`gormdemo`](../../internal/gostorage/gormdemo/)；可选手动演示：[`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go)。

## 场景与范围

本阶段只检查数据库代码的构建、静态诊断和不连接数据库的逻辑。配置与 DSN 测试、sqlx 批量插入参数测试、GORM 模型关联解析及 `DryRun` SQL 测试均纳入默认 `go test ./...`。`DryRun` 检查生成的 SQL 和参数，不执行查询或事务。

独立的 `integration` 测试计划通过 [`sqlxdemo.InitDB`](../../internal/gostorage/sqlxdemo/sqlx_example.go) 建立连接，用 sqlx 插入记录；随后将同一连接交给 [`gormdemo.InitDBByExistDB`](../../internal/gostorage/gormdemo/public.go)，用 GORM 读取、更新并尝试验证回滚。它仍是未运行的手动演示，不能据此声称两套库已在真实 MySQL 上完成往返验证。

仓库中原有的 [`gorm_test.go`](../../test/gostorage/gorm_test.go) 和 [`mysql_test.go`](../../test/gostorage/mysql_test.go) 涉及 `users`、`company` 等已有业务表，部分会修改数据；它们不属于这条最小验证路径。不要一次运行整个 `integration` 测试集合。

## 无数据库验证

在仓库根目录运行：

```bash
go build ./...
go test -timeout 20s ./...
go test -tags=integration -run '^$' ./...
go vet ./internal/gostorage/... ./pkg
go vet -tags=integration ./test/gostorage
```

这些命令不需要 MySQL 服务。`integration` 命令只编译手动测试；`go vet` 查找静态问题，不能证明 SQL 在服务器上执行成功。

## 可选手动演示（不属于当前验收范围）

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

该演示预期创建一个名为 `go_demo_music_<数字>` 的表，只读写这张表，正常结束时尝试自动删除。上述流程未实际运行；若进程被强制终止，可能留下该表。手动运行者可用 `SHOW TABLES LIKE 'go_demo_music_%';` 核对，再对确认属于本测试的表执行 `DROP TABLE 表名;`。

## 边界

静态检查无法确认账号权限、建表、连接池复用、读写、事务回滚及清理在真实 MySQL 上的行为。旧业务表、GORM Gen 和其他手动演示也未做真实数据库验证。若自行运行可选演示，连接失败时检查 `GO_DEMO_CONFIG`、数据库、账号及端口；密码不要写入版本控制，也不要把演示权限授予生产数据库。
