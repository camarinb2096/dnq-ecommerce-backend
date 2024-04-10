package utils

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}
