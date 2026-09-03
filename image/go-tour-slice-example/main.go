package main

import "golang.org/x/tour/pic"

// https://stackoverflow.com/questions/18553246/tour-golang-org36-the-functionality-implemented-by-pic-show
func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := range pic {

		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8((x * y) / 1 * x * y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
