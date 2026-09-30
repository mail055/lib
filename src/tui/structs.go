package tui

import "github.com/rivo/tview"
import (
	"database/sql"

	"github.com/rivo/tview"
)

var app *tview.Application
var page *tview.Flex

var sidebar Sidebar
var content Content
var db *sql.DB

type Sidebar struct {
	flex      *tview.Flex
	searchBar *tview.InputField
	albumList *tview.List
}

type Content struct {
	flex      *tview.Flex
	albumInfo *tview.TextView
	songList  *tview.List
type Album struct {
	ID      string
	Title   string
	Artists string
}
}
