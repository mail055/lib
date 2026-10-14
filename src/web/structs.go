package web

import (
	"database/sql"
	"html/template"
)

var db *sql.DB
var tmpls *template.Template

const schema = `CREATE TABLE IF NOT EXISTS album_artists (
    album_id  TEXT REFERENCES albums (album_id) ON DELETE CASCADE
                   NOT NULL,
    artist_id TEXT REFERENCES artists (artist_id) 
                   NOT NULL,
    PRIMARY KEY (
        album_id,
        artist_id
    )
    ON CONFLICT IGNORE
);

CREATE TABLE IF NOT EXISTS albums (
    album_id     TEXT    PRIMARY KEY,
    album_name   TEXT,
    album_type   TEXT,
    total_tracks INTEGER,
    duration_ms  INTEGER,
    album_url    TEXT,
    album_image  TEXT,
    release_date TEXT
);

CREATE TABLE IF NOT EXISTS artist_tracks (
    artist_id TEXT REFERENCES artists (artist_id) 
                   NOT NULL,
    track_id  TEXT REFERENCES tracks (track_id) ON DELETE CASCADE
                   NOT NULL,
    PRIMARY KEY (
        artist_id,
        track_id
    )
    ON CONFLICT IGNORE
);

CREATE TABLE IF NOT EXISTS artists (
    artist_id   TEXT PRIMARY KEY
                     UNIQUE ON CONFLICT IGNORE,
    artist_name TEXT,
    artist_url  TEXT
);

CREATE TABLE IF NOT EXISTS tracks (
    track_id     TEXT    PRIMARY KEY,
    track_name   TEXT,
    track_url    TEXT,
    duration_ms  INTEGER,
    explicit     INTEGER,
    track_number INTEGER,
    disc_number  INTEGER,
    album_id     TEXT    REFERENCES albums (album_id) ON DELETE CASCADE
);
`

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

type AlbumJSON struct {
	AlbumType            string            `json:"album_type"`
	TotalTracks          int               `json:"total_tracks"`
	ExternalUrls         map[string]string `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []ImagesData      `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Type                 string            `json:"type"`
	Uri                  string            `json:"uri"`
	Artists              []ArtistsData     `json:"artists"`
	Tracks               TracksWrapper     `json:"tracks"`
	Copyrights           []CopyrightsData  `json:"copyrights"`
	ExternalIds          map[string]string `json:"external_ids"`
	Genres               []string          `json:"genres"`
}

type ImagesData struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ArtistsData struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

type TracksWrapper struct {
	Href     string      `json:"href"`
	Limit    int         `json:"limit"`
	Next     *string     `json:"next"`
	Offset   int         `json:"offset"`
	Previous *string     `json:"previous"`
	Total    int         `json:"total"`
	Items    []TrackData `json:"items"`
}

type TrackData struct {
	Artists      []ArtistsData     `json:"artists"`
	DiscNumber   int               `json:"disc_number"`
	DurationMs   int               `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	TrackNumber  int               `json:"track_number"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
	IsLocal      bool              `json:"is_local"`
}

type CopyrightsData struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
