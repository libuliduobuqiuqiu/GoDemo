package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormgendemo/query"
	"log"
)

func ListUsers() {

	db, err := gormdemo.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	query.SetDefault(db)
	users, err := query.User.Find()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Println(u)
	}

	fmt.Println(len(users))
}
