package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/model"
	"log"
)

func ListUsers() {
	db, err := gormdemo.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	var users []*model.User
	db = db.Table(model.UserTableName)
	fmt.Println(model.UserTableName)
	if err := db.Find(&users).Error; err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(len(users))

	for _, u := range users {
		fmt.Println(u)
	}
}
