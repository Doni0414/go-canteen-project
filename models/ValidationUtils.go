package models

import "regexp"

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	pattern := regexp.MustCompile(regex)

	return pattern.MatchString(email)
}
