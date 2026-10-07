package web

import "fmt"

func formatAlbumDuration(durationMS int) string {
	var duration string

	hours, minutes, seconds := parseDuration(durationMS)

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

	if hours >= 1 {
		duration = fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)
	} else {
		duration = fmt.Sprintf("%d:%02d", minutes, seconds)
	}
	return duration
}

func parseDuration(durationMS int) (int, int, int) {
	var hours, minutes, seconds int
	minutes = (durationMS / 1000) / 60
	seconds = (durationMS / 1000) % 60

	if minutes >= 60 {
		hours = minutes / 60
		minutes = minutes % 60
	}

	return hours, minutes, seconds
}
