package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/image-generator/models"
	"github.com/image-generator/store"
)

func GetImages(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		// just show the user templ
		c.File("./static/index.html")
	} else {
		http.Error(c.Writer, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
}


func DecodeRawData(c *gin.Context) {
	// get the images from the file name
	// if the method is get then just render it to the getImages.html

	requestMethod := c.Request.Method
	if requestMethod == http.MethodGet{
		// render the getImages.html for the user
		c.File("./static/getImages.html")
	} else if requestMethod == http.MethodPost{
		// just get the raw file from the user
		file, err := c.FormFile("rawfile")
		fmt.Println(err)
		if err != nil {
			http.Error(c.Writer, "Can't writ ethe file.", http.StatusInternalServerError)
			return
		}
		
		// save the uploaded file name
		info, err1 := os.Stat("./uploads")
		if err1 == nil && info.IsDir(){
			// directory already exist no need to recreate the
			// directory
		} else{
			// directory not present so create it
			err = os.Mkdir("./uploads",os.ModePerm)
			if err != nil{
				http.Error(c.Writer, "Can't create new folder.", http.StatusInternalServerError)
				return
			}
		}

		// create a random file name in order to keep track of it
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		path := "./uploads"+filename
		// save the raw files to the uploads
		c.SaveUploadedFile(file, path)

		// as I have saved the given file to the path now my task will be to save this 
		// record to the database including th person's data

		// pass the db and file info here
		userInfo := models.UserInfo{
			Filename: filename,
			Size: file.Size,
			Path: path,
			CreatedAt: time.Now(),
		}
		// also set the contents in the given file

		// I will pass the userInfo to share for the db and also the db 
		store.SaveFileInfoToDB(userInfo)
	}
}
