package main

import (
	"fmt"
	"godemo/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var mysqlDSN string

const (
	confPath = "/data/GoDemo/configs/local_conf.json"
)

func InitLocalMysqlStr() {
	config := pkg.GetGlobalConfig(confPath)
	mysqlConfig := config.MysqlConfig
	mysqlDSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Prefix)
}

func main() {
	InitLocalMysqlStr()

	if mysqlDSN == "" {
		return
	}

	db, err := gorm.Open(mysql.Open(mysqlDSN))
	if err != nil {
		log.Fatal(err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/gostorage/gormdemo/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)
	// g.ApplyBasic()

	g.Execute()
}
