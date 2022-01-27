package services

import (
	"fmt"
	"regexp"
	"strconv"
)

func IsUniqueEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func IsNumberValid(number string) bool {
	if _, err := strconv.Atoi(number); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
