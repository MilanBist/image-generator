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
	Userid 		 	int64 
	SourceFieldId	int64
	StorageKey		string
	Filename		string
	Mimetype		string
	Width			int
	Height			int
	FileSize		int64
}