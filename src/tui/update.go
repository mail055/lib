package tui

import (
	"fmt"
	"log"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		sb.albumList.AddItem(album.Title,
			" "+album.Artists,
			0,
			nil)
	}

	// Update Album List Text
	title := fmt.Sprintf("Album List (%d)", sb.albumList.GetItemCount())
	sb.albumList.SetTitle(title)
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
		durationSeconds := track.DurationMS / 1000
		duration := fmt.Sprintf("%d:%02d", durationSeconds/60, durationSeconds%60)
		mainText := fmt.Sprintf("[%d] - %s | %s", track.TrackNumber, track.Title, duration)
		ct.trackList.AddItem(mainText,
			" "+track.Artists,
			0,
			nil)
	}
}

func (ct *Content) refreshAlbumInfo(albumID string) {
	ct.albumInfo.text.Clear()
	ct.albumInfo.cover.Clear()

	// Requests for the Album Info
	info, err := AlbumInfoQuery(albumID)
	if err != nil {
		log.Fatal(err)
	}

	// Declares variables
	albumType := cases.Title(language.English).String(info.AlbumType)
	title := info.Title
	url := info.URL
	artists := info.Artists
	releaseDate := info.ReleaseDate
	totalTracks := fmt.Sprintf("%d songs", info.TotalTracks)
	durationSeconds := info.DurationMS / 1000
	duration := fmt.Sprintf("%d min %02d sec", durationSeconds/60, durationSeconds%60)
	ID := info.ID
	image := info.ImageURL

	coverText := fmt.Sprintf("\n\n\nAlbum Cover - %s\n\n\nALBUM COVER WOULD GO HERE", image)
	infoText := fmt.Sprintf(`%s
%s • %s
%s • %s • %s, %s
%s`,
		albumType, title, url, artists,
		releaseDate, totalTracks, duration, ID)

	// Applies the cover and info text there respective UI elements
	ct.albumInfo.cover.SetText(coverText)
	ct.albumInfo.text.SetText(infoText)
}
