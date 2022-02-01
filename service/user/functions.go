package user

import (
	"fmt"
	"net/mail"
)

func CheckMail(email string) (bool,error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false,fmt.Errorf("enter valid email")
	}
	return true,nil

}
