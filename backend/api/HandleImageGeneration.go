package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/image-generator/engine"
	"github.com/image-generator/internal/models"
)

// must satisty this pattern to handle the image
type ImageGeneration interface{
	AddRawFileToDestination(file multipart.File, fileName, userId string) (string,int, error)
	GenerateImage(exactFilePath string, userId string) (engine.AllFiles,int, error)
}

// store the image's metadata to the database as well
type ImageDimensiongetter interface{
	GetImageDimension(file multipart.File) (int, int, error)
}

// what functionality is my this handler is going to have
type ImageGeneratorHandler struct{
	Savator		ImageGeneration
	Dimension 	ImageDimensiongetter
}

// Handle for the raw data part
func(srv *ImageGeneratorHandler) HandleImageGeneration(w http.ResponseWriter, r *http.Request){
	// get the image 
	file,header, err := r.FormFile("file")

	if err != nil {
		fmt.Println(err)
        http.Error(w, "failed to get file", http.StatusBadRequest)
        return
    }

	filename := header.Filename

	// get the userId and create the full fileName
	// get the user_id as well from the context
	data := r.Context().Value("metaData")
	metaData := data.(models.ContextMetaData)

	// check for the file extension
	if filepath.Ext(filename) != ".raw"{
		fmt.Println("Wrong file name. Should be .raw file.")
		http.Error(w, "Wrong file input should be .raw insted.", http.StatusBadRequest)
		return
	}


	fmt.Println("Meta data is; ", metaData)

	// get the dimension of the file 


	// add the file data to the given location
	exactFilePathForRawFile, statusCode, err := srv.Savator.AddRawFileToDestination(file, filename, strconv.Itoa(metaData.UserId))	
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
		return
	}

	size := header.Size
	width, height, err := srv.Dimension.GetImageDimension()

	// now send the fileLocation to the generate image function
	allGeneratedImageFiles,statusCode, err := srv.Savator.GenerateImage(exactFilePathForRawFile, strconv.Itoa(metaData.UserId))
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return	
	}

	fmt.Println(allGeneratedImageFiles)
	
}