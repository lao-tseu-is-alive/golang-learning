package main

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

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
	fmt.Printf("Size is : %v \n", srcImg.Size())
	gocv.Resize(srcImg, &srcImg, image.Pt(0, 0), 0.5, 0.5, gocv.InterpolationCubic)
	fmt.Printf("Size is : %v \n", srcImg.Size())

	window := gocv.NewWindow("OpenCV Hello")
	for {
		window.IMShow(srcImg)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
