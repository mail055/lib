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

type Sidebar struct {
	flex      *tview.Flex
	searchBar *tview.InputField
	albumList *tview.List
}

type Content struct {
	flex      *tview.Flex
	albumInfo *tview.TextView
	trackList *tview.List
}

type Album struct {
	ID      string
	Title   string
	Artists string
}

type Track struct {
	ID      string
	Title   string
	Artists string
}
