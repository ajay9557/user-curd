package users

import "user-curd/entities"

func formQuery(u entities.User) (string, []interface{}) {
	// declare a variable to hold query to be updated
	var query string
	var values []interface{}

	if u.Id < 0 {
		return "", nil
	}
	if u.Name != "" {
		query += " name = ?,"
		values = append(values, u.Name)
	}
	if u.Email != "" {
		query += " email = ?,"
		values = append(values, u.Email)
	}
	if u.Phone != "" {
		query += " phone = ?,"
		values = append(values, u.Phone)
	}
	if u.Age != 0 {
		query += " age = ?,"
		values = append(values, u.Age)
	}
	query = query[:len(query)-1]
	values = append(values, u.Id)
	return query, values
}
