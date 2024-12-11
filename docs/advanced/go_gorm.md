# Gorm

如何设计一个数据库引擎？需要具备哪些功能？
1. 数据库连接（支持不同类型数据进行创建、关闭、管理连接功能，支持连接池管理）
2. ORM核心功能：
    - 模型定义与映射
    - CRUD操作
    - 链式调用
    - 事务支持
3. 查询生成：
    - SQL构造器（动态构造复杂SQL语句）
    - 预编译，参数化（防止SQL注入）
    - 分页、排序、条件查询
4. 模型验证与钩子：
    - 字段验证
5. 高级功能：
    - 支持一对多，多对多关联关系
    - 预加载数据
    - 自定义SQL执行
6. 日志和调试
    - SQL查询日志
    - 支持性能分析


## 数据库连接

mysql两种方式建立连接：
1. 通过dsn（data source name)新建一个mysql连接,然后初始化*gorm.DB;
2. 通过已存在的mysql连接，初始化*gorm.DB;


## ORM核心功能

### 基础CRUD操作

基础的增删改查操作？

### 模型定义

如何模型定义，字段标签的作用？
1. 模型用普通结构体定义，
    - 使用一个默认ID为主键
    - 表名默认将结构体名转换为snake_case结构体并且加上复数
    - 列名默认将字段名转换为snake_case结构体
    - Gorm使用字段CreateAt和UpdateAt用来自动跟踪记录中的创建和更新时间
2. 常用字段标签：
    - column
    - type
    - primary key
    - unique key
    - default
    - comment

### 事务支持

怎么进行事务操作？怎么在事务中执行多个操作？事务底层原理？
调用事务函数，自定义函数中执行的业务逻辑,函数中可以执行多个操作，当某个操作异常或者业务逻辑返回错误，回滚之前操作。
```go
err = db.Transaction(func(tx *gorm.DB) (err error) {
}
```
还可以通过手动执行事务
```go
tx := db.Begin()
// db操作
tx.Create()
// 错误回滚
tx.Rollback()
// 提交操作
tx.Commit()
```

## 查询生成

### 链式调用

链式调用什么用，链式调用怎么用？
简单来说链式调用，更加简洁优雅，同时能让逻辑代码保持连贯性；
```go
  if err := db.Select("id", "email", "username").Limit(10).Find(&userList).Error; err != nil {
		return err
	}

	if err := db.Where("email like ?", "%com").Take(&user).Error; err != nil {
		return err
	}
```
链式调用主要分为三部分：Chain Methods,Finsher Methods, New Session Methods.

### SQL构建器

如何执行原生SQL？DryRun模式有什么用？
```go
// 原生SQL查询
db.Raw().Scan()
// 原生SQL执行
db.Exec()
```
DryRun模式可以在不执行的情况下，生成SQL;
```go
stmt := db.Session(&gorm.Session{DryRun: true}).Where("email like ?", "%.com").Find(&userList).Statement
fmt.Println(stmt.SQL.String())
fmt.Println(stmt.Vars)
```

### 模型验证与钩子

### 钩子
钩子是在什么？在什么时候执行？如果钩子出现异常之后，会导致什么结果？
钩子就是回调函数，可以在创建、更新、删除操作前后调用；
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
```
Gorm中操作默认是开启事务，当钩子出现异常时会对操作进行回滚。

## 高级功能

### 预加载
