package main

type Ray struct {
	Origin Vec3
	Dir    Vec3
}

func (r Ray) at(t float64) Vec3 {
	var v = r.Dir.times(t)
	return *v.add(r.Origin)
}
