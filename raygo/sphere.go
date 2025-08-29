package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (self *Sphere) hit(r Ray, ray_tmin float64, ray_tmax float64) Hit {
	var oc = self.Center.minus(r.Origin)
	var a = r.Dir.lenSq()
	var h = r.Dir.dot(oc)
	var c = oc.lenSq() - self.Radius*self.Radius

	var discriminant = h*h - a*c
	if discriminant < 0.0 {
		return NoHit
	}

	var sqrtd = math.Sqrt(discriminant)

	var root = (h - sqrtd) / a
	if root < ray_tmin || ray_tmax <= root {
		root = (h + sqrtd) / a
		if root <= ray_tmin || ray_tmax <= root {
			return NoHit
		}
	}

	var p = r.at(root)

	return DidHit(p, root, r, p.minus(self.Center).divBy(self.Radius))
}
