package main

import (
	"fmt"
	"sharing-services/drivers/sql"
)

const (
	host         = "192.168.202.128"
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
	defer db.Sql.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('Vickey', 1)`
	_, err = db.Sql.Exec(insertStmt)
	if err != nil {
		fmt.Println(err)
		return
	}
}
