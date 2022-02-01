package users

import (
	"regexp"
)

func emailValidation(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func idCheck(id int) bool {
	return id > 0
}
