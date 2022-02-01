package stores

import "user-crud/models"

func BuildQuery(usr models.User) (string, []interface{}) {
	query := ""
	var values []interface{}
	if usr.Name != "" {
		query += " name = ?, "
		values = append(values, usr.Name)
	}
	if usr.Email != "" {
		query += " email = ?, "
		values = append(values, usr.Email)
	}
	if usr.Phone != "" {
		query += " phone = ?, "
		values = append(values, usr.Phone)
	}
	if usr.Age != 0 {
		query += " age = ? "
		values = append(values, usr.Age)
	}
	values = append(values, usr.Id)
	return query, values
}
