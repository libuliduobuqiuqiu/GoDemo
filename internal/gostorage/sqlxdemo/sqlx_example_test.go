package sqlxdemo

import (
	"reflect"
	"strings"
	"testing"

	"godemo/internal/gostorage/sqlxdemo/model"
)

func TestBuildBatchInsertMusic(t *testing.T) {
	rows := []model.Music{
		{Author: "a", Name: "first", Album: "one", Time: "2024-01-01", MusicType: "pop", Lyrics: "l1", Arranger: "r1"},
		{Author: "b", Name: "second", Album: "two", Time: "2024-01-02", MusicType: "rock", Lyrics: "l2", Arranger: "r2"},
	}
	query, args, err := buildBatchInsertMusic(rows)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(query, "VALUES (?, ?, ?, ?, ?, ?, ?),(?, ?, ?, ?, ?, ?, ?)") {
		t.Fatalf("unexpected SQL: %s", query)
	}
	want := []interface{}{"a", "first", "one", "2024-01-01", "pop", "l1", "r1", "b", "second", "two", "2024-01-02", "rock", "l2", "r2"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("args = %#v, want %#v", args, want)
	}
	if strings.Count(query, "?") != len(args) {
		t.Fatalf("placeholder count = %d, argument count = %d", strings.Count(query, "?"), len(args))
	}
}

func TestEmptyMusicInputDoesNotConnect(t *testing.T) {
	if _, _, err := buildBatchInsertMusic(nil); err == nil {
		t.Fatal("empty batch should return an error")
	}
	rows, err := QueryByNames(nil)
	if err != nil || len(rows) != 0 {
		t.Fatalf("empty name query = %#v, %v", rows, err)
	}
}

func TestBuildMusicNamesQuery(t *testing.T) {
	query, args, err := buildMusicNamesQuery([]string{"first", "second"})
	if err != nil {
		t.Fatal(err)
	}
	if want := "select * from music_music where music_name in (?, ?)"; query != want {
		t.Fatalf("query = %q, want %q", query, want)
	}
	if want := []interface{}{"first", "second"}; !reflect.DeepEqual(args, want) {
		t.Fatalf("args = %#v, want %#v", args, want)
	}
	if _, _, err := buildMusicNamesQuery(nil); err == nil {
		t.Fatal("sqlx.In should reject an empty list")
	}
}
