package helpers

import (
	"net/http"
	"regexp"
)

func IsValidUsername(username string) bool {
	match, err := regexp.MatchString("^[a-z0-9._-]{5,30}$", username)
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
