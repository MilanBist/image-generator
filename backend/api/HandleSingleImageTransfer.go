package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/image-generator/internal/models"
)

// a function which is going to get the byted image based on the request
type SingleImageTransferer interface{
	GetBytedImage(storageKey string) ([]byte, error) 
}

type ImageFetcher interface{
	GetImageStorageKey() (string, error)
}

type SingleImageProperty struct{
	Transfer 	SingleImageTransferer
	Store 		ImageFetcher	
}

func(s *SingleImageProperty) HandleSingleImageProperty(w http.ResponseWriter, r *http.Request){
	// receive the image's metadata in certain format and then get the image from the file storage and return the byted images


	// get the params
	// imageId := r.URL.Query().Get("id")
	storageKey := r.URL.Query().Get("storageKey")
	mimetype := r.URL.Query().Get("mimetype")


	if storageKey == ""{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&map[string]any{
			"status":false,
			"message":"Get the proper credentials.",
		})
		return
	}

	// now based on the storage key read the image and send response to the frontend
	bytedData, err := s.Transfer.GetBytedImage(storageKey)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response)
	}

	fmt.Println("Byted data is: ", bytedData)


	var headerType string
	if mimetype == "image/jpeg"{
		headerType = "image/jpeg"
	} else{
		headerType = "image/png"
	}
	w.Header().Add("Content-Type", headerType)
	w.WriteHeader(http.StatusOK)
	w.Write(bytedData)
}