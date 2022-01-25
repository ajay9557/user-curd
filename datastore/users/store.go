package users

import (
	"database/sql"
	"errors"
	"log"
	"user-curd/entities"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

// GetUserById Retrieve a single user from database
func (s *store) GetUserById(id interface{}) (*entities.User, error) {

	var user entities.User

	//define the query
	query := "select id, name, email, phone, age from user where id = ?"

	// Execute the query and return the user struct and error if any
	err := s.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		return &entities.User{}, errors.New("error fetching from database, id not found")
	}
	return &user, err
}

// GetAllUsers Retrieve all the users form the database
func (s *store) GetAllUsers() ([]*entities.User, error) {
	var user []*entities.User

	// define the query
	query := "select id,name,email,phone,age from user"

	rows, err := s.db.Query(query)
	if err != nil {
		return []*entities.User{}, errors.New("error fetching data from database")
	}

	// iterate through the result and add it to the user slice
	for rows.Next() {
		var u entities.User
		_ = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)
		user = append(user, &u)
	}
	return user, nil
}

// CreateUser Create a new user entry
func (s *store) CreateUser(u entities.User) error {

	// define the query
	query := "insert into user values (?,?,?,?,?)"

	_, err := s.db.Exec(query, u.Id, u.Name, u.Email, u.Phone, u.Age)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func formQuery(u entities.User) (string, []interface{}) {
	// declare a variable to hold fields to be updated
	var fields string
	var values []interface{}

	if u.Id == 0 {
		return "", nil
	}
	if u.Name != "" {
		fields += " name = ?,"
		values = append(values, u.Name)
	}
	if u.Email != "" {
		fields += " email = ?,"
		values = append(values, u.Email)
	}
	if u.Phone != "" {
		fields += " phone = ?,"
		values = append(values, u.Phone)
	}
	if u.Age != 0 {
		fields += " age = ?,"
		values = append(values, u.Age)
	}
	fields = fields[:len(fields)-1]
	values = append(values, u.Id)
	return fields, values
}

// UpdateUser Update the fields based on the input
func (s *store) UpdateUser(u entities.User) error {

	fields, values := formQuery(u)
	//if fields == "" || values == nil {
	//	return errors.New("error, no id provided, cannot update")
	//}

	// form the query with the fields
	query := "update user set" + fields + " where id = ?"

	//query := "update user set name = ? where id = ?"

	// update user set name = ? where id = ?
	_, err := s.db.Exec(query, values...)
	if err != nil {
		return errors.New("error, no id provided, cannot update")
	}
	return nil
}

// DeleteUser Delete record from database based on input id
func (s *store) DeleteUser(id interface{}) error {

	// define the delete query
	query := "delete from user where id = ?"

	_, err := s.db.Exec(query, id)
	if err != nil {
		return errors.New("error, not able to delete data")
	}
	return nil
}
