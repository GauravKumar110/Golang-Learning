package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnection() (*sql.DB, error) {
	fmt.Println("Go MySQL Tutorial")
	/*serverName := "ayopop-core-stage-cluster.cluster-cnwudf2a1war.ap-south-1.rds.amazonaws.com:3306"
	user := "gaurav.kumar"
	password := "EB@3BrDG~~`#$L_4"
	dbName := "develop_ayopop" */

	serverName := "localhost:3306"
	user := "root"
	password := "Gaurav@1234"
	dbName := "blog"

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, serverName, dbName)

	db, err := sql.Open("mysql", connectionString)

	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	return db, nil
}
