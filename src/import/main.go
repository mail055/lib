package _import

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func Main() {
	fmt.Printf(" == Import Flag Enabled ==\n")

	initDB()

	var albumURL string
	fmt.Println("Paste Album URL: ")
	fmt.Scanln(&albumURL)
	albumID := handleAlbumURL(albumURL)
	err := clipboard.WriteAll(albumID)
	fmt.Printf(`
1. Navigate to https://developer.spotify.com/documentation/web-api/reference/get-an-album
and paste (Ctrl + V) the Album ID copied to your clipboard (Album ID: %s)
2. Press "Try it" to enter the ID
3. Copy the enter Response Sample
It should look something like this:
{
	"album_type": "album",
	"total_tracks": 10,
	...
	...
	...
	"genres": []
}

4. Return to this page and copy the sample into the terminal with Ctrl + V and press enter

Enter Response:
`, albumID)

	album := parseAlbumJSON()
	err = album.insertAlbumJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album Added Successfully")
}

func handleAlbumURL(URL string) string {
	splitURL := strings.Split(URL, "?")
	splitURL = strings.Split(splitURL[0], "/album/")
	return splitURL[1]
}

func parseAlbumJSON() *AlbumJSON {
	var album AlbumJSON

	err := json.NewDecoder(os.Stdin).Decode(&album)
	if err != nil {
		log.Fatal(err)
	}

	return &album
}
