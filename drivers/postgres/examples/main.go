package main

import (
	"fmt"
	"sharing-services/drivers/sql"
)

func main() {
	db, err := sql.NewPostgresDbConnection("", "", "", "", 000)
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('Vickey', 1)`
	_, err = db.Sql.Exec(insertStmt)
	if err != nil {
		fmt.Println(err)
		return
	}
}
