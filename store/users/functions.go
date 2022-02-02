package users

import "github.com/tejas/user-crud/models"

func preCheck(user models.User) (string, []interface{}) {

	var query string
	var values []interface{}

	if user.Id < 0 {
		return "", nil
	}

	if user.Name != "" {
		query += " name = ?,"
		values = append(values, user.Name)
	}

	if user.Email != "" {
		query += " email = ?,"
		values = append(values, user.Email)
	}

	if user.Phone != "" {
		query += " phone = ?,"
		values = append(values, user.Phone)
	}

	if user.Age != 0 {
		query += " age = ?,"
		values = append(values, user.Age)
	}

	query = query[:len(query)-1]
	values = append(values, user.Id)
	return query, values
}
