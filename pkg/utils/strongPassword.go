package utils

import "regexp"

func IsStrongPassword(password string) bool {
	var (
		uppercasePattern   = regexp.MustCompile(`[A-Z]`)
		lowercasePattern   = regexp.MustCompile(`[a-z]`)
		numberPattern      = regexp.MustCompile(`[0-9]`)
		specialCharPattern = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\|,.<>\/\?\~]`)
	)
	if len(password) < 8 {
		return false
	}
	if !uppercasePattern.MatchString(password) {
		return false
	}
	if !lowercasePattern.MatchString(password) {
		return false
	}
	if !numberPattern.MatchString(password) {
		return false
	}
	if !specialCharPattern.MatchString(password) {
		return false
	}
	return true
}
