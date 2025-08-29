package main

type Hit struct {
	P         Vec3
	Normal    Vec3
	Mat       Material
	T         float64
	FrontFace bool
	DidHit    bool
}

var NoHit = Hit{ZERO3, ZERO3, &NULL_MAT, 0.0, true, false}

func DidHit(p Vec3, t float64, r Ray, outward_normal Vec3, mat Material) Hit {
	var front_face = r.Dir.dot(outward_normal) < 0.0
	var normal = outward_normal
	if !front_face {
		normal.negate()
	}
	return Hit{p, normal, mat, t, front_face, true}
}

type Object interface {
	hit(r Ray, ray_t Interval) Hit
}
