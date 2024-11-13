package main

import (
	"godemo/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const (
	confPath = "/data/GoDemo/configs/local_conf.json"
)

func main() {
	mysqlDSN := pkg.GenMysqlDSN(confPath)

	if mysqlDSN == "" {
		return
	}

	db, err := gorm.Open(mysql.Open(mysqlDSN))
	if err != nil {
		log.Fatal(err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/gostorage/gormgendemo/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("book"),
		g.GenerateModelAs("model", "MyModel"),
		g.GenerateModel("history"),
		g.GenerateModel("device"))
	// g.ApplyBasic(g.GenerateModel("book"), g.GenerateModel("history"))
	// g.ApplyBasic()

	g.Execute()
}
