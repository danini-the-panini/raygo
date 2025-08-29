package main

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m *Metal) scatter(r Ray, hit Hit) Scatter {
	var reflected = r.Dir.reflect(hit.Normal)
	reflected.normalize().add(RandUnit().times(m.Fuzz))
	var scattered = Ray{hit.P, reflected}
	if scattered.Dir.dot(hit.Normal) <= 0 {
		return NoScatter
	}
	return DidScatter(m.Albedo, scattered)
}
