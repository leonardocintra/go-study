package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	dbString := "server=10.31.0.108;user=loja200;password=!@(Aloja2Z)*&;port=1433"
	fmt.Println(dbString)

	db, err := sql.Open("mssql", dbString)
	if err != nil {
		log.Fatal("Open dbection failed: ", err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("update dbmagazine_xp..CONTRATOS set CODIGOPRODUTO = '000000007' where NINTERNET = 20900660")
	if err != nil {
		log.Fatal("Prepare falhou: ", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	fmt.Println(row)
	fmt.Printf("Bye!")
}
