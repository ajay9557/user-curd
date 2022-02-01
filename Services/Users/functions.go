package Users

import "regexp"

func validateEmail(email string) bool {
	var emreg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return emreg.MatchString(email)

}

func validatePhone(phone string) bool {
	var phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)

	return phoneRegex.MatchString(phone) && len(phone) == 10
}
