package helpers

import (
	"net/http"
	"regexp"
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
