package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/image-generator/api"
	"github.com/image-generator/config"
	"github.com/image-generator/database"
)

func main(){
	// Main task is to configure the port and start the server
	cfg := config.Configure()

	// connection of the database
	conn, err := database.ConnectToPostgres()
	if err != nil{
		fmt.Println(err)
		log.Fatal("Error in configuring the postgres.")
	}

	// based on the connection with the postgres now create the server
	server, err := api.NewServer(conn, cfg)

	port := ":"+cfg.BackendPort

	fmt.Println("Server successfully started! Listening on http://localhost:8081...")
	if err = http.ListenAndServe(port, server.Router); err != nil{
		log.Fatal("Can't start the http server, Error: ",err)
	}
}