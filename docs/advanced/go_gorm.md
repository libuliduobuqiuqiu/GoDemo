# GORM：连接、查询与事务

> 状态：核心 API 已复核，待真实数据库验证；适用版本：仓库的 `gorm.io/gorm` v1.25.12、`gorm.io/driver/mysql` v1.5.7；代码入口：[`gormdemo`](../../internal/gostorage/gormdemo/) 与 [`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go)。

## 场景

仓库用 GORM 演示 MySQL 模型映射、查询、关联和事务。先运行[独立数据库往返测试](go_database.md)，确认配置和连接可用，再按需阅读依赖 `users`、`company` 等预置表的旧演示。这里使用 GORM 的传统 `*gorm.DB` API；官方文档也展示了较新的泛型 API，本仓库尚未采用。

## 连接与模型

[`gormdemo.InitDB`](../../internal/gostorage/gormdemo/public.go) 从 `GO_DEMO_CONFIG` 指定的 JSON 读取 MySQL 配置，再创建 GORM 连接；[`InitDBByExistDB`](../../internal/gostorage/gormdemo/public.go) 可复用已有的 `database/sql` 连接池，往返测试通过 sqlx 的 `*sqlx.DB` 验证此路径。调用方应处理初始化错误，并在持有底层连接池时负责关闭它。旧演示中的 `GetDB()` 在连接失败时调用 `log.Fatal`，不适合在需要恢复错误的库代码中沿用。

[`model.User`](../../internal/gostorage/gormdemo/model/user.go) 明确设置主键、列名和关联。GORM 默认把 `ID` 识别为主键，把字段名转成蛇形列名；`CreatedAt`、`UpdatedAt` 字段只有在模型中定义时才会自动维护。自定义列名用 `gorm:"column:列名"`，多个设置用分号分隔，例如 `gorm:"primaryKey;column:id"`。本仓库的 `User` 没有 `CreatedAt` 或 `UpdatedAt` 字段。[官方约定说明](https://gorm.io/docs/conventions.html)

## 查询与错误处理

传统 API 在执行查询后检查返回的 `*gorm.DB.Error`；`First`、`Take` 找不到记录会返回 `gorm.ErrRecordNotFound`，`Find` 查询结果为空时通常应看 `RowsAffected` 或切片长度。仓库的 [`dao.GetUsers`](../../internal/gostorage/gormdemo/dao/user.go) 展示 `Select`、`Limit`、`Where` 和 `Take`，但它依赖已有的 `users` 表及匹配 `email` 的数据，不适合作为首次运行命令。[官方错误处理说明](https://gorm.io/docs/error_handling.html)

```go
var user model.User
err := db.Where("email = ?", email).Take(&user).Error
if errors.Is(err, gorm.ErrRecordNotFound) {
    // 记录不存在
} else if err != nil {
    // 连接、SQL 或其他错误
}
```

参数值使用 `?` 占位符传入。表名、列名等 SQL 标识符不能作为值参数传入；独立往返测试只用程序生成的固定前缀表名。

## 事务与关联

`db.Transaction` 的回调返回 `nil` 时提交，返回错误时回滚；回调中的操作必须使用传入的 `tx`。[`TestStorageRoundTrip`](../../test/gostorage/roundtrip_test.go) 在事务里更新后故意返回错误，再用 sqlx 核对值未改变。旧的 [`dao.InsertUserByTrans`](../../internal/gostorage/gormdemo/dao/user.go) 也故意返回回滚错误，但随后尝试读取已经回滚的记录，因此不应用作成功路径演示。[官方事务说明](https://gorm.io/docs/transactions.html)

```go
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Model(&user).Update("username", "new-name").Error; err != nil {
        return err
    }
    return nil
})
```

[`dao.UseGormPreload`](../../internal/gostorage/gormdemo/dao/user.go) 对 `Company.Users` 使用 `Preload`，它通过额外查询加载关联；[`dao.UseGormJoin`](../../internal/gostorage/gormdemo/dao/user.go) 对 `User.Company` 使用 `Joins`，通过联接查询加载单值关联。两者依赖既有表和外键数据，尚未纳入独立往返测试。[官方预加载说明](https://gorm.io/docs/preload.html)

## 验证与边界

准备专用 MySQL 数据库后，按[数据库实践](go_database.md)配置并执行：

```bash
go test -tags=integration -run '^TestStorageRoundTrip$' -v ./test/gostorage
```

默认 `go test ./...` 不连接数据库。此测试覆盖连接池复用、读取、更新和事务回滚；旧业务表上的关联、GORM Gen、钩子、日志与所有 DAO 函数仍需逐项验证，不能由该测试推断为已通过。
