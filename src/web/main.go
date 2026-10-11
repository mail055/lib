package web

import (
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

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("GET /{$}", handleIndex)
	mux.HandleFunc("GET /{id}", handleIndex)

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

	searchQuery := r.URL.Query().Get("q")
	selectedAlbum := r.URL.Query().Get("id")

	albums, err := AlbumsQuery(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	albumInfo, err := AlbumInfoQuery(selectedAlbum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
