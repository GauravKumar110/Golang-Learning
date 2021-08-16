package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

/*
type PubSubConfig struct {
	TopicTransactionFailed        string mapstructure:"TOPIC_TRANSACTION_FAILED"
	SubscriptionTransactionFailed string mapstructure:"SUBSCRIPTION_TRANSACTION_FAILED"
	TopicTransactionSuccess       string mapstructure:"TOPIC_TRANSACTION_SUCCESS"
	ProjectId                     string mapstructure:"GCP_PROJECT_ID"
}

type DBDetail struct {
	DBPort     string ""
	DBHost     string ""
	DBUserName string ""
	DBPassword string ""
}

var pubSubConfig PubSubConfig

func initPubSubConfig(environment string) {
	//var pubSubConfig PubSubConfig
	switch strings.ToUpper(environment) {
	case "TEST":

	case "APP":
		viper.Unmarshal(&pubSubConfig)

	}

}*/

func DatabaseConnection() (*sql.DB, error) {
	fmt.Println("Go MySQL Tutorial")
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	//	fmt.Println(viper.ReadInConfig())

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//	fmt.Println(viper.ReadInConfig())
	// so we have to do the type assertion, to get the value
	DBPort, ok := viper.Get("DB.PORT").(string)
	//	fmt.Println(DBPort)
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

	//fmt.Printf("viper : %s = %s \n", "Database Port", DBPort)

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
