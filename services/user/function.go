package user

import (
	"regexp"
	"strconv"
)

func IsUniqueEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)$")
	return emailRegex.MatchString(email)
}

func IsNumberValid(number string) bool {
	var phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phoneRegex.MatchString(number) && len(number) == 10
}

func IsValidId(number string) bool {
	if _, err := strconv.Atoi(number); err != nil {
		return false
	}
	return true
}
