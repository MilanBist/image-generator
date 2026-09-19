package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"github.com/image-generator/engine"
	"github.com/image-generator/internal/models"
)

// must satisty this pattern to handle the image
type ImageGeneration interface{
	AddRawFileToDestination(file multipart.File, fileName, userId string) (string,int, error)
	GenerateImage(exactFilePath string, userId string) (engine.AllFiles,int, error)
}

// store the uploaded and generated files metadata in the database
type UploadGenerationStore interface{
	AddUploadedFiles(uploadedMetaData models.UploadedFilesMetaData) (int, error)
	AddGeneratedFiles(generatedFilesMetaData models.GeneratedImageMetaData) (int64, error)
}

// image dimesion interface
type ImageDimensions interface{
	GetDimension(location string) (int, int, int64, error)
}

// what functionalities this handler is going to contain
type ImageGeneratorHandler struct{
	Savator		ImageGeneration
	Store 		UploadGenerationStore
	Dimension 	ImageDimensions
}

// Handle for the raw data part
func(srv *ImageGeneratorHandler) HandleImageGeneration(w http.ResponseWriter, r *http.Request){
	// get the image 
	file,header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
        http.Error(w, "failed to get file", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Input file missing.",
		})
        return
    }

	filename := header.Filename
	data := r.Context().Value("metaData")
	metaData := data.(models.ContextMetaData)
	// check for the file extension
	if filepath.Ext(filename) != ".raw"{
		fmt.Println("Wrong file name. Should be .raw file.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Input .raw file.",
		})
		return
	}

	fmt.Println("Meta data is; ", metaData)


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

	// add these data to uploadedFIles and return back the id of the uploaded fles
	size := header.Size
	mimeType := "application/octet-stream"
	fmt.Println("The size, height and the width of the file are: ", size)
	fmt.Println("Mimetype is: ", mimeType)
	fmt.Println("Filename is: ",filename)
	fmt.Println("Storage key: ", exactFilePathForRawFile)

	var uploadingMetaData models.UploadedFilesMetaData = models.UploadedFilesMetaData{
		UserId: int64(metaData.UserId),
		Filename: filename,
		StorageKey: exactFilePathForRawFile,
		FileType: ".raw",
		Mimetype: mimeType,
		FileSize: size,
	}

	uploadedId, err := srv.Store.AddUploadedFiles(uploadingMetaData)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	
	allGeneratedImageFiles,statusCode, err := srv.Savator.GenerateImage(exactFilePathForRawFile, strconv.Itoa(metaData.UserId))
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Internal server error",
		}

		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return	
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("For the jpg images: ", allGeneratedImageFiles.JpgFiles)
	fmt.Println("For the png images: ", allGeneratedImageFiles.PngFiles)

	var returningImageResponse models.FileBasedImageGenerationReturn


	// for each of the jpg files valid ones
	for _, value := range allGeneratedImageFiles.JpgFiles{
		height, width, size, err := srv.Dimension.GetDimension(value)
		if err != nil{
			if err.Error()=="invalid"{
				continue
			}		}
		splittedData := strings.Split(value, "/")

		credentials := models.GeneratedImageMetaData{
			Userid: int64(metaData.UserId),
			SourceFieldId: int64(uploadedId),
			Filename: splittedData[len(splittedData)-1],
			StorageKey: value,
			Mimetype: "image/jpg",
			Width: width,
			Height: height,
			FileSize: size,
		}

		id, err := srv.Store.AddGeneratedFiles(credentials)
		if err != nil{
			response := models.Response{
			Success: false,
			Message: err.Error(),
		}

			w.WriteHeader(500)
			json.NewEncoder(w).Encode(response)
			return	
		}
		// also fill the id as well
		credentials.Id = int64(id)


		// also add to the generageted image
		returningImageResponse.GeneratedImage = append(returningImageResponse.GeneratedImage, credentials)
	}

	// for all of the png files being uploaded
	for _, value := range allGeneratedImageFiles.PngFiles{
		height, width, size, err := srv.Dimension.GetDimension(value)
		if err != nil{
			if err.Error()=="invalid"{
				continue
			}		
		}
		splittedData := strings.Split(value, "/")

		credentials := models.GeneratedImageMetaData{
			Userid: int64(metaData.UserId),
			SourceFieldId: int64(uploadedId),
			Filename: splittedData[len(splittedData)-1],
			StorageKey: value,
			Mimetype: "image/png",
			Width: width,
			Height: height,
			FileSize: size,
		}

		id, err := srv.Store.AddGeneratedFiles(credentials)
		if err != nil{
			response := models.Response{
			Success: false,
			Message: err.Error(),
		}
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(response)
			return	
		}
		credentials.Id = int64(id)
		returningImageResponse.GeneratedImage = append(returningImageResponse.GeneratedImage, credentials)

	}


	// add the data of the actual filename
	returningImageResponse.ActualFile.Id = int64(uploadedId)
	returningImageResponse.ActualFile.Name = filename


	// final response link the uploaded files and the other files which are generated
	response := models.Response{
		Success: true,
		Message: "Successfully added the images.",
		Data: returningImageResponse,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&response)

}