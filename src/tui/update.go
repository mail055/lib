package tui

import (
	"fmt"
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
		log.Fatal(err)
	}

	// Adds every album to the albumList
	for _, album := range albums {
		albumIDs = append(albumIDs, album.ID)
		sb.albumList.AddItem(album.Title, album.Artists, 0, nil)
	}
}

func (ct *Content) refreshTrackList(albumID string) {
	ct.trackList.Clear()

	// Requests for the list of tracks
	tracks, err := TracksQuery(albumID)
	if err != nil {
		log.Fatal(err)
	}

	// Adds every track to the trackList
	for _, track := range tracks {
		ct.trackList.AddItem(track.Title, track.Artists, 0, nil)
	}
}

func (ct *Content) refreshAlbumInfo(albumID string) {
	ct.albumInfo.Clear()

	// Requests for the Album Info
	info, err := AlbumInfoQuery(albumID)
	if err != nil {
		log.Fatal(err)
	}

	// Declares variables
	title := info.Title
	artists := info.Artists
	totalTracks := info.TotalTracks
	durationSeconds := info.DurationMS / 1000
	duration := fmt.Sprintf("%d:%02d", durationSeconds/60, durationSeconds%60)
	url := info.URL
	image := info.ImageURL
	releaseDate := info.ReleaseDate
	ID := info.ID

	infoText := fmt.Sprintf(`%s - %s
Total Tracks - %d | Album Length - %s | Release Date - %s
Album URL   - %s
Album Cover - %s
Album ID    - %s`,
		title, artists, totalTracks, duration, releaseDate, url, image, ID)

	// Applies the info to albumInfo
	ct.albumInfo.SetText(infoText)
}
