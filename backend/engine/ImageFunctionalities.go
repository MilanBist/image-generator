package engine


import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"log"
	"os"
	"strconv"
)

func CheckFilePresent(filepath string) bool {
	_, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File not present or error in opening file path. ")
		fmt.Println("Provide file with correct filePath.")
		return false
	}
	return true
}

// make the image grey
func makeGreyScale(filepath string) {
	// open the given file
	file, err1 := os.Open(filepath)
	if err1 != nil {
		log.Println("Error in opening the filepath. ", err1)
	}

	// now find the bound of the given file
	// decode the given file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		log.Println("Error in decoding the file. ", img)
	}

	// now I have the access to the bounds of the given file
	b := img.Bounds()

	// setup for the newGray which will store the gray pix of each of pixels that I got
	grayImage := image.NewGray(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			// get the pixel at each of the point here
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)

			grayPix := 0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)

			// set the gray Pix to the given grayImage
			grayImage.SetGray(x, y, color.Gray{Y: uint8(grayPix)})
		}
	}
	// create a folder named as the gray file
	// foldername is greyscaleImage
	folderPath := "greyscale"

	folderInfo, err3 := os.Stat(folderPath)
	if err3 != nil {
		// folder doesn't exist so create one
		err4 := os.Mkdir(folderPath, 0755)

		if err4 != nil {
			log.Println("Error in making a folder")
		}
	}

	if folderInfo.IsDir() {
		// now get the filepath
		filePath := "grey" + filepath

		newFilePath := folderPath + "/" + filePath

		// create the given file path and just append in it
		newFile, err5 := os.Create(newFilePath)
		if err5 != nil {
			log.Println("Error in making the file for the grey image")
		}

		err6 := jpeg.Encode(newFile, grayImage, &jpeg.Options{Quality: 100})
		if err6 != nil {
			log.Println("Error in encoding the grayImage file", err6)
		}
	}

}

// increase the brightness of the image
func increaseBrightness(filepath string) {
	// add certain value to the r,g and b value in order to make it more brightening

	// open the file
	file, err1 := os.Open(filepath)

	if err1 != nil {
		log.Println("Error in opening the file")
	}
	// decode the file
	img, _, err2 := image.Decode(file)
	if err2 != nil {
		log.Println("Error in decoding the file", err2)
	}
	// get each of the pixel grid
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	// make a new grid with the rgba
	brightedImage := image.NewRGBA(b)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pix := img.At(x, y)
			r, g, b, a := pix.RGBA()

			// convert each of them in the uint8 format
			R := int(uint8(r >> 8))
			G := int(uint8(g >> 8))
			B := int(uint8(b >> 8))
			A := uint8(a >> 8)

			// check for the overflow
			R = R + 20
			if R > 255 {
				R = 255
			}
			G = G + 20
			if G > 255 {
				G = 255
			}
			B = B + 20
			if B > 255 {
				B = 255
			}

			r1 := uint8(R)
			g1 := uint8(G)
			b1 := uint8(B)

			// now store them all in the brightedImge place
			brightedImage.SetRGBA(x, y, color.RGBA{r1, g1, b1, A})
		}
	}
	// add it to the new file present in the new folder of the brightImage
	// make the folder if not present
	folderpath := "brightedImage"
	filepath = "brighted" + filepath
	_, err3 := os.Stat(folderpath)

	if err3 != nil {
		// folder is not present so create it
		err4 := os.Mkdir(folderpath, 0755)
		if err4 != nil {
			log.Println("Error making the file: ", err4)
		}
	}

	// since I had the folder just add to it
	newPath := folderpath + "/" + filepath
	newFile, err5 := os.Create(newPath)
	if err5 != nil {
		log.Println("Error in creating a file. ", newFile)
	}

	// now  just add the things
	err6 := jpeg.Encode(newFile, brightedImage, &jpeg.Options{Quality: 100})
	if err6 != nil {
		log.Fatal("Error in encoding the given file. ", err6)
	}

}

