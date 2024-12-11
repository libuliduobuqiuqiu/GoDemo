package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/model"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

// 不带模型
func ListUsersByTableName() error {
	db, err := gormdemo.InitDB()
	if err != nil {
		return err
	}

	var users = []map[string]interface{}{}
	fmt.Println(len(users))
	db = db.Table(model.UserTableName)
	result := db.Find(&users)
	if err := result.Error; err != nil {
		return err
	}

	fmt.Println(len(users))
	return nil
}

// 带模型
func ListUsersByNotTableName() error {

	db, err := gormdemo.InitDB()
	if err != nil {
		return err
	}

	var users []*model.User
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	fmt.Println(len(users))
	return nil
}

func InsertUser() error {
	db := gormdemo.GetDB()

	var tmpUser model.User
	if err := faker.FakeData(&tmpUser); err != nil {
		return err
	}

	if err := db.Create(&tmpUser).Error; err != nil {
		return err
	}

	return nil
}

func InsertUserByTrans() error {
	var tmpID string

	db := gormdemo.GetDB()
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		tmpUser := &model.User{}
		if err = faker.FakeData(tmpUser); err != nil {
			return
		}

		if err = tx.Create(tmpUser).Error; err != nil {
			return
		}

		tmpID = tmpUser.ID

		// 测试是否回滚
		return fmt.Errorf("test rollback")
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	var user = &model.User{}
	user.ID = tmpID
	if err = db.Take(user).Error; err != nil {
		return err
	}
	fmt.Println("Insert Success", user)

	return err
}

func GetUsers() error {
	db := gormdemo.GetDB()

	var (
		user     model.User
		userList []*model.User
	)

	if err := db.Select("id", "email", "username").Limit(10).Find(&userList).Error; err != nil {
		return err
	}

	if err := db.Where("email like ?", "%com").Take(&user).Error; err != nil {
		return err
	}

	for _, v := range userList {
		fmt.Println(v)
	}

	fmt.Println(user)
	return nil
}

func UseGromDryRun() {
	db := gormdemo.GetDB()

	var userList []*model.User
	stmt := db.Session(&gorm.Session{DryRun: true}).Where("email like ?", "%.com").Find(&userList).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(stmt.Vars)
}

func UpdateUserCompany() error {
	db := gormdemo.GetDB()

	var (
		userList    []*model.User
		companyList []*model.Company
		companyIds  []string
	)

	if err := db.Find(&userList).Error; err != nil {
		return err
	}

	if err := db.Select("id").Find(&companyList).Error; err != nil {
		return err
	}

	for _, v := range companyList {
		companyIds = append(companyIds, v.ID)
	}

	for _, user := range userList {
		if err := db.Table("users").Where("id = ? ", user.ID).Update("company_code", gormdemo.RandSlice(companyIds)).Error; err != nil {
			return err
		}
	}

	return nil
}
