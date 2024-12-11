package model

import (
	"github.com/pborman/uuid"
	"gorm.io/gorm"
)

const UserTableName = "users"

type User struct {
	ID          string  `gorm:"primaryKey,column:id" faker:"-"`
	Email       string  `gorm:"column:email" faker:"email"`
	Password    string  `gorm:"column:password" faker:"password"`
	PhoneNumber string  `gorm:"column:phone_number" faker:"phone_number"`
	UserName    string  `gorm:"column:username" faker:"username"`
	FirstName   string  `gorm:"first_name" faker:"first_name"`
	LastName    string  `gorm:"last_name" faker:"last_name"`
	Century     string  `gorm:"century" faker:"century"`
	Date        string  `gorm:"date" faker:"date"`
	CompanyID   string  `gorm:"column:company_code" faker:"-"`
	Company     Company `gorm:"foreignKey:CompanyID;references:ID" faker:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	// 测试创建是否事务
	// fmt.Println(u.ID)
	// return fmt.Errorf("test rollback hook")
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// 测试创建后回调钩子失败后是否回滚
	// return fmt.Errorf("test after create rollback hook.")
	return nil
}
