package model

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
