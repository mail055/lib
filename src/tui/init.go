package tui

import "github.com/rivo/tview"

func initUI() {
	app = tview.NewApplication()
	page = tview.NewFlex().
		SetDirection(tview.FlexColumn)

	sidebar.initSidebar()
	content.initContent()
}

func (sb *Sidebar) initSidebar() {
	sb.searchBar = tview.NewInputField()
	sb.searchBar.SetTitle("Search Bar").
		SetBorder(true)

	sb.albumList = tview.NewList()
	sb.albumList.SetTitle("Album List").
		SetBorder(true)

	sb.flex = tview.NewFlex().
		SetDirection(tview.FlexRow)
	sb.flex.
		AddItem(sb.searchBar, 3, 0, true).
		AddItem(sb.albumList, 0, 1, false)

	page.AddItem(sb.flex, 0, 30, true)
}

func (ct *Content) initContent() {
	ct.albumInfo = tview.NewTextView()
	ct.albumInfo.SetTitle("Album Info").
		SetBorder(true)

	ct.songList = tview.NewList()
	ct.songList.SetTitle("Songs").
		SetBorder(true)

	ct.flex = tview.NewFlex().
		SetDirection(tview.FlexRow)
	ct.flex.
		AddItem(ct.albumInfo, 0, 40, false).
		AddItem(ct.songList, 0, 60, true)

	page.AddItem(ct.flex, 0, 70, true)
}
