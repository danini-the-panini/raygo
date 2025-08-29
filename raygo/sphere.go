package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
	Mat    Material
}

func (s *Sphere) hit(r Ray, ray_t Interval) Hit {
	var oc = s.Center.minus(r.Origin)
	var a = r.Dir.lenSq()
	var h = r.Dir.dot(oc)
	var c = oc.lenSq() - s.Radius*s.Radius

	var discriminant = h*h - a*c
	if discriminant < 0.0 {
		return NoHit
	}

	var sqrtd = math.Sqrt(discriminant)

	var root = (h - sqrtd) / a
	if !ray_t.surrounds(root) {
		root = (h + sqrtd) / a
		if !ray_t.surrounds(root) {
			return NoHit
		}
	}

	var p = r.at(root)

	return DidHit(p, root, r, p.minus(s.Center).divBy(s.Radius), s.Mat)
}
