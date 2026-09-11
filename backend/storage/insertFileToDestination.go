package storage

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"github.com/image-generator/engine"
)


type StoreFile struct{
	BasePath string
}


// check for path existence
func pathExistence(filePath string) (bool, error){
	_, err := os.Stat(filePath)

	if err == nil{
		return true, nil
	} else if errors.Is(err, os.ErrNotExist){
		return false, err
	}

	return false, err
}

// add file to the given location
func addToLocation(file multipart.File, filepath string) error{

	fmt.Println(filepath)
	inputFile, err := os.Create(filepath)
	if err != nil {
		// Error in creating the file
		fmt.Println("Error in creating the file. Actual error is: ", err)
		return errors.New("Error in creating the file.")
	}

	defer inputFile.Close()

	//copy each item to the given inputFile
	_, err =io.Copy(inputFile, file)
	if err != nil{
		fmt.Println("Error in copying the file item. Actual error: ", err)
		return errors.New("Error in copying the file item.")
	}
	return nil
}



// add the file to the certain location
func (f *StoreFile) AddDataToDestination(file multipart.File, fileName string) (string, int, error){
	fmt.Println("Base path is: ",f.BasePath)
	folderExistence, err := pathExistence(f.BasePath)
	fmt.Println("File existence: ", folderExistence)
	if err == os.ErrNotExist && folderExistence == false{
		// create the folder
		err := os.Mkdir(f.BasePath, 0755)
		if err != nil{
			fmt.Println("Error in creating the folder in destination of .", f.BasePath + fileName)
			return "", http.StatusInternalServerError, errors.New("Error in creating file destination")
		}

		fmt.Println("Created Successfully.")
	}


	err = addToLocation(file, f.BasePath+fileName)
	fmt.Println("File creation failed.")
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("Error in creating the file.")
	}
	return "", http.StatusAccepted, nil
}

// Generate image from the raw file
func (f *StoreFile) GenerateImage(exactFilePath string) (string,int, error){
	//generate image from the raw file
	allImageFiles, baseOutputPath, err := engine.GenerateImageFromRaw(exactFilePath, "../outputImages/")
	
	if err != nil{
		fmt.Println("Internal Server error: ",err)
		return  "", -1, err
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("All images files are: ", allImageFiles)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return baseOutputPath, http.StatusAccepted, nil
}