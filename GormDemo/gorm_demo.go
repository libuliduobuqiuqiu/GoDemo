package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
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
	ID        int    `gorm:"column:id" faker:"-"`
	Author    string `gorm:"column:music_author" faker:"name"`
	Name      string `gorm:"column:music_name" faker:"username"`
	Album     string `gorm:"column:music_album" faker:"word"`
	Time      string `gorm:"column:music_time" faker:"time"`
	MusicType string `gorm:"column:music_type" faker:"title_male"`
	Lyrics    string `gorm:"column:music_lyrics" faker:"word"`
	Arranger  string `gorm:"column:music_arranger" faker:"word"`
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

func queryRecords(db *gorm.DB) {

	music := Music{}
	db.First(&music)
	fmt.Printf("%+v\n", music)

	db.Take(&music)
	fmt.Printf("%+v\n", music)

	db.Last(&music)
	fmt.Printf("%+v\n", music)

}

func insertRecord(db *gorm.DB) {

	var musicList []*Music
	for i := 0; i < 10; i++ {
		tempMusic := Music{}
		err := faker.FakeData(&tempMusic)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		musicList = append(musicList, &tempMusic)
	}

	result := db.Create(musicList)
	for _, m := range musicList {
		fmt.Printf("%+v", m)
	}
	fmt.Println(result.Error, result.RowsAffected)

}

func main() {

	db, err := initDB()
	db = db.Debug()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	queryRecords(db)
	insertRecord(db)
}
