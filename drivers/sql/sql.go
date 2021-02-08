package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Driver interface {
	Insert(context.Context, Rows, ...Field) (string, error)
	Query(context.Context, string, ...interface{}) (Rows, error)
	Update(ctx context.Context, args ...string) error
	Delete(ctx context.Context, table, column, value string) (string, error)
}

type Datatype string

type sqlDriver struct {
	sqlDb *sql.DB
}

func (s sqlDriver) Insert(ctx context.Context, rows Rows, field ...Field) (string, error) {
	panic("implement me")
}

func (s sqlDriver) Query(ctx context.Context, s2 string, i ...interface{}) (Rows, error) {
	panic("implement me")
}

func (s sqlDriver) Update(ctx context.Context, args ...string) error {
	panic("implement me")
}

func (s sqlDriver) Delete(ctx context.Context, table, column, value string) (string, error) {
	panic("implement me")
}

type Field struct {
	Datatype
	Name string
}

type Rows struct {
	Schema string
	Table  string
	Fields []Field
	Values [][]interface{}
}

func NewPostgresDbConnection(host, user, password, dbName string, port int) (Driver, error) {
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

	return &sqlDriver{sqlDb: db}, nil
}
