package users

import (
	"user-curd/models"
	"user-curd/stores"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserStorer struct {
	db *sql.DB
}

func New(db *sql.DB) stores.Store {
	return &UserStorer{db}
}

func (u *UserStorer) InsertDetails(user models.User) error {
	fmt.Println(user.Email)
	res, err := u.db.Exec("INSERT INTO User VALUES (?,?,?,?,?)", user.Id, user.Name, user.Email, user.Phone, user.Age)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println(res)
	//lastId, err := res.LastInsertId()
	//fmt.Println(lastId)

	return nil
}

func (u *UserStorer) UpdateById(user models.User) (err error) {
	if _, err = u.db.Exec("update User set Name=?, Email=?, Phone=?, Age=? where Id = ?", user.Name, user.Email, user.Phone, user.Age, user.Id); err != nil {
		fmt.Println(err)
		return errors.New("t")
	}
	return nil
}
func (u *UserStorer) DeleteById(id int) (err error) {
	if _, err = u.db.Exec("delete from User where id=?", id); err != nil {
		return
	}
	return nil
}

func (u *UserStorer) SearchById(id int) (models.User, error) {
	var user models.User
	row := u.db.QueryRow("SELECT Id,Name,Email,Phone,Age FROM User where Id=?", id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no records found")
		}
		fmt.Print(user)
		return user, fmt.Errorf("userById:%d: %v", id, err)

	}
	return user, nil
}

func (u *UserStorer) SearchDetails() ([]models.User, error) {
	var users []models.User
	results, err := u.db.Query("SELECT Id,Name,Email,Phone,Age FROM User ")
	if err != nil {
		return []models.User{}, errors.New("t")
	}
	defer results.Close()
	for results.Next() {
		var user models.User
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		if err != nil {
			return []models.User{}, errors.New("t")
		}
		users = append(users, user)
	}
	return users, nil
}
