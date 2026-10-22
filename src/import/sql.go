package _import

import (
	"database/sql"
	"log"
	"maps"
	"slices"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() {
	var err error

	// Establishes the databases connection
	db, err = sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Tests the database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (a *AlbumJSON) insertAlbumJSON() error {
	// Ensures insertions do not fail during import
	// If it does, rolls back database
	var err error
	tx, err = db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err = a.insertAlbum(); err != nil {
		return err
	} else if err = a.insertArtists(a.ID); err != nil {
		return err
	} else if err = a.Tracks.insertTracks(a.ID); err != nil {
		return err
	}
	return tx.Commit()
}

func (a *AlbumJSON) insertAlbum() error {
	// Variable Declarations
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

	// Create Query
	albumQuery := `
		INSERT INTO albums (
		    album_id, album_name, album_type, total_tracks,
            duration_ms, album_url, album_image, release_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute Query
	_, err := tx.Exec(albumQuery,
		id, name, albumType, totalTracks,
		duration, url, image, releaseDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *AlbumJSON) insertArtists(albumID string) error {
	for _, artist := range a.Artists {
		// Variable Declarations
		artistID := artist.ID
		artistName := artist.Name
		artistImage := "" // Might remove
		followers := 0    // Also might remove
		artistURL := slices.Collect(maps.Values(artist.ExternalUrls))[0]

		// Create Query
		artistQuery := `
			INSERT INTO artists (
			    artist_id, artist_name, artist_image,
			    followers, artist_url)
			VALUES (?, ?, ?, ?, ?)`

		// Execute Query
		_, err := tx.Exec(artistQuery,
			artistID, artistName, artistImage,
			followers, artistURL,
		)
		if err != nil {
			return err
		}

		// Create Query
		albumArtistQuery := `
			INSERT INTO album_artists (
        		album_id, artist_id)
    		VALUES (?, ?)`

		// Execute Query
		_, err = tx.Exec(albumArtistQuery,
			albumID, artistID,
		)
		if err != nil {
			return err
		}
	}

	for _, track := range a.Tracks.Items {
		for _, artist := range track.Artists {
			// Variable Declarations
			artistID := artist.ID
			artistName := artist.Name
			artistImage := "" // Might remove
			followers := 0    // Also might remove
			artistURL := slices.Collect(maps.Values(artist.ExternalUrls))[0]

			// Create Query
			artistQuery := `
			INSERT INTO artists (
			    artist_id, artist_name, artist_image,
			    followers, artist_url)
			VALUES (?, ?, ?, ?, ?)`

			// Execute Query
			_, err := tx.Exec(artistQuery,
				artistID, artistName, artistImage,
				followers, artistURL,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TracksWrapper) insertTracks(albumID string) error {
	for _, track := range t.Items {
		// Variable Declarations
		trackID := track.ID
		name := track.Name
		url := slices.Collect(maps.Values(track.ExternalUrls))[0]
		durationMS := track.DurationMs
		isrc := ""
		explicit := track.Explicit
		trackNumber := track.TrackNumber
		discNumber := track.DiscNumber

		// Create Query
		trackQuery := `
			INSERT INTO tracks (
                track_id,  track_name, track_url, duration_ms, isrc, 
			    explicit, track_number, disc_number, album_id)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

		// Execute Query
		_, err := tx.Exec(trackQuery,
			trackID, name, url, durationMS, isrc,
			explicit, trackNumber, discNumber, albumID)
		if err != nil {
			return err
		}

		for _, artist := range track.Artists {
			artistID := artist.ID

			// Create Query
			artistTrackQuery := `
			INSERT INTO artist_tracks (
                artist_id, track_id)
            VALUES (?, ?)`

			// Execute Query
			_, err = tx.Exec(artistTrackQuery,
				artistID, trackID,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
