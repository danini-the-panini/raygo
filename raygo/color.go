package main

import (
	"fmt"
	"os"
)

func WriteColor(out *os.File, pixel_color Vec3) {
	var r = pixel_color.X
	var g = pixel_color.Y
	var b = pixel_color.Z

	var rbyte = int(255.999 * r)
	var gbyte = int(255.999 * g)
	var bbyte = int(255.999 * b)

	fmt.Println(rbyte, " ", gbyte, " ", bbyte)
}
