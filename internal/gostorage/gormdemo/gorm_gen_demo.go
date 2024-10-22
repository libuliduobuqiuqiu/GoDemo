package gormdemo

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo/query"
	"godemo/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ExecBookDemo() {
	var confPath = "/data/GoDemo/configs/local_conf.json"
	mysqlStr := pkg.GenMysqlDSN(confPath)

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
