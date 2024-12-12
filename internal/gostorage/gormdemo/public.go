package gormdemo

import (
	"fmt"
	"godemo/pkg"
	"log"
	"math/rand"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	return
}

func InitDBWithNewLogger() (db *gorm.DB, err error) {
	rotateLogger := &lumberjack.Logger{
		Filename:   "gorm.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   true,
	}

	newLogger := logger.New(log.New(rotateLogger, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		})

	dsn := pkg.GenMysqlDSN("")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	db = db.Debug()
	return
}

// InitDBByExistDB
func InitDBByExistDB(existDB gorm.ConnPool) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: existDB,
	}))

	db = db.Debug()
	return
}
