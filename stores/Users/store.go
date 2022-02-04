package stores

import (
	"database/sql"
	"errors"
	"user-curd/model"
	"user-curd/stores"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &dbStore{db}
}

func (s *dbStore) InsertDetails(user *model.User) (*model.User, error) {
	res, err := s.db.Exec("insert into User(Name,Email,Phone,Age) values(?,?,?,?)", user.Name, user.Email, user.Phone, user.Age)
	if err != nil {
		return &model.User{}, errors.New("error in executing insert")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &model.User{}, errors.New("could not get last inserted id")
	}
	user.Id = int(id)
	return user, nil
}

func (s *dbStore) GetById(Id int) (*model.User, error) {
	var user model.User
	row := s.db.QueryRow("SELECT Id,Name,Email,Phone,Age FROM User where Id=?", Id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return &model.User{}, errors.New("no records found")
		}
	}
	return &user, nil
}

func (s *dbStore) GetAll() ([]*model.User, error) {

	var users []*model.User
	results, err := s.db.Query("SELECT Id,Name,Email,Phone,Age FROM User ")
	if err != nil {
		return []*model.User{}, err
	}
	defer results.Close()
	for results.Next() {
		var user model.User
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		if err != nil {
			return []*model.User{}, errors.New("user details not found")
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *dbStore) UpdateById(user *model.User) (*model.User, error) {
	feilds, values := BuildQuery(*user)
	query := "update User set " + feilds + "where Id = ?"
	_, err := s.db.Exec(query, values...)
	if err != nil {
		return &model.User{}, errors.New("error in updating user name")
	}

	usr, err := s.GetById(user.Id)
	return usr, nil
}

func (s *dbStore) DeleteById(Id int) (err error) {
	if _, err = s.db.Exec("delete from User where Id=?", Id); err != nil {
		return
	}
	return nil
}
