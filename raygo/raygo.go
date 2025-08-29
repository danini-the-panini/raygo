package main

import (
	"fmt"
	"math"
	"os"
)

func HitSphere(center Vec3, radius float64, r Ray) float64 {
	var oc = Minus(center, r.Origin)
	var a = LengthSq(r.Dir)
	var h = Dot(r.Dir, oc)
	var c = LengthSq(oc) - radius*radius
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
		var n = Unit(Minus(At(r, t), Vec3{0.0, 0.0, -1.0}))
		return Times(Vec3{n.X + 1.0, n.Y + 1.0, n.Z + 1.0}, 0.5)
	}
	var unit_direction = Unit(r.Dir)
	var a = 0.5 * (unit_direction.Y + 1.0)
	var color1 = Vec3{1.0, 1.0, 1.0}
	var color2 = Vec3{0.5, 0.7, 1.0}
	return *Add(Scale(&color1, 1.0-a), *Scale(&color2, a))
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

	var pixel_delta_u = DivBy(viewport_u, float64(image_width))
	var pixel_delta_v = DivBy(viewport_v, float64(image_height))

	var viewport_upper_left = Minus(Minus(Minus(camera_center,
		Vec3{0.0, 0.0, focal_length}), DivBy(viewport_u, 2.0)), DivBy(viewport_v, 2.0))
	var pixel00_loc = Plus(viewport_upper_left, Times(Plus(pixel_delta_u, pixel_delta_v), 0.5))

	// Render

	fmt.Println("P3\n", image_width, " ", image_height, "\n255")

	for j := range image_height {
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", (image_height - j), " ")
		for i := range image_width {
			var pixel_center = pixel00_loc
			Add(Add(&pixel_center, Times(pixel_delta_u, float64(i))), Times(pixel_delta_v, float64(j)))
			var ray_direction = Minus(pixel_center, camera_center)
			var r = Ray{camera_center, ray_direction}

			var pixel_color = RayColor(r)
			WriteColor(os.Stdout, pixel_color)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}
