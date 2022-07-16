package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Db *sql.DB
}

func CreateDb(cdn string) (*Db, error) {
	db, err := sql.Open("mysql", cdn)
	if err != nil {
		return &Db{}, err
	}
	return &Db{
		Db: db,
	}, nil
}
