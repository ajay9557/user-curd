package stores

import "user-curd/model"

func BuildQuery(user model.User) (string, []interface{}) {
	query := ""
	var values []interface{}
	if user.Name != "" {
		query += " Name = ?, "
		values = append(values, user.Name)
	}
	if user.Email != "" {
		query += " Email = ?, "
		values = append(values, user.Email)
	}
	if user.Phone != "" {
		query += " Phone = ?, "
		values = append(values, user.Phone)
	}
	if user.Age != "0" {
		query += " Age = ?,"
		values = append(values, user.Age)
	}
	values = append(values, user.Id)
	return query, values
}
