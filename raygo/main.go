package main

import (
	"math"
	"math/rand"
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

	// var mat_ground = Lambertian{Vec3{0.8, 0.8, 0.0}}
	// var mat_center = Lambertian{Vec3{0.1, 0.2, 0.5}}
	// var mat_left = Dielectric{1.50}
	// var mat_bubble = Dielectric{1.00 / 1.50}
	// var mat_right = Metal{Vec3{0.8, 0.6, 0.2}, 1.0}

	// var ground = Sphere{Vec3{0.0, -100.5, -1.0}, 100.0, &mat_ground}
	// var sphere_center = Sphere{Vec3{0.0, 0.0, -1.2}, 0.5, &mat_center}
	// var sphere_left = Sphere{Vec3{-1.0, 0.0, -1.0}, 0.5, &mat_left}
	// var sphere_bubble = Sphere{Vec3{-1.0, 0.0, -1.0}, 0.4, &mat_bubble}
	// var sphere_right = Sphere{Vec3{1.0, 0.0, -1.0}, 0.5, &mat_right}

	// world.add(&ground)
	// world.add(&sphere_center)
	// world.add(&sphere_left)
	// world.add(&sphere_bubble)
	// world.add(&sphere_right)

	var mat_ground = Lambertian{Vec3{0.5, 0.5, 0.5}}
	var ground = Sphere{Vec3{0, -1000.0, 0.0}, 1000.0, &mat_ground}
	world.add(&ground)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			var choose_mat = rand.Float64()
			var center = Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.minus(Vec3{4.0, 0.2, 0.0}).len() > 0.9 {
				var sphere_mat Material

				if choose_mat < 0.8 {
					// diffuse
					var albedo = RandVec3().mul(RandVec3())
					sphere_mat = &Lambertian{albedo}
					var sphere = Sphere{center, 0.2, sphere_mat}
					world.add(&sphere)
				} else if choose_mat < 0.95 {
					// metal
					var albedo = Interval{0.5, 1.0}.randVec3()
					var fuzz = Interval{0.0, 0.5}.rand()
					sphere_mat = &Metal{albedo, fuzz}
					var sphere = Sphere{center, 0.2, sphere_mat}
					world.add(&sphere)
				} else {
					// glass
					sphere_mat = &Dielectric{1.5}
					var sphere = Sphere{center, 0.2, sphere_mat}
					world.add(&sphere)
				}
			}
		}
	}

	var mat1 = Dielectric{1.5}
	var sphere1 = Sphere{Vec3{0.0, 1.0, 0.0}, 1.0, &mat1}
	world.add(&sphere1)

	var mat2 = Lambertian{Vec3{0.4, 0.2, 0.1}}
	var sphere2 = Sphere{Vec3{-4.0, 1.0, 0.0}, 1.0, &mat2}
	world.add(&sphere2)

	var mat3 = Metal{Vec3{0.7, 0.6, 0.5}, 0.0}
	var sphere3 = Sphere{Vec3{4.0, 1.0, 0.0}, 1.0, &mat3}
	world.add(&sphere3)

	var cam = NewCamera(
		16.0/9.0,
		1200,
		20.0,
		Vec3{13.0, 2.0, 3.0},
		Vec3{0.0, 0.0, -1.0},
		Y_UP,
		0.6,
		10.0,
		10,
		50,
	)

	cam.render(&world)
}
