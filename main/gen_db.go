package main

import (
	"fmt"
	"log"
	gormDemo "sunrun/GormDemo"
)

func GenRenameSQL() {
	db, err := gormDemo.InitDB("company2")
	if err != nil {
		log.Fatal(fmt.Sprintf("Init DB Failed: ", err))
	}

	gormDemo.GenerateSQL(db)
}

func GenShowTableStructSQL() {
	db, err := gormDemo.InitDB("company2")
	if err != nil {
		log.Fatal(err)
	}
	tableNames := gormDemo.FilterADTable(db)

	cmdbConnect, err := gormDemo.InitDB("company")
	if err != nil {
		log.Fatal(err)
	}
	gormDemo.FilterTableField(cmdbConnect, tableNames)
}
