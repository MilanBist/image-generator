package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	BackendPort string
}


func Configure() *Config{
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error in loading environment variable",err)
	}

	// if no error just load the environment variable of the port
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	newConfig := &Config{}
	newConfig.BackendPort = port
	return newConfig
}
