package web

import "fmt"

func parseDuration(durationMS int) string {
	var hours, minutes, seconds int
	var duration string
	minutes = (durationMS / 1000) / 60
	seconds = (durationMS / 1000) % 60

	if minutes >= 60 {
		hours = minutes / 60
		minutes = minutes % 60
	}

	if hours >= 1 {
		duration = fmt.Sprintf("%d hr %d min", hours, minutes)
	} else {
		duration = fmt.Sprintf("%d min %d sec", minutes, seconds)
	}
	return duration
}
