package database

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


func ConnectToPostgres()(*pgxpool.Pool, error){
	_ = godotenv.Load()
	databaseUrl := os.Getenv("DATABASE_URL")
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil{
		fmt.Println("Error in creating the connection pool.")
		return nil, err
	}

	    // Actually verify the connection
    if err := conn.Ping(context.Background()); err != nil {
        conn.Close()
        return nil, err
    }

	fmt.Println("Connected to postgresql successfully.")

	return conn, nil
}