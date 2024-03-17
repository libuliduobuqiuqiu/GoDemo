package main

import (
	"fmt"
	"log"
	"sunrun/goconcurrency"
	methoddemo "sunrun/gomanual/methodDemo"
	"sunrun/gomanual/reflectdemo"
	"sunrun/gostorage/standardmysql"
	"sunrun/public"
)

func startChat() {
	goconcurrency.StartChat()
}

func insertSqlByReflect() {
	d := reflectdemo.Device{Name: "zhangsan", Address: "127.0.0.1", Port: 8080}

	// Get Generate SQL string
	sql, sqlValues, err := reflectdemo.GenerateMysqlString(&d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)

	// Get Global Config about mysql
	globalConfig := public.GetGlobalConfig()
	db, err := standardmysql.GetDB(globalConfig.MysqlConfig)

	// Exec sql string
	_, err = db.Exec(sql, sqlValues...)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// methoddemo.MethodUseDiffReceiver()
	// regexdemo.InterpreteDataGroup()
	methoddemo.ScanDevice()
}
