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
