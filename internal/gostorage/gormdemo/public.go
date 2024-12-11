package gormdemo

import (
	"fmt"
	"godemo/pkg"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func init() {
	tmp, err := InitDB()
	if err != nil {
		gormDB = tmp
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandSlice[T any](d []T) T {
	index := rand.Intn(len(d))
	return d[index]
}

func PrintRecord(record interface{}, result *gorm.DB) {
	fmt.Printf("%v\n", record)
	fmt.Println(result.Error, result.RowsAffected)

}

func PrintRecords[T any](records []T, results *gorm.DB) {
	for _, v := range records {
		fmt.Println(v)
	}

	fmt.Println("Rows：", results.RowsAffected, ", Error: ", results.Error)
}

func GetDB() *gorm.DB {
	var err error
	if gormDB == nil {
		gormDB, err = InitDB()
		if err != nil {
			log.Fatal(err)
		}
	}

	return gormDB
}

func InitDB() (db *gorm.DB, err error) {
	dsn := pkg.GenMysqlDSN("")
	db, err = gorm.Open(mysql.Open(dsn))
	db = db.Debug()
	return
}

func InitDBByExistDB(existDB gorm.ConnPool) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: existDB,
	}))

	db = db.Debug()
	return
}
