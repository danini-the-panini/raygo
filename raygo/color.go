package main

import (
	"fmt"
	"math"
	"os"
)

var BLACK = ZERO3
var WHITE = Vec3{1.0, 1.0, 1.0}

func linearToGamma(linear_component float64) float64 {
	if linear_component > 0.0 {
		return math.Sqrt(linear_component)
	}

	return 0.0
}

func WriteColor(out *os.File, pixel_color Vec3) {
	var r = linearToGamma(pixel_color.X)
	var g = linearToGamma(pixel_color.Y)
	var b = linearToGamma(pixel_color.Z)

	var intensity = Interval{0.000, 0.999}
	var rbyte = int(255.0 * intensity.clamp(r))
	var gbyte = int(255.0 * intensity.clamp(g))
	var bbyte = int(255.0 * intensity.clamp(b))

	fmt.Println(rbyte, " ", gbyte, " ", bbyte)
}
