package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type sqlDriver struct {
	Sql *sql.DB
}

func NewPostgresDbConnection(host, user, password, dbName string, port int) (*sqlDriver, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return &sqlDriver{Sql: db}, nil
}
