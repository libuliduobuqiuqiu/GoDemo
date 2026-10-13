//go:build integration

package gostorage

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/sqlxdemo"
	"gorm.io/gorm"
)

// TestStorageRoundTrip uses an isolated table in the configured test database.
func TestStorageRoundTrip(t *testing.T) {
	db, err := sqlxdemo.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	table := fmt.Sprintf("go_demo_music_%d", time.Now().UnixNano())
	created := false
	t.Cleanup(func() {
		if created {
			if _, err := db.Exec("DROP TABLE " + table); err != nil {
				t.Errorf("drop test table %s: %v", table, err)
			}
		}
		if err := db.Close(); err != nil {
			t.Errorf("close database: %v", err)
		}
	})

	createTable := "CREATE TABLE " + table + " (" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, " +
		"music_author VARCHAR(100) NOT NULL, " +
		"music_name VARCHAR(100) NOT NULL) ENGINE=InnoDB"
	if _, err := db.Exec(createTable); err != nil {
		t.Fatal(err)
	}
	created = true

	result, err := db.Exec("INSERT INTO "+table+" (music_author, music_name) VALUES (?, ?)", "demo", "first")
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	gormDB, err := gormdemo.InitDBByExistDB(db)
	if err != nil {
		t.Fatal(err)
	}
	var row struct {
		ID     int64  `db:"id" gorm:"column:id"`
		Author string `db:"music_author" gorm:"column:music_author"`
		Name   string `db:"music_name" gorm:"column:music_name"`
	}
	if err := gormDB.Table(table).Where("id = ?", id).Take(&row).Error; err != nil {
		t.Fatal(err)
	}
	if row.ID != id || row.Author != "demo" || row.Name != "first" {
		t.Fatalf("GORM read unexpected row: %+v", row)
	}

	if err := gormDB.Table(table).Where("id = ?", id).Update("music_name", "updated").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Get(&row, "SELECT id, music_author, music_name FROM "+table+" WHERE id = ?", id); err != nil {
		t.Fatal(err)
	}
	if row.Name != "updated" {
		t.Fatalf("sqlx read music_name = %q, want updated", row.Name)
	}

	rollbackErr := errors.New("rollback demonstration")
	err = gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(table).Where("id = ?", id).Update("music_name", "rolled back").Error; err != nil {
			return err
		}
		return rollbackErr
	})
	if !errors.Is(err, rollbackErr) {
		t.Fatalf("transaction error = %v, want %v", err, rollbackErr)
	}
	if err := db.Get(&row, "SELECT id, music_author, music_name FROM "+table+" WHERE id = ?", id); err != nil {
		t.Fatal(err)
	}
	if row.Name != "updated" {
		t.Fatalf("rolled-back music_name = %q, want updated", row.Name)
	}
}
