package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/model"
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
