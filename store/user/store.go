package user

import (
	"database/sql"
	"fmt"
	"zopsmart/user-curd/model"
)

type DbStore struct {
	db *sql.DB
}

func New(db *sql.DB) DbStore {
	return DbStore{db: db}
}

func (s *DbStore) GetUserById(id int) (model.User, error) {
	var u model.User
	if id < 1 {
		return u, fmt.Errorf("not a valid Id")
	}
	dis := s.db.QueryRow("SELECT * FROM User WHERE id = ?", id)
	dis.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)
	return u, nil
}

func (s *DbStore) GetAllUsers() ([]model.User, error) {
	var users []model.User
	rows, _ := s.db.Query("Select * From User")
	defer rows.Close()
	for rows.Next() {
		var usr model.User
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age); err != nil {
			fmt.Errorf("%v", err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("%v", err)
	}
	return users, nil

}

func (s *DbStore) AddUser(usr model.User) (int, error) {
	res, err := s.db.Exec("INSERT INTO User(Name,Email,Phone,Age) VALUES (?,?,?,?)", usr.Name, usr.Email, usr.Phone, usr.Age)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}
	return int(lastId), nil
}

func (s *DbStore) UpdateUser(ur model.User) error {
	if ur.Id < 1 {
		return fmt.Errorf("not a valid Id")
	}
	_, err := s.db.Exec("Update User set name=?, email=?,phone=?,age=? Where id=?", ur.Name, ur.Email, ur.Phone, ur.Age, ur.Id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil

}

func (s *DbStore) DeleteUser(id int) error {
	if id < 1 {
		return fmt.Errorf("not a valid Id")
	}
	_, err := s.db.Exec("Delete FROM User WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil

}
