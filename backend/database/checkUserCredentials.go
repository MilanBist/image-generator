package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

	// user exist
	return true, id
}
