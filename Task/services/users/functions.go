package service

import (
	"regexp"

)

// validateId utility to validate the id
func validateId(id interface{}) bool {

	// check if id is of type int and is a positive number
	idInt, ok := id.(int)
	if !ok {
		return false
	}
	if idInt < 0 {
		return false
	}
	return true
}

// validateEmail utility to validate email ids
func validateEmail(email string) bool {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// validatePhone utility to validate phone number
func validatePhone(phone string) bool {

	phoneRegex := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phoneRegex.MatchString(phone) && len(phone) <= 10
}
