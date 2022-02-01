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

func (connection *SqlDb) Create(user models.User) error {
	_, error := connection.Db.Exec("insert into users(name, email, phone, age) values(?, ?, ?, ?)", user.Name, user.Email, user.Phone, user.Age)
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

func (connection *SqlDb) Update(user models.User) error {

	query := "update users set"
	var args []interface{}

	if user.Name != "" {
		query += " name=?,"
		args = append(args, user.Name)
	}

	if user.Email != "" {
		query += " email=?,"
		args = append(args, user.Email)
	}

	if user.Phone != "" {
		query += " phone=?,"
		args = append(args, user.Phone)
	}

	if user.Age != 0 {
		query += " age=?"
		args = append(args, user.Age)
	}

	if user.Id > 0 {
		query += " where id=? "
		args = append(args, user.Id)
	} else {
		return errors.New("INVALID ID")
	}

	_, err := connection.Db.Exec(query, args...)
	if err != nil {
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
