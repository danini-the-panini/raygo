package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

type Pixel struct {
	X     int
	Y     int
	Color Vec3
}

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
	u                   Vec3
	v                   Vec3
	w                   Vec3
	defocus_angle       float64
	defocus_disk_u      Vec3
	defocus_disc_v      Vec3
}

func NewCamera(
	aspect_ratio float64,
	image_width int,
	vfov float64,
	lookfrom Vec3,
	lookat Vec3,
	vup Vec3,
	defocus_angle float64,
	focus_dist float64,
	samples_per_pixel int,
	max_depth int,
) Camera {
	var image_height = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	var center = lookfrom

	// Determine viewport dimensions.
	var theta = Deg2Rad(vfov)
	var h = math.Tan(theta / 2.0)
	var viewport_height = 2.0 * h * focus_dist
	var viewport_width = viewport_height * (float64(image_width) / float64(image_height))

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	var w = lookfrom.minus(lookat).unit()
	var u = vup.cross(w).unit()
	var v = w.cross(u)

	// Calculate the vectors across the horizontal and down the verical viewport edges.
	var viewport_u = u.times(viewport_width)
	var viewport_v = v.inverse().times(viewport_height)

	// Calculate the horizontal and vertical delta vectors from pizel to pixel.
	var pixel_delta_u = viewport_u.divBy(float64(image_width))
	var pixel_delta_v = viewport_v.divBy(float64(image_height))

	// Calculate the location of the upper left pixel.
	var viewport_upper_left = center.minus(w.times(focus_dist)).minus(viewport_u.divBy(2.0)).minus(viewport_v.divBy(2.0))
	var pixel00_loc = viewport_upper_left.plus(pixel_delta_u.plus(pixel_delta_v).times(0.5))

	// Calculate the camera defocus disk basis vectors.
	var defocus_radius = focus_dist * math.Tan(Deg2Rad(defocus_angle/2.0))
	var defocus_disk_u = u.times(defocus_radius)
	var defocus_disk_v = v.times(defocus_radius)

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
		u, v, w,
		defocus_angle,
		defocus_disk_u,
		defocus_disk_v,
	}
}

func (cam *Camera) worker(world *Group, jobs <-chan Pixel, results chan<- Pixel, wg *sync.WaitGroup) {
	defer wg.Done()
	for pixel := range jobs {
		for range cam.samples_per_pixel {
			var r = cam.getRay(pixel.X, pixel.Y)
			pixel.Color.add(cam.rayColor(r, cam.max_depth, world))
		}
		pixel.Color.scale(cam.pixel_samples_scale)
		results <- pixel
		var progress = float64(len(results)) / float64(cam.image_width*cam.image_height)
		fmt.Fprint(os.Stderr, "\rScalines remaining: ", math.Round(progress*100.0), "% ")
	}
}

func (cam *Camera) render(world *Group) {
	fmt.Println("P3\n", cam.image_width, " ", cam.image_height, "\n255")
	var pixels = make([]Vec3, cam.image_width*cam.image_height)
	for i := range len(pixels) {
		pixels[i] = BLACK
	}

	var num_workers = runtime.NumCPU()
	num_workers = 1

	var jobs = make(chan Pixel, len(pixels))
	var results = make(chan Pixel, len(pixels))
	var wg sync.WaitGroup

	for w := 1; w <= num_workers; w++ {
		wg.Add(1)
		go cam.worker(world, jobs, results, &wg)
	}

	for j := range cam.image_height {
		for i := range cam.image_width {
			jobs <- Pixel{i, j, BLACK}
		}
	}
	close(jobs)

	wg.Wait()
	close(results)

	for pixel := range results {
		pixels[pixel.Y*cam.image_width+pixel.X] = pixel.Color
	}

	for _, pixel := range pixels {
		WriteColor(os.Stdout, pixel)
	}

	// fmt.Fprint(os.Stderr, "\rDone.                                 ")
}

func (cam *Camera) getRay(i int, j int) Ray {
	// Construct a camera ray originating from the origin and directed at randomly sampled
	// point around the pixel location i, j.

	var offset = SampleSquare()
	var pixel_sample = cam.pixel00_loc.
		plus(cam.pixel_delta_u.times(float64(i) + offset.X)).
		plus(cam.pixel_delta_v.times(float64(j) + offset.Y))

	var ray_origin = cam.center
	if cam.defocus_angle > 0.0 {
		ray_origin = cam.defocusDiskSample()
	}
	var ray_direction = pixel_sample.minus(ray_origin)

	return Ray{ray_origin, ray_direction}
}

func (cam *Camera) defocusDiskSample() Vec3 {
	var p = RandUnitDisk()
	return cam.center.plus(cam.defocus_disk_u.times(p.X)).plus(cam.defocus_disc_v.times(p.Y))
}

func (cam *Camera) rayColor(r Ray, depth int, world *Group) Vec3 {
	if depth <= 0 {
		return BLACK
	}

	var hit = world.hit(r, Interval{0.001, math.Inf(1)})
	if hit.DidHit {
		var scat = hit.Mat.scatter(r, hit)
		if scat.DidScatter {
			return scat.Attenuation.mul(cam.rayColor(scat.Scattered, depth-1, world))
		}
		return BLACK
	}
	var unit_direction = r.Dir.unit()
	var a = 0.5 * (unit_direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.times(1.0 - a).plus(Vec3{0.5, 0.7, 1.0}.times(a))
}
