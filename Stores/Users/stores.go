package Users

// import "database/sql"

// type user struct {
// 	db *sql.DB
// }

// // type user struct {
// // 	userStore store.Finder
// // }

// func New(db *sql.DB) *user {
// 	return &user{db: db}
// }

// func (u *user) Get(id int) {

// }

import (
	"Icrud/Stores"
	"Icrud/TModels"
	"database/sql"
	"errors"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) Stores.IStore {
	return &dbStore{db: db}
}

func (s *dbStore) GetUsers() ([]TModels.User, error) {
	var users []TModels.User

	rows, err := s.db.Query("select * from HUser")
	if err != nil {
		return []TModels.User{}, errors.New("cannot fetch users")
	}

	for rows.Next() {
		var u TModels.User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)
		users = append(users, u)
	}
	return users, err

}

func (s *dbStore) InsertUser(u TModels.User) (int, error) {
	sq := "Insert into HUser (name,email,phone,age) values (?,?,?,?)"
	res, err := s.db.Exec(sq, u.Name, u.Email, u.Phone, u.Age)
	if err != nil {
		return -1, err
	}
	LastId, err := res.LastInsertId()
	if err != nil {
		return -1, errors.New("error during accessing the id")
	}
	fmt.Printf("The last inserted row id: %d\n", LastId)
	return int(LastId), nil
}

// func (s *dbStore) UpdateUser(u TModels.User,id)(int,error){

// }

func (s *dbStore) UserById(id int) (TModels.User, error) {
	var user TModels.User
	if id < 1 {
		return user, errors.New("negative Id")
	}

	rows := s.db.QueryRow("select id,name,email,phone,age from HUser where id = ?", id)
	if rows.Err() != nil {
		// panic(rows.Err())
		return user, rows.Err()
	}

	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	fmt.Println(user)
	// if err != nil {
	// 	panic(err)
	// }
	return user, err

}

func (s *dbStore) DeleteUserById(id int) (int, error) {
	var iid int
	if id < 1 {
		return iid, errors.New("negative Id")
	}

	result, err := s.db.Exec("delete from HUser where id=?", id)
	if err != nil {
		return iid, err
	}
	rr, _ := result.RowsAffected()
	return int(rr), nil
}

func (s *dbStore) UpdateUserById(u TModels.User, id int) (int, error) {
	var iid int

	if id < 1 {
		return iid, errors.New("negative Id")
	}

	sq := "Update HUser set  name = ?,email = ?,phone = ?,age = ? where id =?"
	res, err := s.db.Exec(sq, u.Name, u.Email, u.Phone, u.Age, id)
	if err != nil {
		return -1, err
	}
	re, err := res.RowsAffected()
	if err != nil {
		return -1, errors.New("error during Updating the id")
	}

	return int(re), nil
}

func (s *dbStore) GetEmail(email string) (bool, error) {
	// Emails := []string{}
	searchQ := "Select email from HUser where email=?"
	rows := s.db.QueryRow(searchQ, email)

	if rows.Err() != nil {
		return true, errors.New("DataBase error")
	}
	var eml string

	err := rows.Scan(&eml)
	if err == nil {
		return true, errors.New("email Existed")
	}
	// Emails = append(Emails, eml)

	return false, nil
}
