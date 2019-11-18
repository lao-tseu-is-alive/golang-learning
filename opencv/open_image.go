package main

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
)

/*
	https://gocv.io/writing-code/more-examples/
*/

func main() {
	const defaultImagePath = "images/big_1600x1200.jpg"
	var imagePath string = defaultImagePath
	flag.StringVar(&imagePath, "file", defaultImagePath, "the path to an image file")
	flag.Parse()
	srcImg := gocv.IMRead(imagePath, gocv.IMReadColor)
	defer srcImg.Close()
	if srcImg.Empty() {
		fmt.Printf("Error reading image from: %v\n", imagePath)
		return
	}
	window := gocv.NewWindow("OpenCV Hello")
	for {
		window.IMShow(srcImg)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
