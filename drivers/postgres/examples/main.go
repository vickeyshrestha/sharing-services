package main

import (
	sql2 "database/sql"
	"fmt"
	"github/sharing-services/drivers/sql"
)

const (
	host         = "localhost"
	user         = "postgres"
	password     = "admin"
	databaseName = "godzilla"
	port         = 5104
)

func main() {
	db, err := sql.NewPostgresDbConnection(host, user, password, databaseName, port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Sql *sql2.DB) {
		err := Sql.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(db.Sql)

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('Vickey', 1)`
	_, err = db.Sql.Exec(insertStmt)
	if err != nil {
		fmt.Println(err)
		return
	}
}
