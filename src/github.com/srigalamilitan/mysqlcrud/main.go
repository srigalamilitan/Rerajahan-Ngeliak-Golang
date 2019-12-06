package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID       int64  `json:"id"`
	Username string `json:"name"`
}

func main() {
	fmt.Println("Go mysql Tutorial")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to Mysql database")
	insert, err := db.Query("INSERT into users(username) values('dari program')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted into users table")

	results, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var users Users
		err = results.Scan(&users.ID, &users.Username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("USERS==========")
		fmt.Printf("ID  %d\n", users.ID)
		fmt.Println("Name : " + users.Username)
	}

}

// install terlebih dahulu mysql
//go get github.com/go-sql-driver/mysql
