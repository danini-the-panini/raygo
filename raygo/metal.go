package main

type Metal struct {
	Albedo Vec3
}

func (m *Metal) scatter(r Ray, hit Hit) Scatter {
	var reflected = r.Dir.reflect(hit.Normal)
	return DidScatter(m.Albedo, Ray{hit.P, reflected})
}
