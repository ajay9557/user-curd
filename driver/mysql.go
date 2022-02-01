package driver

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySql(dr string) (*sql.DB, error) {
	conn := "root:@tcp(localhost:3306)/test"

	db, err := sql.Open(dr, conn)

	if err != nil {
		return nil, errors.New("Could not able to connet to the given database driver")
	}

	return db, nil
}
