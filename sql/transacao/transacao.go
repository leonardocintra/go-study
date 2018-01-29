package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")
	stmt.Exec(200, "Bia")
	stmt.Exec(201, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave dupliada
	// nesse caso nao insere nem a Bia nem o Carlos

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
