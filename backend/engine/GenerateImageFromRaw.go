package engine

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// to get all of the file name
type AllFiles struct{
	JpgFiles []string
	PngFiles []string
}

var allfile AllFiles

// To create the new file
var jpgFileCount int
var pngFileCount int

// which is the previous file
var whichFile string = "none"
var imageData = make([]byte, 512)


var pngHeaders = []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
// var jpgHeaders = []byte{0xff, 0xd8, 0xff}
var jpgForthHeader = []byte{0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef}

// check for path existence
func pathExistence(folderPath string) (bool, error){
	_, err := os.Stat(folderPath)

	if err == nil{
		return true, nil
	} else if errors.Is(err, os.ErrNotExist){
		return false, err
	}

	return false, err
}





// If the file size is less then 512 bytes just return false
func fileCheck(info os.FileInfo) bool{
	// find the size
	fileSize := int(info.Size())
	if fileSize < 512{
		return  false
	}
	return true
}

// forth value is diffrent for the data
func isForthValid(forth byte) bool{
	for _, value := range jpgForthHeader{
		if value == forth{
			return  true
		}
	}
	return false
}

func checkForJpg(data []byte) bool{
	if data[0] == 0xff && data[1] == 0xd8 && data[2] == 0xff && isForthValid(data[3]) == true{
		return true
	}
	return false
}

func checkForPng(data []byte) bool{
	if bytes.Contains(data, pngHeaders){
		return true
	}
	return false
}

func getImages(filename, baseOutputPath string) (error){
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Error in opening the file. ")
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, _ := os.Stat(filename)
	isValidSize := fileCheck(fileInfo)
	
	if isValidSize == true{

		var imageFilePath string
		var imageFile *os.File
		for{
			// continue reading till we find end of the file
			// in each round read 512 bytes
			n, err := file.Read(imageData)
			fmt.Println("[IMAGE GENERATION ENGINE]: The length of the data is: ", n)

			if err != nil && err != io.EOF{
				fmt.Println("Some error occured in engine. Acutal error: ", err)
				return err
			}

			

			if n > 0{
				// get the file data of 512 bytes
				data := imageData[:n]
				if checkForJpg(data) == true{
					// close the previous file and set the jpg as previous for the next file
					if jpgFileCount == 0{
						if whichFile == "png"{
							imageFile.Close()
						}
						imageFilePath = filepath.Join(baseOutputPath, "img"+ strconv.Itoa(jpgFileCount)+ ".jpg")
					} else{
						imageFile.Close()
						imageFilePath = filepath.Join(baseOutputPath, "img"+ strconv.Itoa(jpgFileCount)+ ".jpg")
					}
					imageFile, err = os.OpenFile(imageFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil{
						fmt.Println("Error in opening the file. Actual error: ", err)
						return err
					}
					whichFile = "jpg"
					jpgFileCount += 1		
					allfile.JpgFiles = append(allfile.JpgFiles, imageFilePath)		
				}else if checkForPng(data) == true{
					// close the previous file path
					if pngFileCount == 0{
						if whichFile == "jpg"{
							imageFile.Close()
						}
						imageFilePath = filepath.Join(baseOutputPath, "img"+ strconv.Itoa(pngFileCount)+ ".png")
					} else{
						imageFile.Close()
						imageFilePath = filepath.Join(baseOutputPath, "img"+ strconv.Itoa(pngFileCount)+ ".png")
					}
					imageFile, err = os.OpenFile(imageFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil{
						fmt.Println("Error in opening the file. Actual error: ", err)
						return err
					}
					whichFile = "png"
					pngFileCount += 1
					// add to the png file 
					allfile.PngFiles = append(allfile.PngFiles, imageFilePath)
				}

				// write the data simply to the file opened
				imageFile.Write(data)			
			}

			// if n < 0 just break
			if err == io.EOF {
				break
			}
		}
		imageFile.Close()
	}
	return nil
}



func GenerateImageFromRaw(rawImageFilepath, baseOutputPath string) (AllFiles, string, error){
	folderExistence, err := pathExistence(baseOutputPath)
	fmt.Println("Base output path is: ", baseOutputPath)
	fmt.Println(folderExistence, err)
	if err == os.ErrNotExist || folderExistence == false{
		// create the folder
		err := os.MkdirAll(baseOutputPath, 0755)
		if err != nil{
			fmt.Println("Error in creating the folder in destination of .", baseOutputPath)
			return allfile,"", errors.New("Error in creating file destination")
		}
	}

	// get the images based on the rawImage filepath
	fmt.Println("Raw image file path is: ", rawImageFilepath)
	err = getImages(rawImageFilepath, baseOutputPath)
	if err != nil{
		fmt.Println("Error in engine.: ", err)
		return allfile,"", err
	}

	return allfile,baseOutputPath, nil
}