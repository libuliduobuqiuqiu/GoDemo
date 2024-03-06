package standardmysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"sunrun/public"
)

func GetDB(config public.MysqlConfig) (db *sql.DB, err error) {
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Prefix))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ExecSQLStr() {
	globalConfig := public.GetGlobalConfig()
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
