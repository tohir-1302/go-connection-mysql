package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `users` (`id`, `name`, `age`) VALUES (2, 'Mrfozil', 27)")
	if err != nil {
		panic(err)
	}
	insert.Close()
	fmt.Println("saqlandi")
}
