package main

import (
	"math"
)

func RayColor(r Ray, world *Group) Vec3 {
	var hit = world.hit(r, Interval{0.0, math.Inf(1)})
	if hit.DidHit {
		return hit.Normal.plus(Vec3{1.0, 1.0, 1.0}).times(0.5)
	}
	var unit_direction = r.Dir.unit()
	var a = 0.5 * (unit_direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.times(1.0 - a).plus(Vec3{0.5, 0.7, 1.0}.times(a))
}

func main() {
	var world = NewGroup()

	var mat_ground = Lambertian{Vec3{0.8, 0.8, 0.0}}
	var mat_center = Lambertian{Vec3{0.1, 0.2, 0.5}}
	var mat_left = Metal{Vec3{0.8, 0.8, 0.8}, 0.3}
	var mat_right = Metal{Vec3{0.8, 0.6, 0.2}, 1.0}

	var ground = Sphere{Vec3{0.0, -100.5, -1.0}, 100.0, &mat_ground}
	var sphere_center = Sphere{Vec3{0.0, 0.0, -1.2}, 0.5, &mat_center}
	var sphere_left = Sphere{Vec3{-1.0, 0.0, -1.2}, 0.5, &mat_left}
	var sphere_right = Sphere{Vec3{1.0, 0.0, -1.2}, 0.5, &mat_right}

	world.add(&ground)
	world.add(&sphere_center)
	world.add(&sphere_left)
	world.add(&sphere_right)

	var cam = NewCamera(16.0/9.0, 400, 100, 50)

	cam.render(&world)
}
