package main

import (
	"fmt"
	"os"

	"github.com/danini-the-panini/raygo/raygo/util"
)

func main() {

	// Image

	var image_width = 256
	var image_height = 256

	// Render

	fmt.Println("P3\n", image_width, " ", image_height, "\n255")

	for j := range image_height {
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", (image_height - j), " ")
		for i := range image_width {
			var pixel_color = util.Color{
				X: float64(i) / float64(image_width-1),
				Y: float64(j) / float64(image_height-1),
				Z: 0.0,
			}
			util.WriteColor(os.Stdout, pixel_color)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}
