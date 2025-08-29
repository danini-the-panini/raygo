package main

import (
	"fmt"
	"math"
	"os"
)

type Camera struct {
	image_width         int
	image_height        int
	samples_per_pixel   int
	max_depth           int
	pixel_samples_scale float64
	center              Vec3
	pixel00_loc         Vec3
	pixel_delta_u       Vec3
	pixel_delta_v       Vec3
}

func NewCamera(aspect_ratio float64, image_width int, samples_per_pixel int, max_depth int) Camera {
	var image_height = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	var center = Vec3{0.0, 0.0, 0.0}

	// Determine viewport dimensions.
	var focal_length = 1.0
	var viewport_height = 2.0
	var viewport_width = viewport_height * (float64(image_width) / float64(image_height))

	// Calculate the vectors across the horizontal and down the verical viewport edges.
	var viewport_u = Vec3{viewport_width, 0.0, 0.0}
	var viewport_v = Vec3{0.0, -viewport_height, 0.0}

	// Calculate the horizontal and vertical delta vectors from pizel to pixel.
	var pixel_delta_u = viewport_u.divBy(float64(image_width))
	var pixel_delta_v = viewport_v.divBy(float64(image_height))

	// Calculate the location of the upper left pixel.
	var viewport_upper_left = center.minus(Vec3{0.0, 0.0, focal_length}).minus(viewport_u.divBy(2.0)).minus(viewport_v.divBy(2.0))
	var pixel00_loc = viewport_upper_left.plus(pixel_delta_u.plus(pixel_delta_v).times(0.5))

	return Camera{
		image_width,
		image_height,
		samples_per_pixel,
		max_depth,
		1.0 / float64(samples_per_pixel),
		center,
		pixel00_loc,
		pixel_delta_u,
		pixel_delta_v,
	}
}

func (cam *Camera) render(world *Group) {
	fmt.Println("P3\n", cam.image_width, " ", cam.image_height, "\n255")

	for j := range cam.image_height {
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", (cam.image_height - j), " ")
		for i := range cam.image_width {
			var pixel_color = BLACK
			for range cam.samples_per_pixel {
				var r = cam.getRay(i, j)
				pixel_color.add(cam.rayColor(r, cam.max_depth, world))
			}
			WriteColor(os.Stdout, pixel_color.times(cam.pixel_samples_scale))
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}

func (cam *Camera) getRay(i int, j int) Ray {
	// Construct a camera ray originating from the origin and directed at randomly sampled
	// point around the pixel location i, j.

	var offset = SampleSquare()
	var pixel_sample = cam.pixel00_loc.
		plus(cam.pixel_delta_u.times(float64(i) + offset.X)).
		plus(cam.pixel_delta_v.times(float64(j) + offset.Y))

	var ray_origin = cam.center
	var ray_direction = pixel_sample.minus(ray_origin)

	return Ray{ray_origin, ray_direction}
}

func (cam *Camera) rayColor(r Ray, depth int, world *Group) Vec3 {
	if depth <= 0 {
		return BLACK
	}

	var hit = world.hit(r, Interval{0.001, math.Inf(1)})
	if hit.DidHit {
		var direction = RandHemi(hit.Normal)
		direction.add(RandUnit())
		return cam.rayColor(Ray{hit.P, direction}, depth-1, world).times(0.5)
	}
	var unit_direction = r.Dir.unit()
	var a = 0.5 * (unit_direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.times(1.0 - a).plus(Vec3{0.5, 0.7, 1.0}.times(a))
}
