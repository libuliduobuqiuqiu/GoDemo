package model

import "database/sql/driver"

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

func (m Music) Value() (driver.Value, error) {
	return []interface{}{m.Author, m.Name, m.Album, m.Time, m.MusicType, m.Lyrics, m.Arranger}, nil
}
