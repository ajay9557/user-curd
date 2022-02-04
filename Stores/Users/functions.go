package Users

import "Icrud/TModels"

func formUpdateQuery(u TModels.User) (string, []interface{}) {
	var where string = ""
	var value []interface{}
	if u.Name != "" {
		where = where + "name = ?,"
		value = append(value, u.Name)
	}
	if u.Email != "" {
		where = where + "email = ?,"
		value = append(value, u.Email)
	}
	if u.Phone != "" {
		where = where + "phone = ?,"
		value = append(value, u.Phone)
	}
	if u.Age != 0 {
		where = where + "age = ?,"
		value = append(value, u.Age)
	}

	return where, value
}
