package tui

import (
	"database/sql"
	"log"

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

func AlbumsQuery() ([]Album, error) {
	var albums []Album

	// Creates a query that gets every Album ID, Album Name, and Artist Name(s)
	query := `
    SELECT 
        al.album_id,
        al.album_name,
        GROUP_CONCAT(ar.artist_name, ', ') AS artists
    FROM albums AS al
    JOIN album_artists AS aa ON al.album_id = aa.album_id
    JOIN artists AS ar ON aa.artist_id = ar.artist_id
    GROUP BY al.album_id, al.album_name
    ORDER BY al.album_name;` // Likely allow for user adjustment in future (e.g. ID, Artist, Release Date, etc.)

	// Performs the query on the database
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Receives the query response and assigns variables to them
	for rows.Next() {
		var nextAlbum Album
		if err := rows.Scan(&nextAlbum.ID, &nextAlbum.Title, &nextAlbum.Artists); err != nil {
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

func TracksQuery(albumID string) ([]Track, error) {
	var tracks []Track

	// Creates a query that gets every Track ID, Track Name, and Artist Name(s)
	query := `
	SELECT
		tr.track_id,
		tr.track_name,
		GROUP_CONCAT(ar.artist_name, ', ') AS artists
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
		var nextTrack Track
		if err := rows.Scan(&nextTrack.ID, &nextTrack.Title, &nextTrack.Artists); err != nil {
			return nil, err
		}
		tracks = append(tracks, nextTrack)
	}

	// Checks for any hidden errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Returns the list of tracks if all error checking passes
	return tracks, nil
}

func AlbumInfoQuery(albumID string) (*AlbumInfo, error) {
	var info AlbumInfo

	// Creates a query that gets all the album info
	query := `
    SELECT 
        al.album_id,
        al.album_name,
        GROUP_CONCAT(ar.artist_name, ', ') AS artists,
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
			&info.TotalTracks, &info.DurationMS, &info.URL,
			&info.ImageURL, &info.ReleaseDate); err != nil {
			return nil, err
		}
	}

	// Checks for any hidden errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Returns the list of tracks if all error checking passes
	return &info, nil
}
