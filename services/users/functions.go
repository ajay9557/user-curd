package users

import (
	"regexp"

)

func ValidId(id int) bool {
	if id < 0 {
		return false
	}

	return true
}

func ValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func ValidPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)

	return phoneRegex.MatchString(phone) && len(phone) == 10
}

