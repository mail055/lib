package tui

import (
	"database/sql"

	"github.com/rivo/tview"
)

var app *tview.Application
var page *tview.Flex

var sidebar Sidebar
var content Content
var db *sql.DB
var albumIDs []string
var searchQuery = ""

type Sidebar struct {
	flex      *tview.Flex
	searchBar *tview.InputField
	albumList *tview.List
}

type Content struct {
	flex      *tview.Flex
	albumInfo AlbumInfo
	trackList *tview.List
}

type AlbumInfo struct {
	flex  *tview.Flex
	cover *tview.TextView
	text  *tview.TextView
}

type AlbumStruct struct {
	ID      string
	Title   string
	Artists string
}

type TrackStruct struct {
	ID          string
	TrackNumber int
	Title       string
	Artists     string
	DurationMS  int
}

type AlbumInfoStruct struct {
	ID          string
	Title       string
	AlbumType   string
	Artists     string
	TotalTracks int
	DurationMS  int
	URL         string
	ImageURL    string
	ReleaseDate string
}
