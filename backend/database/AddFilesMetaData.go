package database

import (
	"context"
	"fmt"

	"github.com/image-generator/internal/models"
)

func(p *PostgresData) AddUploadedFiles(uploadedMetaData models.UploadedFilesMetaData)(int, error){
	
	var id int64

	err := p.Db.QueryRow(
		context.Background(),
		`INSERT INTO "uploadedFiles" ("userId", "fileName", "storageKey", "fileType", "mimeType", "fileSize") 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`,
		uploadedMetaData.UserId,
		uploadedMetaData.Filename,
		uploadedMetaData.StorageKey,
		uploadedMetaData.FileType,
		uploadedMetaData.Mimetype,
		uploadedMetaData.FileSize,
	).Scan(&id)

	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return -1, err
	}
		return int(id), nil
}

	

func(p *PostgresData) AddGeneratedFiles(generatedFilesMetaData models.GeneratedImageMetaData)(int64, error){
	var id int64
	err := p.Db.QueryRow(
		context.Background(),
		`INSERT INTO "images" ("userId", "sourceFileld", "storageKey", "fileName", "mimeType", "width", "height", "fileSize") 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`,
		generatedFilesMetaData.Userid,
		generatedFilesMetaData.SourceFieldId,
		generatedFilesMetaData.Filename,
		generatedFilesMetaData.Mimetype,
		generatedFilesMetaData.Width,
		generatedFilesMetaData.Height,
		generatedFilesMetaData.FileSize,
	).Scan(&id)
	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return id, err
	}
	return id, nil
}
	
	
