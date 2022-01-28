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
func (s *store) GetUserById(id int) (*entities.User, error) {

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

// UpdateUser Update the fields based on the input
func (s *store) UpdateUser(u entities.User) error {

	fields, values := formQuery(u)
	// form the query with the fields
	query := "update user set" + fields + " where id = ?"
	_, err := s.db.Exec(query, values...)
	if err != nil {
		return errors.New("error, no id provided, cannot update")
	}
	return nil
}

// DeleteUser Delete record from database based on input id
func (s *store) DeleteUser(id int) error {

	// define the delete query
	query := "delete from user where id = ?"

	_, err := s.db.Exec(query, id)
	if err != nil {
		return errors.New("error, not able to delete data")
	}
	return nil
}
