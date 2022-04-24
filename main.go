package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Name string `json:"name"`
	Age uint16 `json:"age"`
}

func main()  {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert, err := db.Query("INSERT INTO `users` ( `name`, `age`) VALUES ('Og`abek', 22)")

	selects, err := db.Query("select name, age from users")
	if err != nil {
		panic(err)
	}

	for selects.Next(){
		var user Users
		err = selects.Scan(&user.Name, &user.Age)

		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("user: %s yoshi %d", user.Name, user.Age))
	}
}