// decrease the brightness of the image
func decreaseBrightness(filepath string) {
	// subtract the given value with the delta value to make it lose it being more brightening
	// open the file
	file, err1 := os.Open(filepath)
	if err1 != nil {
		log.Println("Error in opening the file. ", err1)
	}

	// decode the contents and make a new RGB there
	img, _, err2 := image.Decode(file)
	if err2 != nil {
		log.Println("Error in decoding the file. ", err2)
	}

	// find the bounds
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	darkenedImage := image.NewNRGBA(b)
	// get each of the pixel and make a new whole grid from it with the RGBA
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			R := int(uint8(r >> 8))
			G := int(uint8(g >> 8))
			B := int(uint8(b >> 8))
			A := uint8(a >> 8)

			gamma := 40
			// now subtract certain
			R = R - gamma
			if R < 0 {
				R = 0
			}
			G = G - gamma
			if G < 0 {
				G = 0
			}
			B = B - gamma
			if B < 0 {
				B = 0
			}

			// again convert each of the following in the same type
			r1 := uint8(R)
			g1 := uint8(G)
			b1 := uint8(B)

			// set at the position of the given one
			darkenedImage.SetNRGBA(x, y, color.NRGBA{r1, g1, b1, A})
		}
	}

	folderPath := "darkened"
	// check if the given file exist or not
	_, err5 := os.Stat(folderPath)
	if err5 != nil {
		// since the folder doesn't exist create this folder
		err6 := os.Mkdir(folderPath, 0755)
		if err6 != nil {
			log.Print("Error in opening the file. ", err6)
		}
	}
	// now save it to the file of the given one
	newPath := folderPath + "/" + "darkened" + filepath
	finalFile, err3 := os.Create(newPath)
	if err3 != nil {
		log.Println("Error in creating a file. ", err3)
	}

	// now add the data to the given file
	err4 := jpeg.Encode(finalFile, darkenedImage, &jpeg.Options{Quality: 100})
	if err4 != nil {
		log.Println("Error in encoding the file. ", err4)
	}
}

// all done
func performTransformations(specificity, filepath string) {
	switch specificity {
	case "greyscale":
		makeGreyScale(filepath)
	case "increaseBrightness":
		increaseBrightness(filepath)
	case "decreaseBrightness":
		decreaseBrightness(filepath)
	}
}

// making the image large
func makeLarge(filepath string, times int) {
	// read the file and get it bounds and create a new bound from it which will be of double size perfectly

	// open the file
	file, err1 := os.Open(filepath)

	if err1 != nil {
		log.Println("Error in opening the file for resize. ", err1)
	}

	// decode the file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		log.Println("Error in decoding the file. ", err2)
	}

	// find the bounds of the given image and just make it double
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	newMaxX := maxX * times
	newMaxY := maxY * times

	newRectange := image.Rect(0, 0, newMaxX, newMaxY)
	// make a newRGBA color model to store the new Image\
	largedImage := image.NewNRGBA(newRectange)

	// loop over each of the pixels as a grid
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// get the pixel at the given region
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)
			A := uint8(a >> 8)

			if times == 1 {
				largedImage.SetNRGBA(x, y, color.NRGBA{R, G, B, A})
				continue
			}

			// set at the position 2x,2y then 2x,2y+1 then 2x+1, 2y and at last 2x+1, 2y+1

			if times == 2 {
				largedImage.SetNRGBA(2*x, 2*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x+1, 2*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x, 2*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x+1, 2*y+1, color.NRGBA{R, G, B, A})
				continue
			}

			if times == 3 {
				largedImage.SetNRGBA(3*x, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x, 3*y+2, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y+2, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y+2, color.NRGBA{R, G, B, A})
			}

		}
	}

	// make a newFile called as largeimages
	folderPath := "largeFolder"

	_, err5 := os.Stat(folderPath)
	if err5 != nil {
		// means like the folder doesn't exist so that create
		err6 := os.Mkdir(folderPath, 0755)
		if err6 != nil {
			log.Println("Error in creating a new folder to create a new folder for it")
		}
	}
	// create a new file path and also
	filepath = folderPath + "/" + strconv.Itoa(times) + "times" + "larged" + filepath
	// just print it in the new image
	newImage, err3 := os.Create(filepath)

	if err3 != nil {
		log.Println("Error in creating a new file.", err3)
	}
	err4 := jpeg.Encode(newImage, largedImage, &jpeg.Options{Quality: 100})
	if err4 != nil {
		log.Println("Error in encoding the image in new jpeg file.")
	}

}

// what to do under resizing
func performResize(specificity, filepath string) {

	switch specificity {
	case "onex":
		makeLarge(filepath, 1)
	case "twox":
		makeLarge(filepath, 2)
	case "threex":
		makeLarge(filepath, 3)
	}
}

