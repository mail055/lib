package web

import (
	"database/sql"
	"html/template"
)

var db *sql.DB
var tmpls *template.Template

type PageData struct {
	Albums      []AlbumStruct
	SearchQuery string
	AlbumInfo   *AlbumInfoStruct
	Tracks      []TrackStruct
}

type AlbumStruct struct {
	ID       string
	Title    string
	Artists  string
	ImageURL string
}

type TrackStruct struct {
	ID          string
	TrackNumber int
	Title       string
	Artists     string
	DurationMS  int
	Duration    string
	TrackURL    string
	DiscNumber  int
	Explicit    int
}

type AlbumInfoStruct struct {
	ID          string
	Title       string
	AlbumType   string
	Artists     string
	TotalTracks int
	DurationMS  int
	Duration    string
	URL         string
	ImageURL    string
	ReleaseDate string
}
