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

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleIndex)

	// Runs the server
	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Return error if an invalid address is found
	if r.URL.Path != "/" && r.URL.Path != "/album" {
		http.NotFound(w, r)
		return
	}

	searchQuery := r.URL.Query().Get("id")
	selectedAlbum := r.URL.Query().Get("id")

	albums, err := AlbumsQuery("")
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

	err = tmpls.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
