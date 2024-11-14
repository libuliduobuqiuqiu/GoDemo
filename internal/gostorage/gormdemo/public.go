package gormdemo

import (
	"fmt"
	"godemo/pkg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	data := pkg.GetGlobalConfig("/data/GoDemo/configs/local_conf.json")
	config := data.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.Username,
		config.Password, config.Host, config.Port, config.Prefix)

	db, err = gorm.Open(mysql.Open(dsn))
	return
}
