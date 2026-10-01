package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func initUI() {
	app = tview.NewApplication().
		EnableMouse(true)
	page = tview.NewFlex().
		SetDirection(tview.FlexColumn)

	sidebar.initSidebar()
	content.initContent()
}

func (sb *Sidebar) initSidebar() {
	sb.initSearchBar()
	sb.initAlbumList()

	sb.flex = tview.NewFlex().
		SetDirection(tview.FlexRow)
	sb.flex.
		AddItem(sb.searchBar, 3, 0, true).
		AddItem(sb.albumList, 0, 1, false)

	page.AddItem(sb.flex, 0, 30, true)
}

func (sb *Sidebar) initSearchBar() {
	sb.searchBar = tview.NewInputField()

	// Border Attributes
	sb.searchBar.SetTitle("Search Bar").
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	// Label, Field, and Placeholder Styling
	sb.searchBar.SetLabel(" > ").
		SetLabelColor(tcell.ColorDefault).
		SetFieldStyle(tcell.StyleDefault).
		SetPlaceholder("Enter Album / Artist Name").
		SetPlaceholderStyle(
			tcell.StyleDefault.
				Foreground(tcell.ColorDarkGray).
				Background(tcell.ColorDefault))

	// Update Album List when typed in
	sb.searchBar.SetChangedFunc(func(text string) {
		searchQuery = sb.searchBar.GetText()
		sb.refreshAlbumList()
	})
}

func (sb *Sidebar) initAlbumList() {
	sb.albumList = tview.NewList()

	// Border Attributes
	title := fmt.Sprintf("Album List (%d)", sb.albumList.GetItemCount())
	sb.albumList.SetTitle(title).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	// Main Text, Secondary Text, and Selected Text Styling
	sb.albumList.SetMainTextStyle(tcell.StyleDefault).
		SetSecondaryTextColor(tcell.ColorDarkGray).
		SetSelectedStyle(tcell.StyleDefault.
			Foreground(tcell.ColorSpringGreen).
			Background(tcell.ColorDefault))

	// Update Track List and Album Info when scrolling
	sb.albumList.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		currentAlbumID := albumIDs[index]
		content.refreshAlbumInfo(currentAlbumID)
		content.refreshTrackList(currentAlbumID)
	})
}

func (ct *Content) initContent() {
	ct.albumInfo.initAlbumInfo()
	ct.initTrackList()

	ct.flex = tview.NewFlex().
		SetDirection(tview.FlexRow)
	ct.flex.
		AddItem(ct.albumInfo.flex, 0, 40, false).
		AddItem(ct.trackList, 0, 60, true)

	page.AddItem(ct.flex, 0, 70, true)
}

func (ai *AlbumInfo) initAlbumInfo() {
	ai.initAlbumCover()
	ai.initAlbumText()

	ai.flex = tview.NewFlex().
		SetDirection(tview.FlexColumn)
	ai.flex.SetTitle("Album Info").
		SetBorder(true)
	ai.flex.
		AddItem(ai.cover, 0, 40, false).
		AddItem(ai.text, 0, 60, false)
}

func (ai *AlbumInfo) initAlbumCover() {
	ai.cover = tview.NewTextView()
	ai.cover.SetBorderPadding(1, 1, 1, 1)
	ai.cover.SetTextAlign(tview.AlignCenter)
}

func (ai *AlbumInfo) initAlbumText() {
	ai.text = tview.NewTextView()
	ai.text.SetBorderPadding(1, 1, 1, 1)
}

func (ct *Content) initTrackList() {
	ct.trackList = tview.NewList()

	// Border Attributes
	ct.trackList.SetTitle("Songs").
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	// Main Text, Secondary Text, and Selected Text Styling
	ct.trackList.SetMainTextStyle(tcell.StyleDefault).
		SetSecondaryTextColor(tcell.ColorDarkGray).
		SetSelectedStyle(tcell.StyleDefault.
			Foreground(tcell.ColorSpringGreen).
			Background(tcell.ColorDefault))
}
