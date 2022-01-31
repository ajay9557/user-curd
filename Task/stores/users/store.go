package store

import (
	"context"
	"database/sql"
	"errors"
//	"fmt"

	"zopsmart/Task/models"

	_ "github.com/go-sql-driver/mysql"
)

type DbStore struct {
	db *sql.DB
}

func New(db *sql.DB) *DbStore {
	return &DbStore{
		db: db,
	}
}

func (s *DbStore) GetAllUsersStore() ([]models.User,error) {
	var users []models.User 
	res, err := s.db.Query("select Id,Name,Email,Phone,Age from user;")

	if err!= nil {
		return []models.User{},errors.New("error while fetching data")
	}
	for res.Next() {
		var u models.User
		res.Scan(&u.Id,&u.Name,&u.Email,&u.Phone,&u.Age)
		users = append(users,u)
	}

	return users,nil

}

func (s *DbStore) GetUserById(id int) (models.User, error) {
	res, err := s.db.Query("select Id,Name,Email,Phone,Age  from user where id = ?", id)
	var u models.User
	if err != nil {
		return u, err
	}

	defer res.Close()
	for res.Next() {
		err = res.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)
		if err != nil {
			return u, err
		}
	}

	return u, nil

}

func (s *DbStore) Create(Id int, Name string, Email string, Phone string, Age int) error {

	query := "insert into users(Id, Name, Email, Phone, Age) VALUES(?,?,?,?,?)"

	_, err := s.db.ExecContext(context.TODO(), query, Id, Name, Email, Phone, Age)

	if err != nil {
		return err
	}

	return nil
}



func (u *DbStore) Update(user models.User) error {

	
	query := "update user set"
	fields, values := find(user)
	query += fields + " where id = ?"

	_, err := u.db.Exec(query, values...)

	if err != nil {
		return err
	}

	return nil
}


func (s *DbStore) Delete(id int) error {

	query := "delete from user where id = ?"
	_, err := s.db.ExecContext(context.TODO(), query, id)

	if err != nil {
		return err
	}

	return nil
}

// func (s *DbStore) GetMail(mail string) (bool, error) {
// 	Email := []string{}
// 	query := "select Email from user where Email=?;"

// 	res, err := s.db.Query(query, mail)
// 	if err != nil {
// 		return false, errors.New("error")
// 	}
// 	defer res.Close()
// 	for res.Next() {
// 		err = res.Scan(&mail)
// 		if err != nil {
// 			return false, errors.New("error")
// 		}
// 		Email = append(Email, mail)
// 	}

// 	return len(Email) == 0, nil
// }
