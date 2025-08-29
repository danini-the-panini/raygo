package util

import (
	"fmt"
	"os"
)

type Color Vec3

func WriteColor(out *os.File, pixel_color Color) {
	var r = pixel_color.X
	var g = pixel_color.Y
	var b = pixel_color.Z

	var rbyte = int(255.999 * r)
	var gbyte = int(255.999 * g)
	var bbyte = int(255.999 * b)

	fmt.Println(rbyte, " ", gbyte, " ", bbyte)
}
