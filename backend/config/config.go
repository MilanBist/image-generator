package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	db *sql.DB
}

func(config *Config) connectToDB() (error, bool){
	// just connect to the db 
	err := godotenv.Load()

	if err != nil{
		// just return false
		return errors.New("Error in returing the msg."), false
	}

	// get all the data from the environment variable
	return nil, true
}

func Port() *Config{
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error in loading environment variable",err)
	}

	// if no error just load the environment variable of the port
	port := os.Getenv("PORT")
	fmt.Println(port)

	if port == ""{
		port = "8080"
	}
	newConfig := &Config{}
	newConfig.Port = port
	// also connect to the db
	newConfig.connectToDB()
	return newConfig
}
