package main

import (
	"fmt"
	"os"
)

var BLACK = Vec3{0.0, 0.0, 0.0}
var WHITE = Vec3{1.0, 1.0, 1.0}

func WriteColor(out *os.File, pixel_color Vec3) {
	var r = pixel_color.X
	var g = pixel_color.Y
	var b = pixel_color.Z

	var intensity = Interval{0.000, 0.999}
	var rbyte = int(255.0 * intensity.clamp(r))
	var gbyte = int(255.0 * intensity.clamp(g))
	var bbyte = int(255.0 * intensity.clamp(b))

	fmt.Println(rbyte, " ", gbyte, " ", bbyte)
}
