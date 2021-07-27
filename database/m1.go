package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	id    int
	email string
	phone string
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "gaurav.kumar:EB@3BrDG~~`#$L_4@tcp(ayopop-core-stage-cluster.cluster-cnwudf2a1war.ap-south-1.rds.amazonaws.com:3306)/develop_ayopop")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results, err := db.Query("Select id, email, phone from members order by id desc limit 3,2,k")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user users
		err := results.Scan(&user.id, &user.email, &user.phone)
		if err != nil {
			panic(err.Error())
		}

		json.NewEncoder(w).Encode(user)

		log.Print(user.id)
		log.Printf(user.email)
		log.Printf(user.phone)
	}

	// defer the close till after the main function has finished
	// executing

	fmt.Println("Go MySQL Tutorial1")

}
