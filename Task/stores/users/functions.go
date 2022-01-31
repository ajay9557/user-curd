package store

import "zopsmart/Task/models"

func find(user models.User) (string, []interface{}){

	var query string
	var values []interface{}

	if user.Id < 0 {
		return "",nil
	}

	if user.Name != "" {
		query += " Name = ?,"
		values = append(values, user.Name)
	}

	if user.Email != "" {
		query += " Email = ?,"
		values = append(values, user.Email)
	}

	if user.Phone != "" {
		query += " Phone = ?,"
		values = append(values, user.Phone)
	}

	if  user.Age >= 0 {
		query += " Age = ?,"
		values = append(values, user.Age)
	}

	query = query[:len(query)-1]
	values = append(values, user.Id)
	return query, values
}