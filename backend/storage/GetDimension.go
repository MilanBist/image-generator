package storage


import (
	"image"
	"mime/multipart"
)

// return the dimension of the file here
func GetImageDimension(file multipart.File)(int,int, error){
	config, _, err := image.DecodeConfig(file)
	if err != nil{
		return -1, -1, err
	}
	config.
	return config.Height, config.Width, nil
}