package main

import (
	"log"
	"sunrun/GormDemo"
)

func main() {

	db, err := gormDemo.InitDB()
	db = db.Debug()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	gormDemo.InsertRows(db) // 批量插入1000条数据
}