// making the image look blur
func makeBlur(filepath string) {
	// open the given image
	file, err1 := os.Open(filepath)

	if err1 != nil {
		log.Println("Error in opening the image file. ", err1)
	}

	// decode the given file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		log.Println("Error in decoding the image file. ", err2)
	}

	// get the bounds of the given image
	b := img.Bounds()

	// set the new NRGBA for the given image
	blurredImage := image.NewNRGBA(b)

	// get the max X and max Y till where to get the loop
	maxX := b.Max.X
	maxY := b.Max.Y

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// get the pixel at the given range
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			r = r * 0
			g = g * 0
			b = b * 0
			A := uint8(a >> 8)

			// since I got the r,g,b,a value for the given pixel now replace it with the pixel
			// Containing 3x3 matrix including the point itself
			var count uint32
			count = 0
			for i := x - 1; i <= x+1; i++ {
				for j := y - 1; j <= y+1; j++ {
					//  get the image pixel at each of the place
					if j >= 0 && i >= 0 && i < maxX && j < maxY {
						// get the r,g,b and a value at that position
						initialPixel := img.At(i, j)
						r1, g1, b1, _ := initialPixel.RGBA()
						r += r1
						g += g1
						b += b1
						count += 1
					}
				}
			}

			// find the average of all of the given values and then do it
			r = r / count
			g = g / count
			b = b / count

			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)

			blurredImage.SetNRGBA(x, y, color.NRGBA{R, G, B, A})
		}
	}

	// create a newDirectory to store blur
	folderPath := "blur"
	_, err4 := os.Stat(folderPath)

	if err4 != nil{
		// means it doesn't exist so create it
		err5 := os.Mkdir("blur", 0755)
		if err5 != nil{
			log.Println("Error in creating a newFolder.")
		}

	}
	
	// write in the file insdie of the folder
	filepath = folderPath +"/"+ "blured" + filepath


	// create a new file and just save the given content
	imageFile, err3 := os.Create(filepath)

	if err3 != nil {
		log.Println("Error in creating the file. ", err3)
	}

	// Now insert in to this image
	err7 := jpeg.Encode(imageFile, blurredImage, nil)

	if err7 != nil{
		log.Println("Error in encoding inside the file.", err7)
	} else{
		fmt.Println("Please check in the blur folder.")
	}
}

// making the image look more sharp enough
func makeSharpen(filepath string) {
	// // Sharpening of image via Leplacian filter
	// // basic lap matrix [0 1 0][1-4 1][0 1 0]

	// lapMatrix := [][]int{{0,1,0},{1,-4,1},{0,1,0}}

	// // open the image file
	// file, err1 := os.Open(filepath)

	// if err1 != nil{
	// 	fmt.Println("Error in opening the file with path ", filepath)
	// 	log.Fatal(err1)
	// }

	// // find boundary and get each of the given pixels
	// img, _, err2 := image.Decode(file)

	// if err2 != nil{
	// 	fmt.Println("Error in decoding the file. ")
	// 	log.Fatal(err2)
	// }

	// // find the bounds of teh given files
	// b := img.Bounds()
	// maxY := b.Max.Y
	// maxX := b.Max.X

	// newImg := image.NewNRGBA(b)
	// for y:=0; y<maxY; y++{
	// 	for x:=0; x<maxX; x++{
	// 		// get the pixels at each of the positon
	// 		pixel := img.At(x,y)
	// 		r,g,b,a := pixel.RGBA()

	// 		// now get the matrix of the pixels of the neighbouring pixels from image
	// 		var lapr, lapg, lapb uint32
	// 		if x>0 && x < maxX-1 && y>0 && y<maxX-1{
	// 		for i:=x-1; i<=x+1; i++{
	// 			for j:=y-1; j<= y+1; j++{
	// 				// get the pixels at each of the given position
	// 				currentPix := img.At(i,j)
	// 				r1,g1,b1, _ := currentPix.RGBA()
					
	// 			}
	// 		}
	// 	}
	// 	}
	// }
}

// detecting the edges of the images
func detectEdge(filepath string) {

}

func performFilters(specificity, filepath string) {
	// blur sharpen and edge detection
	switch specificity {
	case "blur":
		makeBlur(filepath)
	case "sharpen":
		makeSharpen(filepath)
	case "edge":
		detectEdge(filepath)
	}
}

func PerformAction(command, specificity, filepath string) {
	// if command is transformations
	switch command {
	case "transformations":
		performTransformations(specificity, filepath)

	case "resize":
		performResize(specificity, filepath)

	case "filters":
		performFilters(specificity, filepath)
	}
}