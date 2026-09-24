package models

import "time"


type ContextMetaData struct{
	UserId 		int
	Email 		string
}

type UploadedFilesMetaData struct{
	Id 			int64		`json:"uploadedId"`
	UserId		int64		`json:"userId,omitempty"`
	Filename	string		`json:"filename"`
	StorageKey	string		`json:"storageKey,omitempty"`
	FileType	string		`json:"fileType"`
	Mimetype	string		`json:"mimeType"`
	FileSize 	any		    `json:"filesize"`
	Width		any			`json:"width"`
	Height		any			`json:"height"`
	CreatedAt 	time.Time 	`json:"createdAt,omitempty"`
}

type GeneratedImageMetaData struct{
	Id				any     `json:"id,omitempty"`
	Userid 		 	int64 	`json:"userId"`
	SourceFieldId	int64	`json:"sourceFileId"`
	StorageKey		string	`json:"storageKey"`
	Filename		string	`json:"filename"`
	Mimetype		string	`json:"mimetype"`
	Width			int		`json:"width"`
	Height			int		`json:"height"`
	FileSize		int64	`json:"filesize"`
}
type actualFile struct{
	Name		string 		`json:"name"`
	Id			int64  		`json:"id"` 			
}
type FileBasedImageGenerationReturn struct{
	ActualFile		actualFile				`json:"inputFile"`
	GeneratedImage	[]GeneratedImageMetaData 	`json:"images"`
	
}


type History struct {
	ID            int64          `json:"id"`
	UserID        int64          `json:"userId"`
	SourceFileID  int64          `json:"sourceFileId"`
	OutputImageID int64          `json:"outputImageId"`
	OperationType string         `json:"operationType"`
	Prompt        *string        `json:"prompt"`
	Parameters    map[string]any `json:"parameters"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"createdAt"`
	CompletedAt   time.Time      `json:"completedAt"`
}