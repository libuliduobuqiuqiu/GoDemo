package gormgendemo

import (
	"fmt"
	"godemo/internal/gostorage/gormgendemo/query"
	"godemo/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ExecBookDemo() {
	mysqlStr, err := pkg.GenMysqlDSN("")
	if err != nil {
		log.Print(err)
		return
	}

	db, err := gorm.Open(mysql.Open(mysqlStr))

	if err != nil {
		log.Fatal(err)
	}

	query.SetDefault(db)
	books, err := query.Book.Find()
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
}
