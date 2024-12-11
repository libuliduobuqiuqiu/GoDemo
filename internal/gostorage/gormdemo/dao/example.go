package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/model"
	"log"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

// insertOneRow 单挑记录插入
func insertOneRow() {
	var tmpUser *model.User
	db := gormdemo.GetDB()

	// 生成随机数据
	err := faker.FakeData(&tmpUser)
	if err != nil {
		log.Fatalf(err.Error())
	}

	result := db.Create(tmpUser)
	if result.Error != nil {
		log.Fatalf(result.Error.Error())
	}

	fmt.Println("result RowsAffected: ", result.RowsAffected)
	fmt.Printf("%+v\n", tmpUser)

}

// InsertRows 批量插入
func InsertRows() {
	db := gormdemo.GetDB()

	var users []*model.User
	for i := 0; i < 1000; i++ {
		tmpUser := &model.User{}
		err := faker.FakeData(&tmpUser)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, tmpUser)
	}

	result := db.Create(users)

	if result.Error != nil {
		log.Fatalf(result.Error.Error())
	}

	fmt.Println("RowsAffected: ", result.RowsAffected)

	for _, m := range users {
		fmt.Printf("%+v\n", m)
	}
}

// simpleQueryRow 简单查询
func simpleQueryRow() {
	db := gormdemo.GetDB()

	// 查询第一条记录（主键升序）
	var firstUser *model.User
	result := db.First(&firstUser)
	gormdemo.PrintRecord(firstUser, result)

	// 仅当有一个ID主键时，可直接定义User时把ID初始化
	firstIDUser2 := &model.User{ID: "e8efff22-a497-4a88-be1e-5123eb23ff75"}
	result = db.First(&firstIDUser2)
	gormdemo.PrintRecord(firstIDUser2, result)

	// 查询表中第一条记录（没有指定排序字段）
	var firstUser2 *model.User
	result = db.Take(&firstUser2)
	gormdemo.PrintRecord(firstUser2, result)

	// 查询表中最后一条记录（主键排序）
	var lastUser *model.User
	result = db.Last(&lastUser)
	gormdemo.PrintRecord(lastUser, result)

	// 查询当前所有记录
	var users []*model.User
	result = db.Find(&users)
	gormdemo.PrintRecords(users, result)

}

// condQueryRow 条件查询
func condQueryRow() {
	db := gormdemo.GetDB()

	// 查询当前username为condQueryRow的第一条记录（Struct方式）
	var tmpUser1 *model.User
	result := db.Where(&model.User{UserName: "qNptxqb"}).First(&tmpUser1)
	gormdemo.PrintRecord(tmpUser1, result)

	// 查询当前username为condQueryRow的第一条记录（Map方式）
	var tmpUser2 *model.User
	result = db.Where(map[string]interface{}{"username": "qNptxqb"}).First(&tmpUser2)
	gormdemo.PrintRecord(tmpUser2, result)

	// 指定Century查询字段查询记录
	var tmpUser3 []*model.User
	result = db.Where(&model.User{Century: "VII", UserName: "jaQlaFs"}, "Century").Find(&tmpUser3)
	gormdemo.PrintRecords(tmpUser3, result)

	// String 条件，直接写表达式
	var tmpUser4 *model.User
	result = db.Where("username = ?", "qNptxqb").First(&tmpUser4)
	gormdemo.PrintRecord(tmpUser4, result)

	var users []*model.User
	result = db.Where("date > ?", "2010-10-1").Find(&users)
	gormdemo.PrintRecords(users, result)

	// Order排序（默认升序）
	var users2 []*model.User
	result = db.Where("date > ?", "2010-10-1").Order("date").Find(&users2)
	gormdemo.PrintRecords(users2, result)

	// 查询特定的字段，不返回所有字段
	var tmpUser5 *model.User
	result = db.Select("username", "date").Where("username = ?", "qNptxqb").First(&tmpUser5)
	gormdemo.PrintRecord(tmpUser5, result)
}

