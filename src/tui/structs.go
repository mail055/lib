package tui

import "github.com/rivo/tview"

var app *tview.Application
var page *tview.Flex

var sidebar Sidebar
var content Content

type Sidebar struct {
	flex      *tview.Flex
	searchBar *tview.InputField
	albumList *tview.List
}

type Content struct {
	flex      *tview.Flex
	albumInfo *tview.TextView
	songList  *tview.List
}
