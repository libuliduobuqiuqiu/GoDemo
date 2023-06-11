package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/jinzhu/configor"
	"github.com/pborman/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func printRecord(u *User, result *gorm.DB) {
	fmt.Printf("%v\n", u)
	fmt.Println(result.Error, result.RowsAffected)
}

func printRecords(u []User, result *gorm.DB) {

	for _, u := range u {
		fmt.Println(u)
	}
	fmt.Println(result.Error, result.RowsAffected)

}

type Config struct {
	MysqlConfig DBConfig `json:"mysql"`
}

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"prefix"`
}

type User struct {
	ID          string `gorm:"primaryKey,column:id" faker:"-"`
	Email       string `gorm:"column:email" faker:"email"`
	Password    string `gorm:"column:password" faker:"password"`
	PhoneNumber string `gorm:"column:phone_number" faker:"phone_number"`
	UserName    string `gorm:"column:username" faker:"username"`
	FirstName   string `gorm:"first_name" faker:"first_name"`
	LastName    string `gorm:"last_name" faker:"last_name"`
	Century     string `gorm:"century" faker:"century"`
	Date        string `gorm:"date" faker:"date"`
}

func (u User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}

func initDB() (db *gorm.DB, err error) {
	confPath := "conf.json"
	if _, err = os.Stat(confPath); err != nil {
		return
	}

	var config Config
	if err = configor.Load(&config, confPath); err != nil {
		return
	}

	// 新建Database Gorm连接
	mysqlConfig := &config.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}

// insertOneRow 单挑记录插入
func insertOneRow(db *gorm.DB) {
	var tmpUser *User

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

// insertRows 批量插入
func insertRows(db *gorm.DB) {
	var users []*User
	for i := 0; i < 10; i++ {
		tmpUser := User{}
		err := faker.FakeData(&tmpUser)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, &tmpUser)
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
func simpleQueryRow(db *gorm.DB) {

	// 查询第一条记录（主键升序）
	var firstUser *User
	result := db.First(&firstUser)
	printRecord(firstUser, result)

	// 仅当有一个ID主键时，可直接定义User时把ID初始化
	firstIDUser2 := &User{ID: "e8efff22-a497-4a88-be1e-5123eb23ff75"}
	result = db.First(&firstIDUser2)
	printRecord(firstIDUser2, result)

	// 查询表中第一条记录（没有指定排序字段）
	var firstUser2 *User
	result = db.Take(&firstUser2)
	printRecord(firstUser2, result)

	// 查询表中最后一条记录（主键排序）
	var lastUser *User
	result = db.Last(&lastUser)
	printRecord(lastUser, result)

	// 查询当前所有记录
	var users []User
	result = db.Find(&users)
	printRecords(users, result)

}

// condQueryRow 条件查询
func condQueryRow(db *gorm.DB) {

	// 查询当前username为condQueryRow的第一条记录（Struct方式）
	var tmpUser1 *User
	result := db.Where(&User{UserName: "qNptxqb"}).First(&tmpUser1)
	printRecord(tmpUser1, result)

	// 查询当前username为condQueryRow的第一条记录（Map方式）
	var tmpUser2 *User
	result = db.Where(map[string]interface{}{"username": "qNptxqb"}).First(&tmpUser2)
	printRecord(tmpUser2, result)

	// 指定Century查询字段查询记录
	var tmpUser3 []User
	result = db.Where(&User{Century: "VII", UserName: "jaQlaFs"}, "Century").Find(&tmpUser3)
	printRecords(tmpUser3, result)

	// String 条件，直接写表达式
	var tmpUser4 *User
	result = db.Where("username = ?", "qNptxqb").First(&tmpUser4)
	printRecord(tmpUser4, result)

	var users []User
	result = db.Where("date > ?", "2010-10-1").Find(&users)
	printRecords(users, result)

	// Order排序（默认升序）
	var users2 []User
	result = db.Where("date > ?", "2010-10-1").Order("date").Find(&users2)
	printRecords(users2, result)

	// 查询特定的字段，不返回所有字段
	var tmpUser5 *User
	result = db.Select("username", "date").Where("username = ?", "qNptxqb").First(&tmpUser5)
	printRecord(tmpUser5, result)
}

type APIUser struct {
	ID        string `gorm:"primaryKey,column:id"`
	UserName  string `gorm:"column:username"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
}

// advancedQueryRow 高级查询
func advancedQueryRow(db *gorm.DB) {

	// 智能选择字段，如果经常只需要查询某些字段，可以重新定义小结构体
	var apiUser []APIUser
	result := db.Model(&User{}).Find(&apiUser)
	for _, user := range apiUser {
		fmt.Println(user)
	}
	fmt.Println(result.Error, result.RowsAffected)

	// 扫描结果绑定值map[string]interface{} 或者 []map[string]interface{}
	var users []map[string]interface{}
	result = db.Model(&User{}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.Error, result.RowsAffected)

	// Pluck查询单个列，并将结果扫描到切片
	var emails []string
	result = db.Model(&User{}).Pluck("email", &emails)
	fmt.Println(emails)
	fmt.Println(result.Error, result.RowsAffected)

	// Count查询
	var count int64
	result = db.Model(&User{}).Where("date > ?", "2012-10-22").Count(&count)
	fmt.Println(count)
	fmt.Println(result.Error, result.RowsAffected)
}

func updateRow(db *gorm.DB) {
	// Save会保存所有字段，即使字段是零值，如果保存的值没有主键，就会创建，否则则是更新指定记录
	result := db.Save(&User{ID: "e8efff22-a497-4a88-be1e-5123eb23ff75", UserName: "zhangsan", Date: "2023-12-12"})
	fmt.Println(result.Error, result.RowsAffected)

	// 更新单个列
	result = db.Model(&User{}).Where("username = ?", "jaQlaFs").Update("first_name", "zhangsan")
	fmt.Println(result.Error, result.RowsAffected)

	// 更新多个列
	result = db.Model(&User{}).Where("username = ?", "zhangsan").Updates(User{FirstName: "zhangsan2", LastName: "zhangsan3"})
	fmt.Println(result.Error, result.RowsAffected)

	// 更新指定列(Select指定last_name)
	result = db.Model(&User{}).Where("username = ?", "zhangsan").Select("last_name").Updates(User{FirstName: "zhangsan2", LastName: "zhangsan4"})
	fmt.Println(result.Error, result.RowsAffected)
}

func main() {

	db, err := initDB()
	db = db.Debug()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	updateRow(db)
}
