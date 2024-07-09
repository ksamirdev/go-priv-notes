package helpers

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

var (
	ErrInvalidUsername = "Username is invalid. Please make sure it's in between 5-30 characters with only english alphabets, numeric value, and special characters ('_', '-')!"
	ErrInvalidPin      = "Pin should be numeric with 6 values"
)

func IsValidUsername(username string) bool {
	match, err := regexp.MatchString("^[a-z0-9_-]{5,30}$", username)
	if err != nil {
		return false
	}

	return match
}

func IsValidPin(pin string) bool {
	match, err := regexp.MatchString("^\\d{6}$", pin)
	if err != nil {
		return false
	}

	return match
}

func IsURLEncodedFormValid(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/x-www-form-urlencoded"
}

func ReadableTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(diff.Seconds()))
	case diff < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
	case diff < time.Hour*24:
		return fmt.Sprintf("%d hours ago", int(diff.Hours()))
	case diff < time.Hour*24*7:
		return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
	case diff < time.Hour*24*30:
		return fmt.Sprintf("%d weeks ago", int(diff.Hours()/(24*7)))
	case diff < time.Hour*24*365:
		return fmt.Sprintf("%d months ago", int(diff.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d years ago", int(diff.Hours()/(24*365)))
	}
}
