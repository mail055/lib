# Lib

## Description
A Fast, Self-Hosted, Spotify-Like Library Interface written in Go

## Set Up
Navigate to the directory where you would like the application to be stored and perform:
```bash
# Clone the repository
git clone 'github.com/mail055/lib'
# Ensure you have a Go Interpreter installed on your device
go run src/main.go

# You may also run:
go run src/main.go -terminal # For the Terminal UI
go run src/main.go -import # For the command line import
```

At this point, the application should have created data.db, however it will contain no data

## Importing Albums
Begin by pressing the + icon found near the bottom left corner.

This will bring up a modal with steps on how import an album, though I will provide a detailed explanation here.

1. Navigate to the Album you desire to import into the application. Right click it > Share > Copy Link
2. Go to [this](https://www.unchainedmusic.io/tools/spotify-uri-finder/) website and paste the link into the text area.

    The website will return 3 values: the URI, the ID, and the link. Copy the ID
3. Go to [this](https://developer.spotify.com/documentation/web-api/reference/get-an-album) website (ensure you are logged into Spotify) and paste the ID into its respective field.
4. Press the 'Try it' button. This will update the Response Sample below
5. Copy the **entire** response sample

    It should look something like this:
```go
{
    "album_type": "album",
    "total_tracks": 18,
    // ...
    // More Information
    // ...
    "genres": []
}
``` 
6. Return to the application and paste the response sample into the text area
7. Press 'Add'

If all steps were done correctly, the album should be automatically added to the list. 

If unsuccessful:
- Ensure you copy the entire response sample, squiggly brackets inclusive and try again
- Make sure the album is not already imported
- Check the notification you get bottom right to diagnose the issue.
- Refresh the page (Ctrl + Shift + R) to flush cache
- Restart the application
