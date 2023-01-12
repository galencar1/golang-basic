package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go_test:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Zaraki Kenpachi", 1)
	stmt.Exec("Yama-Jii", 2)

	// delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(6)
}
