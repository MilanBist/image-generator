package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/image-generator/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type PostgresData struct{
	Db 		*pgxpool.Pool
}

func checkPasswordHash(password, hashedPassword string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func(p *PostgresData) CheckPastInitialization(email string) (bool, int){
	// if the user exist return true and if not return false
	ctx := context.Background()
	query := `
			SELECT "id" FROM "users"
			WHERE "email" = $1
		`

	var id int
	id = -1

	// now get the id
	p.Db.QueryRow(ctx, query, email).Scan(&id)
	if id == -1{
		return false, id
	}
	return true, id
}

func(p *PostgresData) LoginUser(loginCredentials models.Login) (int, error){
	ctx := context.Background()
	query := `
			SELECT "passwordHash", "id" FROM "users"
			WHERE "email" = $1
		`
	var password string
	var id int

	fmt.Println("Id and password are: ", password)

	err1 := p.Db.QueryRow(ctx, query, loginCredentials.Email).Scan(&password, &id)
	if err1 != nil{
		fmt.Println("Error is: ", err1)
		return -1, err1
	}

	if id == -1 || password == ""{
		return  id, errors.New("No password or id detected.")
	}
	isSame := checkPasswordHash(loginCredentials.Password, password)
	fmt.Println("Is same: ", isSame)

	fmt.Println("Correct from loginData.")
	if isSame == true{
		return id,nil
	}
	fmt.Println("Reaching here.")
	
	return -1, errors.New("No same password.")
}
