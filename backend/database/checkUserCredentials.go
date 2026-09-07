package database

import (
	"context"
	"fmt"

	"github.com/image-generator/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hashedPassword string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func UserExistence(email string, db *pgxpool.Pool) (bool, int){
	// if the user exist return true and if not return false
	ctx := context.Background()
	query := `
			SELECT "id" FROM "users"
			WHERE "email" = $1
		`

	var id int
	id = -1

	// now get the id
	db.QueryRow(ctx, query, email).Scan(&id)
	if id == -1{
		return false, id
	}
	return true, id
}

func LoginData(loginCredentials models.Login, db *pgxpool.Pool) (int){
	ctx := context.Background()
	query := `
			SELECT "passwordHash", "id" FROM "users"
			WHERE "email" = $1
		`
	var password string
	var id int

	fmt.Println("Id and password are: ", password)

	err1 := db.QueryRow(ctx, query, loginCredentials.Email).Scan(&password, &id)
	if err1 != nil{
		fmt.Println("Error is: ", err1)
		return -1
	}

	if id == -1 || password == ""{
		return  id
	}
	isSame := checkPasswordHash(loginCredentials.Password, password)
	fmt.Println("Is same: ", isSame)

	fmt.Println("Correct from loginData.")
	if isSame == true{
		return id
	}
	fmt.Println("Reaching here.")
	
	return -1
}
