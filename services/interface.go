package services

type Services interface {
	IsUniqueEmail(email string) bool
	IsNumberValid(number string) bool
}
