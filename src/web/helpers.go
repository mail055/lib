package web

import (
	"encoding/json"
	"fmt"
)

func formatAlbumDuration(durationMS int) string {
	var duration string

	hours, minutes, seconds := parseDuration(durationMS)

	// Create string in format H hr M min or M min S sec
	if hours >= 1 {
		duration = fmt.Sprintf("%d hr %d min", hours, minutes)
	} else {
		duration = fmt.Sprintf("%d min %d sec", minutes, seconds)
	}
	return duration
}

func formatTrackDuration(durationMS int) string {
	var duration string

	hours, minutes, seconds := parseDuration(durationMS)

	// Create string in format HH:MM:SS or MM:SS
	if hours >= 1 {
		duration = fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	} else {
		duration = fmt.Sprintf("%d:%02d", minutes, seconds)
	}
	return duration
}

func parseDuration(durationMS int) (int, int, int) {
	// Declare hours, minutes, and seconds
	var hours, minutes, seconds int
	minutes = (durationMS / 1000) / 60
	seconds = (durationMS / 1000) % 60

	// Evaluate hours if minutes greater than or equal to 60
	if minutes >= 60 {
		hours = minutes / 60
		minutes = minutes % 60
	}

	return hours, minutes, seconds
}

func parseAlbumJSON(input string) (*AlbumJSON, error) {
	var album AlbumJSON

	if err := json.Unmarshal([]byte(input), &album); err != nil {
		return nil, err
	}

	return &album, nil
}
