package Users

import (
	"database/sql"
	"errors"
	"fmt"
	"user-curd/model"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) *dbStore {
	return &dbStore{db}
}

func (s *dbStore) InsertDetails(user model.User) error {
	fmt.Println(user.Email)
	res, err := s.db.Exec("INSERT INTO User VALUES (?,?,?,?,?)", user.Id, user.Name, user.Email, user.Phone, user.Age)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println(res)

	return nil
}

func (s *dbStore) GetById(Id int) (model.User, error) {
	var us model.User
	row := s.db.QueryRow("SELECT Id,Name,Email,Phone,Age FROM User where Id=?", Id)

	err := row.Scan(&us.Id, &us.Name, &us.Email, &us.Phone, &us.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return us, fmt.Errorf("no records found")
		}
		fmt.Print(us)
		return us, fmt.Errorf("userById:%d: %v", Id, err)

	}
	return us, nil
}

func (s *dbStore) GetAll() ([]model.User, error) {

	var users []model.User
	results, err := s.db.Query("SELECT Id,Name,Email,Phone,Age FROM User ")
	if err != nil {
		return []model.User{}, errors.New("t")
	}
	defer results.Close()
	for results.Next() {
		var user model.User
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		if err != nil {
			return []model.User{}, errors.New("t")
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *dbStore) UpdateById(user model.User) (err error) {
	if _, err = s.db.Exec("update User set Name=?, Email=?, Phone=?, Age=? where Id = ?", user.Name, user.Email, user.Phone, user.Age, user.Id); err != nil {
		fmt.Println(err)
		return errors.New("t")
	}
	return nil
}
func (s *dbStore) DeleteById(Id int) (err error) {
	if _, err = s.db.Exec("delete from User where Id=?", Id); err != nil {
		return
	}
	return nil
}
