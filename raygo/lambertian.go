package main

type Lambertian struct {
	Albedo Vec3
}

func (m *Lambertian) scatter(r Ray, hit Hit) Scatter {
	var scatter_direction = hit.Normal.plus(RandUnit())
	if scatter_direction.nearZero() {
		scatter_direction = hit.Normal
	}
	return DidScatter(m.Albedo, Ray{hit.P, scatter_direction})
}
