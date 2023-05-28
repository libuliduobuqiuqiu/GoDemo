package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	MysqlConfig DBConfig `json:"mysql"`
}

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"prefix"`
}

type Music struct {
	ID        int    `gorm:"column:id"`
	Author    string `gorm:"column:music_author"`
	Name      string `gorm:"column:music_name"`
	Album     string `gorm:"column:music_album"`
	Time      string `gorm:"column:music_time"`
	MusicType string `gorm:"column:music_type"`
	Lyrics    string `gorm:"column:music_lyrics"`
	Arranger  string `gorm:"column:music_arranger"`
}

func (m Music) TableName() string {
	return "music_music"
}

func initDB() (db *gorm.DB, err error) {
	confPath := "conf.json"
	if _, err = os.Stat(confPath); err != nil {
		return
	}

	var config Config
	if err = configor.Load(&config, confPath); err != nil {
		return
	}

	// 新建Database Gorm连接
	mysqlConfig := &config.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	var musicList []Music
	result := db.Find(&musicList)

	if result.Error != nil {
		log.Fatalf(result.Error.Error())
		return
	}

	fmt.Println("music_music 总计： ", result.RowsAffected)

	for _, music_item := range musicList {

		fmt.Printf("%+v\n", music_item)

	}

}
