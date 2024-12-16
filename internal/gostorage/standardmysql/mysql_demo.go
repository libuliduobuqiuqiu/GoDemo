package standardmysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godemo/pkg"
	"log"
	"strings"
)

func GetDB(config pkg.MysqlConfig) (db *sql.DB, err error) {
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Prefix))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ExecSQLStr() {
	globalConfig := pkg.GetGlobalConfig("")
	mysqlConfig := globalConfig.MysqlConfig

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Prefix))
	if err != nil {
		log.Fatal(err)
	}

	data := [][]string{
		{"4", "test4"},
		{"5", "test5"},
		{"6", "test6"},
	}

	var valueStrings []string
	var values []interface{}

	for _, d := range data {
		valueStrings = append(valueStrings, "(?, ?)")
		values = append(values, d[0], d[1])
	}

	sql := "INSERT INTO test_table(id, data) VALUES " + strings.Join(valueStrings, ",")
	_, err = db.Exec(sql, values...)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser() error {
	dsn := pkg.GenMysqlDSN("")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	sql := "update users set last_name = 'Gleasons' where id = 'fff2a321-cfb7-48ae-bad1-935f99d9fe75'"
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(rowsAffected)
	return nil
}