type APIUser struct {
	ID        string `gorm:"primaryKey,column:id"`
	UserName  string `gorm:"column:username"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
}

// advancedQueryRow 高级查询
func advancedQueryRow() {
	db := gormdemo.GetDB()

	// 智能选择字段，如果经常只需要查询某些字段，可以重新定义小结构体
	var apiUser []APIUser
	result := db.Model(&model.User{}).Find(&apiUser)
	for _, user := range apiUser {
		fmt.Println(user)
	}
	fmt.Println(result.Error, result.RowsAffected)

	// 扫描结果绑定值map[string]interface{} 或者 []map[string]interface{}
	var users []map[string]interface{}
	result = db.Model(&model.User{}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.Error, result.RowsAffected)

	// Pluck查询单个列，并将结果扫描到切片
	var emails []string
	result = db.Model(&model.User{}).Pluck("email", &emails)
	fmt.Println(emails)
	fmt.Println(result.Error, result.RowsAffected)

	// Count查询
	var count int64
	result = db.Model(&model.User{}).Where("date > ?", "2012-10-22").Count(&count)
	fmt.Println(count)
	fmt.Println(result.Error, result.RowsAffected)
}

// updateRow 更新操作
func updateRow() {
	db := gormdemo.GetDB()

	// Save会保存所有字段，即使字段是零值，如果保存的值没有主键，就会创建，否则则是更新指定记录
	result := db.Save(&model.User{ID: "e8efff22-a497-4a88-be1e-5123eb23ff75", UserName: "zhangsan", Date: "2023-12-12"})
	fmt.Println(result.Error, result.RowsAffected)

	// 更新单个列
	result = db.Model(&model.User{}).Where("username = ?", "jaQlaFs").Update("first_name", "zhangsan")
	fmt.Println(result.Error, result.RowsAffected)

	// 更新多个列
	result = db.Model(&model.User{}).Where("username = ?", "zhangsan").Updates(model.User{FirstName: "zhangsan2", LastName: "zhangsan3"})
	fmt.Println(result.Error, result.RowsAffected)

	// 更新指定列(Select指定last_name)
	result = db.Model(&model.User{}).Where("username = ?", "zhangsan").Select("last_name").Updates(model.User{FirstName: "zhangsan2", LastName: "zhangsan4"})
	fmt.Println(result.Error, result.RowsAffected)
}

func deleteRows() {
	db := gormdemo.GetDB()

	// 指定匹配字段删除数据
	result := db.Delete(&model.User{}, map[string]interface{}{"username": "NJrauTj"})
	fmt.Println(result.Error, result.RowsAffected)

	result = db.Delete(&model.User{}, "username = ?", "NJrauTj")
	fmt.Println(result.Error, result.RowsAffected)

	// Where指定字段匹配删除数据
	result = db.Where("username = ? and phone_number = ?", "jXQKmPv", "574-821-9631").Delete(&model.User{})
	fmt.Println(result.Error, result.RowsAffected)

	// 批量删除的两种方式
	result = db.Where("email like ?", "%.com%").Delete(&model.User{})
	fmt.Println(result.Error, result.RowsAffected)

	result = db.Delete(&model.User{}, "email like ?", "%.com%")
	fmt.Println(result.Error, result.RowsAffected)
}

// execSQL 执行原生SQL语句
func execSQL() {
	db := gormdemo.GetDB()

	// 将查询SQL的结果映射到指定的单个变量中
	var oneUser model.User
	result := db.Raw("SELECT * FROM user LIMIT 1").Scan(&oneUser)
	fmt.Println(oneUser)
	fmt.Println(result.Error, result.RowsAffected)

	// 将查询SQL的批量结果映射到列表中
	var users []model.User
	result = db.Raw("SELECT * FROM user").Scan(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.Error, result.RowsAffected)

	var updateUser model.User
	result = db.Raw("UPDATE users SET username = ? where id = ?", "toms jobs", "ab6f089b-3272-49b5-858f-a93ed5a43b4f").Scan(&updateUser)
	fmt.Println(updateUser)
	fmt.Println(result.Error, result.RowsAffected)

	// 直接通过Exec函数执行Update操作，不返回任何查询结果？
	result = db.Exec("UPDATE user SET username = ? where id = ?", "toms jobs", "ab6f089b-3272-49b5-858f-a93ed5a43b4f")
	fmt.Println(result.Error, result.RowsAffected)

	// DryRun模式，在不执行的情况下生成SQL及其参数，可以用于准备或测试的SQL
	var tmpUsers []APIUser
	stmt := db.Session(&gorm.Session{DryRun: true}).Model(&model.User{}).Find(&tmpUsers).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(stmt.Vars)
}
