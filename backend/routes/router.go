package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/image-generator/handlers"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// first get the images from the raw data
	r.GET("/", handlers.GetImages)                           // show the user /getImages/decodeRawFile of html
	r.GET("/getImages/decodeRawFile", handlers.DecodeRawData) // let user get the file and response with the .jpg files
	r.POST("/getImages/decodeRawFile", handlers.DecodeRawData)

	// get the transformation
	r.GET("/transformation", handlers.Transformation)
	r.GET("/transformation/greyscale", handlers.TransformationGrey)
	r.GET("/transformation/increaseBrightness", handlers.TransformationIncreaseBrightness)
	r.GET("/transformation/decreaseBrightness", handlers.TransformationDecreaseBrightness)

	// get for the ressize
	r.GET("/resize", handlers.Resize)
	r.GET("/resize/makeLarge/:times", handlers.ResizeMakeLarge)

	// get for the filters
	r.GET("/filters", handlers.Filters)
	r.GET("/filters/blur", handlers.FilterBlur)
	r.GET("/filters/sharpen", handlers.FilterSharpen)
	r.GET("/filters/edge", handlers.FilterEdge)

	return r
}
