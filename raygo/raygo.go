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

	var sphere1 = Sphere{Vec3{0.0, 0.0, -1.0}, 0.5}
	var sphere2 = Sphere{Vec3{0.0, -100.5, -1.0}, 100.0}

	world.add(&sphere1)
	world.add(&sphere2)

	var cam = NewCamera(16.0/9.0, 400)

	cam.render(&world)
}
