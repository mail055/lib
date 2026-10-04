package _import

type AlbumJSON struct {
	AlbumType            string            `json:"album_type"`
	TotalTracks          int               `json:"total_tracks"`
	ExternalUrls         map[string]string `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []ImagesData      `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Type                 string            `json:"type"`
	Uri                  string            `json:"uri"`
	Artists              []ArtistsData     `json:"artists"`
	Tracks               TracksWrapper     `json:"tracks"`
	Copyrights           []CopyrightsData  `json:"copyrights"`
	ExternalIds          map[string]string `json:"external_ids"`
	Genres               []string          `json:"genres"`
}

type ImagesData struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ArtistsData struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

type TracksWrapper struct {
	Href     string      `json:"href"`
	Limit    int         `json:"limit"`
	Next     *string     `json:"next"`
	Offset   int         `json:"offset"`
	Previous *string     `json:"previous"`
	Total    int         `json:"total"`
	Items    []TrackData `json:"items"`
}

type TrackData struct {
	Artists      []ArtistsData     `json:"artists"`
	DiscNumber   int               `json:"disc_number"`
	DurationMs   int               `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	TrackNumber  int               `json:"track_number"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
	IsLocal      bool              `json:"is_local"`
}

type CopyrightsData struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
