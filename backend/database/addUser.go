package database

import (
	"context"
	"errors"
	"github.com/image-generator/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", errors.New("Error in generating the hash password.")
	}
	return string(bytes), err
}

func RegisterUser(credentials models.Register, db *pgxpool.Pool) (int, error){
	ctx := context.Background()

	hashedPassword, err := hashPassword(credentials.Password)
	if err != nil{
		return -1, errors.New("Error in generating the hashed password.")
	}

	credentials.Password = hashedPassword

	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "users"("userName", "email", "passwordHash")
			 VALUES ($1, $2, $3)
			 RETURNING "id"
			 `


	var id int
	err = db.QueryRow(ctx, query, credentials.Username, credentials.Email, credentials.Password).Scan(&id)
	if err != nil{
		return -1, err
	}

	return id, nil
}