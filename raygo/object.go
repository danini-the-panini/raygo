package main

type Hit struct {
	P         Vec3
	Normal    Vec3
	T         float64
	FrontFace bool
	DidHit    bool
}

var NoHit = Hit{Vec3{0.0, 0.0, 0.0}, Vec3{0.0, 0.0, 0.0}, 0.0, true, false}

func DidHit(p Vec3, t float64, r Ray, outward_normal Vec3) Hit {
	var front_face = r.Dir.dot(outward_normal) < 0.0
	var normal = outward_normal
	if !front_face {
		normal.negate()
	}
	return Hit{p, normal, t, front_face, true}
}

type Object interface {
	hit(r Ray, ray_tmin float64, ray_tmax float64) Hit
}
