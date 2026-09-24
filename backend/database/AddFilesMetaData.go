package database

import (
	"context"
	"errors"
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
		`INSERT INTO "images" ("userId", "sourceFileId", "storageKey", "fileName", "mimeType", "width", "height", "fileSize") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id"`,
		generatedFilesMetaData.Userid,
		generatedFilesMetaData.SourceFieldId,
		generatedFilesMetaData.StorageKey,
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
	
func (p *PostgresData) AddToHistoryOfUser(historyData models.History) (error){
	var id int
	query := `INSERT INTO "history"("userId", "sourceFileId", "outputImageId", "operationType", "parameters", "status")
			VALUES($1, $2, $3, $4, $5, $6) RETURNING "id"
	`
	err := p.Db.QueryRow(
		context.Background(),
		query, 
		historyData.UserID, 
		historyData.SourceFileID, 
		historyData.OutputImageID,
		historyData.OperationType,
		historyData.Parameters,
		historyData.Status,
	).Scan(&id)

	fmt.Println("The error is: ", err)
	if err != nil{
		fmt.Println(err)
		return errors.New("Error in adding to database.")
	}
	return nil
}

func(p *PostgresData) GetAllUploadedFiles(userId int)([]models.UploadedFilesMetaData, error){
	query := `SELECT "id", "fileName","fileType", "mimeType", "fileSize", "width", "height", "createdAt"
			FROM "uploadedFiles" WHERE "userId" = $1
	`

	ctx := context.Background()

	rows, err := p.Db.Query(ctx, query, userId)
	

	if err != nil{
		fmt.Println("Error in getting the uploadedFiles: ", err)
		return []models.UploadedFilesMetaData{}, errors.New("error in getting data")
	}

	defer rows.Close()

	var uploadedFiles []models.UploadedFilesMetaData

	for rows.Next(){
		var uploadedModel models.UploadedFilesMetaData
		err := rows.Scan(&uploadedModel.Id, 
		&uploadedModel.Filename,
		&uploadedModel.FileType,
		&uploadedModel.Mimetype,
		&uploadedModel.FileSize,
		&uploadedModel.Width,
		&uploadedModel.Height,
		&uploadedModel.CreatedAt)

		if err != nil{
			fmt.Println(err)
			fmt.Println("Error in gettting metadata from database.")
			return []models.UploadedFilesMetaData{}, errors.New("Error in getting the value of rows.")
		}


		uploadedFiles = append(uploadedFiles, uploadedModel)
	}


	return uploadedFiles, nil



	
}