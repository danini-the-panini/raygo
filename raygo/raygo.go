package main

import (
	"fmt"
	"math"
	"os"
)

func HitSphere(center Vec3, radius float64, r Ray) float64 {
	var oc = center.minus(r.Origin)
	var a = r.Dir.lenSq()
	var h = r.Dir.dot(oc)
	var c = oc.lenSq() - radius*radius
	var discriminant = h*h - a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (h - math.Sqrt(discriminant)) / a
	}
}

func RayColor(r Ray) Vec3 {
	var t = HitSphere(Vec3{0.0, 0.0, -1.0}, 0.5, r)
	if t > 0.0 {
		var n = r.at(t).minus(Vec3{0.0, 0.0, -1.0}).unit()
		return Vec3{n.X + 1.0, n.Y + 1.0, n.Z + 1.0}.times(0.5)
	}
	var unit_direction = r.Dir.unit()
	var a = 0.5 * (unit_direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.times(1.0 - a).plus(Vec3{0.5, 0.7, 1.0}.times(a))
}

func main() {

	// Image

	var aspect_ratio = 16.0 / 9.0
	var image_width = 400

	var image_height = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	// Camera

	var focal_length = 1.0
	var viewport_height = 2.0
	var viewport_width = viewport_height * (float64(image_width) / float64(image_height))
	var camera_center = Vec3{0.0, 0.0, 0.0}

	var viewport_u = Vec3{viewport_width, 0.0, 0.0}
	var viewport_v = Vec3{0.0, -viewport_height, 0.0}

	var pixel_delta_u = viewport_u.divBy(float64(image_width))
	var pixel_delta_v = viewport_v.divBy(float64(image_height))

	var viewport_upper_left = camera_center.minus(Vec3{0.0, 0.0, focal_length}).minus(viewport_u.divBy(2.0)).minus(viewport_v.divBy(2.0))
	var pixel00_loc = viewport_upper_left.plus(pixel_delta_u.plus(pixel_delta_v).times(0.5))

	// Render

	fmt.Println("P3\n", image_width, " ", image_height, "\n255")

	for j := range image_height {
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", (image_height - j), " ")
		for i := range image_width {
			var pixel_center = pixel00_loc.plus(pixel_delta_u.times(float64(i))).plus(pixel_delta_v.times(float64(j)))
			var ray_direction = pixel_center.minus(camera_center)
			var r = Ray{camera_center, ray_direction}

			var pixel_color = RayColor(r)
			WriteColor(os.Stdout, pixel_color)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}
