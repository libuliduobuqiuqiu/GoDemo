package SqlxDemo

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Config struct {
	Mysql DBConfig `json:"mysql"`
}

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Prefix   string `json:"prefix"`
}

type Music struct {
	ID        int    `db:"id"`
	Author    string `db:"music_author"`
	Name      string `db:"music_name"`
	Album     string `db:"music_album"`
	Time      string `db:"music_time"`
	MusicType string `db:"music_type"`
	Lyrics    string `db:"music_lyrics"`
	Arranger  string `db:"music_arranger"`
}

// InitDB 初始化DB连接
func InitDB() error {
	// 读取conf.json配置文件
	confDir := "conf.json"
	_, err := os.Stat(confDir)
	if err != nil {
		return err
	}

	// 解析conf.json配置文件
	var conf Config
	err = configor.Load(&conf, confDir)
	if err != nil {
		return err
	}

	// 新建DB连接
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host,
		conf.Mysql.Prefix)
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

// SimpleQueryRow 数据库操作
func SimpleQueryRow() {
	err := InitDB()
	if err != nil {
		fmt.Printf("Connect to Database failed: %v \n", err)
		return
	}

	sqlStr := "select * from music_music"
	var music []Music
	err = db.Select(&music, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err: %v \n", err)
		return
	}
	for _, v := range music {
		fmt.Println(v)
	}
}
