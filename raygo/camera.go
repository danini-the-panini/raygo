package main

import (
	"fmt"
	"math"
	"os"
)

type Camera struct {
	image_width   int
	image_height  int
	center        Vec3
	pixel00_loc   Vec3
	pixel_delta_u Vec3
	pixel_delta_v Vec3
}

func NewCamera(aspect_ratio float64, image_width int) Camera {
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
		center,
		pixel00_loc,
		pixel_delta_u,
		pixel_delta_v,
	}
}

func (self *Camera) render(world *Group) {
	fmt.Println("P3\n", self.image_width, " ", self.image_height, "\n255")

	for j := range self.image_height {
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", (self.image_height - j), " ")
		for i := range self.image_width {
			var pixel_center = self.pixel00_loc.plus(self.pixel_delta_u.times(float64(i))).plus(self.pixel_delta_v.times(float64(j)))
			var ray_direction = pixel_center.minus(self.center)
			var r = Ray{self.center, ray_direction}

			var pixel_color = self.rayColor(r, world)
			WriteColor(os.Stdout, pixel_color)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                                 ")
}

func (self *Camera) rayColor(r Ray, world *Group) Vec3 {
	var hit = world.hit(r, Interval{0.0, math.Inf(1)})
	if hit.DidHit {
		return hit.Normal.plus(Vec3{1.0, 1.0, 1.0}).times(0.5)
	}
	var unit_direction = r.Dir.unit()
	var a = 0.5 * (unit_direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.times(1.0 - a).plus(Vec3{0.5, 0.7, 1.0}.times(a))
}
