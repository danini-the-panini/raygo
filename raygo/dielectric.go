package main

import (
	"math"
	"math/rand"
)

func reflectance(cosine float64, refraction_index float64) float64 {
	// Use Schlick's approximation for reflectance.
	var r0 = (1.0 - refraction_index) / (1.0 + refraction_index)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow((1.0-cosine), 5.0)
}

type Dielectric struct {
	RefractionIndex float64
}

func (m *Dielectric) scatter(r Ray, hit Hit) Scatter {
	var ri = m.RefractionIndex
	if hit.FrontFace {
		ri = 1.0 / ri
	}

	var unit_direction = r.Dir.unit()
	var cos_theta = math.Min(unit_direction.inverse().dot(hit.Normal), 1.0)
	var sin_theta = math.Sqrt(1.0 - cos_theta*cos_theta)

	var cannot_refract = ri*sin_theta > 1.0
	var direction Vec3

	if cannot_refract || reflectance(cos_theta, ri) > rand.Float64() {
		direction = unit_direction.reflect(hit.Normal)
	} else {
		direction = unit_direction.refract(hit.Normal, ri)
	}

	return DidScatter(WHITE, Ray{hit.P, direction})
}
