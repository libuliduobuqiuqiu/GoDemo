package gormdemo

import (
	"reflect"
	"strings"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRoundTripSQLWithoutDatabase(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "demo:secret@tcp(127.0.0.1:3306)/go_demo",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{DryRun: true, DisableAutomaticPing: true, SkipDefaultTransaction: true})
	if err != nil {
		t.Fatal(err)
	}

	const table = "go_demo_music_123"
	var row struct {
		ID     int64  `gorm:"column:id"`
		Author string `gorm:"column:music_author"`
		Name   string `gorm:"column:music_name"`
	}
	read := db.Table(table).Where("id = ?", int64(42)).Take(&row)
	if read.Error != nil {
		t.Fatal(read.Error)
	}
	if sql := read.Statement.SQL.String(); !strings.Contains(sql, "FROM `"+table+"`") || !strings.Contains(sql, "WHERE id = ?") {
		t.Fatalf("unexpected read SQL: %s", sql)
	}
	if !reflect.DeepEqual(read.Statement.Vars, []interface{}{int64(42), 1}) {
		t.Fatalf("read args = %#v", read.Statement.Vars)
	}

	update := db.Table(table).Where("id = ?", int64(42)).Update("music_name", "updated")
	if update.Error != nil {
		t.Fatal(update.Error)
	}
	if sql := update.Statement.SQL.String(); !strings.Contains(sql, "UPDATE `"+table+"`") || !strings.Contains(sql, "`music_name`=?") || !strings.Contains(sql, "WHERE id = ?") {
		t.Fatalf("unexpected update SQL: %s", sql)
	}
	if !reflect.DeepEqual(update.Statement.Vars, []interface{}{"updated", int64(42)}) {
		t.Fatalf("update args = %#v", update.Statement.Vars)
	}
}
