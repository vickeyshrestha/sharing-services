package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Driver interface {
	Insert(context.Context, Rows, ...Field) string
	Query(context.Context, string, ...interface{}) (Rows, error)
	Exec(context.Context, string, ...interface{}) (int64, error)
	Delete(context.Context, string, string, string) (int64, error)
	Update(context.Context, ...string) (int64, error)
}

type Datatype string

type Rows struct {
	Schema string
	Table  string
	Fields []Field
	Values [][]interface{}
}

type Field struct {
	Datatype
	Name string
}

func CreateDbConnection(host, user, password, dbName string, port int) (Driver, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// TODO: return the driver
	return nil, nil
}
