package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func DatabaseConnection() (*sql.DB, error) {
	fmt.Println("Go MySQL Tutorial")
	/*serverName := "ayopop-core-stage-cluster.cluster-cnwudf2a1war.ap-south-1.rds.amazonaws.com:3306"
	user := "gaurav.kumar"
	password := "EB@3BrDG~~`#$L_4"
	dbName := "develop_ayopop" */

	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	//viper.SetDefault("DB.HOST", "127.0.0.1")

	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DBPort, ok := viper.Get("DB.PORT").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	DBHost, ok := viper.Get("DB.HOST").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	DBUserName, ok := viper.Get("DB.USERNAME").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	DBPassword, ok := viper.Get("DB.PASWORD").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dbName1, ok := viper.Get("DB.NAME").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	fmt.Printf("viper : %s = %s \n", "Database Port", DBPort)

	serverName := DBHost + ":" + DBPort
	user := DBUserName
	password := DBPassword
	dbName := dbName1

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
