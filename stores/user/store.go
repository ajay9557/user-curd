package user

import (
	"database/sql"
	"errors"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/stores"
)

type SqlDb struct {
	Db *sql.DB
}

func New(db *sql.DB) stores.Crud {
	return &SqlDb{Db: db}
}

func (connection *SqlDb) Create(name string, email string, phone string, age int) error {
	_, error := connection.Db.Exec("insert into users(name, email, phone, age) values(?, ?, ?, ?)", name, email, phone, age)
	if error != nil {
		return errors.New("ERROR IN INSERTING DATA")
	}
	return nil
}

func (connection *SqlDb) ReadOne(id int) (models.User, error) {
	user := models.User{}
	row := connection.Db.QueryRow("select * from users where id=?", id)
	_ = row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	return user, nil
}

func (connection *SqlDb) ReadAll() ([]models.User, error) {
	userList := []models.User{}
	rows, error := connection.Db.Query("select * from users")
	if error != nil {
		return userList, errors.New("ERROR IN FETCHING ROWS")
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		_ = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		userList = append(userList, user)
	}
	return userList, nil
}

func (connection *SqlDb) Update(id int, name string, email string, phone string, age int) error {
	_, error := connection.Db.Exec("update users set name=?, email=?, phone=?, age=? where id=?", name, email, phone, age, id)
	if error != nil {
		return errors.New("FAILED TO UPDATE USER DATA")
	}
	return nil
}

func (connection *SqlDb) Delete(id int) error {
	_, error := connection.Db.Exec("delete from users where id=?", id)
	if error != nil {
		return errors.New("FAILED TO DELETE USER")
	}
	return nil
}
