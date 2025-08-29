package main

import (
	"fmt"
	"os"
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
			var r = float64(i) / float64(image_width-1)
			var g = float64(j) / float64(image_height-1)
			var b = 0.0

			var ir = int(255.999 * r)
			var ig = int(255.999 * g)
			var ib = int(255.999 * b)

			fmt.Println(ir, " ", ig, " ", ib)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}
