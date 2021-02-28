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
	insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	_, err = db.Sql.Exec(insertStmt)
	if err != nil {
		fmt.Println(err)
		return
	}

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	_, err = db.Sql.Exec(insertDynStmt, "Jane", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
}
