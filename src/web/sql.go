package web

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

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

func AlbumsQuery(searchQuery string) ([]AlbumStruct, error) {
	var albums []AlbumStruct
	var whereClause string
	var args []any

	// Creates a WHERE clause if the searchQuery is not empty
	if searchQuery != "" {
		whereClause = "WHERE al.album_name LIKE ? OR ar.artist_name LIKE ? OR al.album_id LIKE ?"
		searchTerm := "%" + searchQuery + "%"
		args = []any{searchTerm, searchTerm, searchTerm}
	}

	// Creates a query that gets every Album ID, Album Name, and Artist Name(s)
	query := fmt.Sprintf(`
    SELECT 
        al.album_id,
        al.album_name,
        GROUP_CONCAT(ar.artist_name, ', ') AS artists,
        al.album_image
    FROM albums AS al
    JOIN album_artists AS aa ON al.album_id = aa.album_id
    JOIN artists AS ar ON aa.artist_id = ar.artist_id
    %s
    GROUP BY al.album_id, al.album_name
    ORDER BY al.album_name;`, whereClause) // Likely allow for user adjustment in future (e.g. ID, Artist, Release Date, etc.)

	// Performs the query on the database
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Receives the query response and assigns variables to them
	for rows.Next() {
		var nextAlbum AlbumStruct
		if err := rows.Scan(&nextAlbum.ID, &nextAlbum.Title, &nextAlbum.Artists, &nextAlbum.ImageURL); err != nil {
			return nil, err
		}
		albums = append(albums, nextAlbum)
	}

	// Checks for any hidden errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Returns the list of albums if all error checking passes
	return albums, nil
}

func TracksQuery(albumID string) ([]TrackStruct, error) {
	var tracks []TrackStruct

	// Creates a query that gets every Track ID, Track Name, and Artist Name(s)
	query := `
	SELECT
		tr.track_id,
		tr.track_number,
		tr.track_name,
		GROUP_CONCAT(ar.artist_name, ', ') AS artists,
		tr.duration_ms
	FROM tracks AS tr
	JOIN artist_tracks AS atr ON tr.track_id = atr.track_id
	JOIN artists AS ar ON atr.artist_id = ar.artist_id
	WHERE tr.album_id = ?
	GROUP BY tr.track_id, tr.track_name
	ORDER BY tr.track_number;` // Also allow for user adjustment

	// Performs the query on the database
	rows, err := db.Query(query, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Receives the query response and assigns variables to them
	for rows.Next() {
		var nextTrack TrackStruct
		if err := rows.Scan(&nextTrack.ID,
			&nextTrack.TrackNumber, &nextTrack.Title,
			&nextTrack.Artists, &nextTrack.DurationMS); err != nil {
			return nil, err
		}
		nextTrack.Duration = formatTrackDuration(nextTrack.DurationMS)
		tracks = append(tracks, nextTrack)
	}

	// Checks for any hidden errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Returns the list of tracks if all error checking passes
	return tracks, nil
}

func AlbumInfoQuery(albumID string) (*AlbumInfoStruct, error) {
	var info AlbumInfoStruct

	if albumID == "" {
		return nil, nil
	}

	// Creates a query that gets all the album info
	query := `
    SELECT 
        al.album_id,
        al.album_name,
        GROUP_CONCAT(ar.artist_name, ', ') AS artists,
        al.album_type,
    	al.total_tracks,
    	al.duration_ms,
        al.album_url,
        al.album_image,
    	al.release_date
    FROM albums AS al
    JOIN album_artists AS aa ON al.album_id = aa.album_id
    JOIN artists AS ar ON aa.artist_id = ar.artist_id
    WHERE al.album_id = ?
    GROUP BY al.album_id, al.album_name;`

	// Performs the query on the database
	rows, err := db.Query(query, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Receives the query response and assigns variables to them
	for rows.Next() {
		if err := rows.Scan(&info.ID, &info.Title, &info.Artists,
			&info.AlbumType, &info.TotalTracks, &info.DurationMS,
			&info.URL, &info.ImageURL, &info.ReleaseDate); err != nil {
			return nil, err
		}
		info.AlbumType = strings.Title(info.AlbumType)
		info.Duration = formatAlbumDuration(info.DurationMS)
	}

	// Checks for any hidden errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Returns the list of tracks if all error checking passes
	return &info, nil
}
