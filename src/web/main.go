package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Main() {
	// Manage Database Initialisation
	initDB()
	defer db.Close()

	var err error

	tmpls, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Initialise Multiplexer
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("GET /{$}", handleIndex)
	mux.HandleFunc("GET /{id}", handleIndex)
	mux.HandleFunc("POST /import", handleImport)
	mux.HandleFunc("DELETE /album/{id}", handleDelete)

	// Runs the server
	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Return error if an invalid address is found
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get variables from url
	searchQuery := r.URL.Query().Get("q")
	selectedAlbum := r.URL.Query().Get("id")

	// Get Albums
	albums, err := AlbumsQuery(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get Album Info
	albumInfo, err := AlbumInfoQuery(selectedAlbum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get Tracks
	tracks, err := TracksQuery(selectedAlbum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Compile Data
	data := PageData{
		Albums:      albums,
		SearchQuery: searchQuery,
		AlbumInfo:   albumInfo,
		Tracks:      tracks,
	}

	// Handle Get Requests
	if r.Header.Get("HX-Request") == "true" {
		switch r.Header.Get("HX-Target") {
		case "album-results":
			err = tmpls.ExecuteTemplate(w, "album-results", data)
		case "content":
			err = tmpls.ExecuteTemplate(w, "content", data)
		default:
			err = tmpls.ExecuteTemplate(w, "index.html", data)
		}
	} else {
		err = tmpls.ExecuteTemplate(w, "index.html", data)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleImport(w http.ResponseWriter, r *http.Request) {
	var msg string
	input := r.FormValue("input")

	album, err := parseAlbumJSON(input)
	if err != nil {
		msg = fmt.Sprintf("Error parsing album JSON: %v", err)
	} else {
		// Check if album already exists in the database
		exists, err := albumExists(album.ID)
		if err != nil {
			msg = fmt.Sprintf("Database error: %v", err)
		} else if exists {
			msg = fmt.Sprintf(`Album "%s" is already in your library`, album.Name)
		} else if err = album.insertAlbumJSON(); err != nil {
			msg = fmt.Sprintf("Error inserting album: %v", err)
		} else {
			msg = fmt.Sprintf(`Album "%s" Successfully Imported`, album.Name)
		}
	}

	// Show notification
	triggers := map[string]any{
		"showNotification": msg,
		"refreshAlbums":    true,
	}

	if jsonBytes, err := json.Marshal(triggers); err == nil {
		w.Header().Set("HX-Trigger", string(jsonBytes))
	}

	w.WriteHeader(http.StatusOK)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	var msg string
	id := r.PathValue("id")

	// Get the album info
	info, err := AlbumInfoQuery(id)
	if err != nil {
		msg = fmt.Sprintf("Error getting album info: %v", err)
	}

	// Attempt to delete the album
	err = deleteAlbum(id)
	if err != nil {
		msg = fmt.Sprintf("Error deleting album: %v", err)
	}

	if info != nil {
		msg = fmt.Sprintf("Album '%s' Successfully Deleted", info.Title)
	}

	// Trigger notification
	triggers := map[string]any{
		"showNotification": msg,
		"refreshAlbums":    true,
	}

	if jsonBytes, err := json.Marshal(triggers); err == nil {
		w.Header().Set("HX-Trigger", string(jsonBytes))
	}

	// Clear out content
	emptyData := PageData{}
	w.Header().Set("Content-Type", "text/html")
	if err := tmpls.ExecuteTemplate(w, "content", emptyData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
