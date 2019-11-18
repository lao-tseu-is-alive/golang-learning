package main

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	const model = "models/res10_300x300_ssd_iter_140000.caffemodel"
	const prototxt = "models/deploy.prototxt.txt"
	const defaultImagePath = "images/test02.jpg"
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

	gocv.Resize(srcImg, &srcImg, image.Pt(300, 300), 0, 0, gocv.InterpolationCubic)
	fmt.Printf("Size is : %v \n", srcImg.Size())

	fmt.Println("Reading model ")
	net := gocv.ReadNetFromCaffe(prototxt, model)
	mean := gocv.Scalar{
		Val1: 104.0,
		Val2: 177.0,
		Val3: 123.0,
		Val4: 0,
	}
	blob := gocv.BlobFromImage(srcImg, 1.0, image.Pt(300, 300), mean, true, true)
	net.SetInput(blob, "")
	prob := net.Forward("")
	size := prob.Size()
	fmt.Println("PROB SIZE : ", size)
	for i := 0; i < prob.Total(); i += 7 {
		confidence := prob.GetFloatAt(0, i+2)
		if confidence > 0.4 {
			fmt.Printf("%d confidence: %v\n", i, confidence)
		}
	}

	performDetection(&srcImg, prob)

	blob.Close()
	prob.Close()

	window := gocv.NewWindow("OpenCV Face detect")
	for {
		window.IMShow(srcImg)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}

// performDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func performDetection(frame *gocv.Mat, results gocv.Mat) {
	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.4 {
			left := int(results.GetFloatAt(0, i+3) * float32(frame.Cols()))
			top := int(results.GetFloatAt(0, i+4) * float32(frame.Rows()))
			right := int(results.GetFloatAt(0, i+5) * float32(frame.Cols()))
			bottom := int(results.GetFloatAt(0, i+6) * float32(frame.Rows()))
			gocv.Rectangle(frame, image.Rect(left, top, right, bottom), color.RGBA{0, 255, 0, 0}, 2)
		}
	}
}
