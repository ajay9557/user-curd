package users

import (
	"layer/user/models"
)

func formUpdateQuery(u models.User) (fields string, args []interface{}) {
	if u.Name != "" {
		fields += " name = ?,"
		args = append(args, u.Name)
	}

	if u.Email != "" {
		fields += " email = ?,"
		args = append(args, u.Email)
	}

	if u.Phone != "" {
		fields += " phone = ?,"
		args = append(args, u.Phone)
	}

	if u.Age > 0 {
		fields += " age = ?,"
		args = append(args, u.Age)
	}

	return
}
