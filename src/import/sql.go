package _import

import (
	"database/sql"
	"log"
	"maps"
	"slices"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error

	// Establishes the databases connection
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Tests the database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (a *AlbumJSON) insertAlbumJSON() error {
	a.insertAlbum()
	a.insertArtist(a.ID)
	a.Tracks.insertTracks(a.ID)
	return nil
}

func (a *AlbumJSON) insertAlbum() {
	id := a.ID
	name := a.Name
	albumType := a.AlbumType
	totalTracks := a.TotalTracks
	var duration int
	for _, track := range a.Tracks.Items {
		duration += track.DurationMs
	}
	url := slices.Collect(maps.Values(a.ExternalUrls))[0]
	image := a.Images[0].Url
	releaseDate := a.ReleaseDate

	albumQuery := `
		INSERT INTO albums (
		    album_id, album_name, album_type, total_tracks,
            duration_ms, album_url, album_image, release_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(albumQuery,
		id, name, albumType, totalTracks,
		duration, url, image, releaseDate,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *AlbumJSON) insertArtist(albumID string) {
	for _, artist := range a.Artists {
		artistID := artist.ID
		artistName := artist.Name
		artistImage := "" // Might remove
		followers := 0    // Also might remove
		artistURL := slices.Collect(maps.Values(artist.ExternalUrls))[0]

		artistQuery := `
			INSERT INTO artists (
			    artist_id, artist_name, artist_image,
			    followers, artist_url)
			VALUES (?, ?, ?, ?, ?)`

		_, err := db.Exec(artistQuery,
			artistID, artistName, artistImage,
			followers, artistURL)
		if err != nil {
			log.Fatal(err)
		}

		albumArtistQuery := `
			INSERT INTO album_artists (
        		album_id, artist_id)
    		VALUES (?, ?)`

		_, err = db.Exec(albumArtistQuery,
			albumID, artistID,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, track := range a.Tracks.Items {
		for _, artist := range track.Artists {
			artistID := artist.ID
			artistName := artist.Name
			artistImage := "" // Might remove
			followers := 0    // Also might remove
			artistURL := slices.Collect(maps.Values(artist.ExternalUrls))[0]

			artistQuery := `
			INSERT INTO artists (
			    artist_id, artist_name, artist_image,
			    followers, artist_url)
			VALUES (?, ?, ?, ?, ?)`

			_, err := db.Exec(artistQuery,
				artistID, artistName, artistImage,
				followers, artistURL)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (t *TracksWrapper) insertTracks(albumID string) {
	for _, track := range t.Items {
		trackID := track.ID
		name := track.Name
		url := slices.Collect(maps.Values(track.ExternalUrls))[0]
		durationMS := track.DurationMs
		isrc := ""
		explicit := track.Explicit
		trackNumber := track.TrackNumber
		discNumber := track.DiscNumber

		trackQuery := `
			INSERT INTO tracks (
                track_id,  track_name, track_url, duration_ms, isrc, 
			    explicit, track_number, disc_number, album_id)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

		_, err := db.Exec(trackQuery,
			trackID, name, url, durationMS, isrc,
			explicit, trackNumber, discNumber, albumID)
		if err != nil {
			log.Fatal(err)
		}

		for _, artist := range track.Artists {
			artistID := artist.ID

			artistTrackQuery := `
			INSERT INTO artist_tracks (
                artist_id, track_id)
            VALUES (?, ?)`

			_, err = db.Exec(artistTrackQuery,
				artistID, trackID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
