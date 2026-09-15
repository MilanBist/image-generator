package engine

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)
type Dimension struct{
	FileType string
}

func validateImage(location string) bool{
	file, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// if it can decode then just return file
	_, _, err = image.DecodeConfig(file)
	file.Close()
	return err == nil

}

func(d *Dimension) GetDimension(location string) (int, int, int64, error){
	if !validateImage(location){
		return -1, -1, -1, errors.New("invalid")
	}

	file, _ := os.Open(location)
	config, _, _ := image.DecodeConfig(file)

	info, err := os.Stat(location)
	if err != nil {
		return -1,-1,-1,err
	}

	fmt.Println("Width:", config.Width)
	fmt.Println("Height:", config.Height)
	fmt.Println("Size: ", info.Size())

	return config.Width, config.Height,info.Size(), nil
}