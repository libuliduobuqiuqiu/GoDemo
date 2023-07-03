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
	gormDemo.SelectCompanyRows(db)
}
