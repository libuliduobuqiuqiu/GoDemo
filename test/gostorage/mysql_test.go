package gostorage

import (
	"fmt"
	"godemo/internal/gomanual/reflectdemo"
	"godemo/internal/gostorage/standardmysql"
	"godemo/pkg"
	"log"
	"testing"
)

func insertSqlByReflect() {
	d := reflectdemo.Device{Name: "zhangsan", Address: "127.0.0.1", Port: 8080}

	// Get Generate SQL string
	sql, sqlValues, err := reflectdemo.GenerateMysqlString(&d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)

	// Get Global Config about mysql
	globalConfig := pkg.GetGlobalConfig("")
	db, err := standardmysql.GetDB(globalConfig.MysqlConfig)

	// Exec sql string
	_, err = db.Exec(sql, sqlValues...)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSqlGenerate(t *testing.T) {
	insertSqlByReflect()
}

func TestUpdateUser(t *testing.T) {
	if err := standardmysql.UpdateUser(); err != nil {
		t.Fatal(err)
	}
}
