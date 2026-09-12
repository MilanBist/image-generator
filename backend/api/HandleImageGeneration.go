package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/image-generator/internal/models"
)

// must satisty this pattern to handle the image
type ImageGeneration interface{
	AddDataToDestination(file multipart.File, fileName string) (string,int, error)
	GenerateImage(exactFilePath string) (string,int, error)
}

// what functionality is my this handler is going to have
type ImageGeneratorHandler struct{
	Savator		ImageGeneration
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
	filename  = "user"+strconv.Itoa(metaData.UserId) +"_"+ filename

	// check for the file extension
	if filepath.Ext(filename) != ".raw"{
		fmt.Println("Wrong file name. Should be .raw file.")
		http.Error(w, "Wrong file input should be .raw insted.", http.StatusBadRequest)
		return
	}


	// add the file data to the given location
	exactFilePath, statusCode, err := srv.Savator.AddDataToDestination(file, filename)	
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
		return
	}

	// now send the fileLocation to the generate image function
	_,statusCode, err = srv.Savator.GenerateImage(exactFilePath)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return	
	}
	
}