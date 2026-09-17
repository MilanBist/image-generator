package models


type ContextMetaData struct{
	UserId 		int
	Email 		string
}

type UploadedFilesMetaData struct{
	UserId		int64
	Filename	string
	StorageKey	string
	FileType	string
	Mimetype	string
	FileSize 	int64
	Width		int
	Height		int
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