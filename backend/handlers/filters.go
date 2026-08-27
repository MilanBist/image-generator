package handlers

import "github.com/gin-gonic/gin"

func Filters(c *gin.Context){
	// show the user the filters tab with the option to select the file

}

func FilterBlur(c *gin.Context){
	// make the image blur and just return the blurred file
}

func FilterEdge(c *gin.Context){
	// detect the edge in the image and just return the edge detected image
}

func FilterSharpen(c *gin.Context){
	// fitler the image and make it look sharpen
}