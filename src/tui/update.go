package tui

import (
	"log"
)

func updateUI() {
	sidebar.refreshAlbumList()
}

func (sb *Sidebar) refreshAlbumList() {
	sb.albumList.Clear()

	// Requests for the list of Albums
	albums, err := AlbumsQuery()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Adds every album to the albumList
	for _, album := range albums {
		secondaryText := album.Artists + "\n" + album.ID
		sb.albumList.AddItem(album.Title, secondaryText, 0, nil)
	}
}
