package tui

import (
	"log"
)

func updateUI() {
	sidebar.refreshAlbumList()
}

func (sb *Sidebar) refreshAlbumList() {
	albumIDs = nil
	sb.albumList.Clear()

	// Requests for the list of Albums
	albums, err := AlbumsQuery()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Adds every album to the albumList
	for _, album := range albums {
		secondaryText := album.Artists + "\n" + album.ID
		albumIDs = append(albumIDs, album.ID)
		sb.albumList.AddItem(album.Title, secondaryText, 0, nil)
	}
}

func (ct *Content) refreshSongList(albumID string) {
	ct.trackList.Clear()

	// Requests for the list of tracks
	tracks, err := TracksQuery(albumID)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Adds every track to the trackList
	for _, track := range tracks {
		ct.trackList.AddItem(track.Title, track.Artists, 0, nil)
	}
}
