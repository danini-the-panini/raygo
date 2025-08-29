package main

type Ray struct {
	Origin Vec3
	Dir    Vec3
}

func At(r Ray, t float64) Vec3 {
	var v = Times(r.Dir, t)
	return *Add(&v, r.Origin)
}
