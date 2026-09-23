package database

import (
	"context"
	"fmt"
)

// get the image's storageKey based on the imageId and userId
func(p *PostgresData) GetImageStorageKey() (string, error){
	var userId, id int
	var storageKey string
	query := `SELECT "storageKey" FROM "images" WHERE "userId" = $1 AND "id" = $2`
	err := p.Db.QueryRow(
		context.Background(),
		query,
		userId,
		id,
	).Scan(&storageKey)

	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return "", err
	}
	return storageKey, nil
}