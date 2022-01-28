package model

import "fmt"

const (
	ErrPro = "Id should not be nill"
)

type Err struct {
	Id int
}

func (er Err) Error() string {
	return fmt.Sprint(ErrPro, er.Id)
}
