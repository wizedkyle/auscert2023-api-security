package utils

import "time"

// GenerateExpiration
// Generates a date based on the supplied duration (which is the number of days).
func GenerateExpiration(duration int) time.Time {
	if duration == 0 { // No expiry
		now := time.Now()
		expiry := time.Date(2099, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
		return expiry
	} else {
		expiry := time.Now().AddDate(0, 0, duration)
		return expiry
	}
}
