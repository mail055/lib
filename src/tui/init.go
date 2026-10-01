package tui

import "github.com/rivo/tview"

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
	sb.searchBar.SetTitle("Search Bar").
		SetBorder(true)

	sb.searchBar.SetChangedFunc(func(text string) {
		searchQuery = "%" + sb.searchBar.GetText() + "%"
		sb.refreshAlbumList()
	})
}

func (sb *Sidebar) initAlbumList() {
	sb.albumList = tview.NewList()
	sb.albumList.SetTitle("Album List").
		SetBorder(true)

	sb.albumList.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		currentAlbumID := albumIDs[index]
		content.refreshTrackList(currentAlbumID)
		content.refreshAlbumInfo(currentAlbumID)
	})
}

func (ct *Content) initContent() {
	ct.initAlbumInfo()
	ct.initTrackList()

	ct.flex = tview.NewFlex().
		SetDirection(tview.FlexRow)
	ct.flex.
		AddItem(ct.albumInfo, 0, 40, false).
		AddItem(ct.trackList, 0, 60, true)

	page.AddItem(ct.flex, 0, 70, true)
}

func (ct *Content) initAlbumInfo() {
	ct.albumInfo = tview.NewTextView()
	ct.albumInfo.SetTitle("Album Info").
		SetBorder(true)
}

func (ct *Content) initTrackList() {
	ct.trackList = tview.NewList()
	ct.trackList.SetTitle("Songs").
		SetBorder(true)
}
