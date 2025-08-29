package main

type Scatter struct {
	Attenuation Vec3
	Scattered   Ray
	DidScatter  bool
}

var NoScatter = Scatter{BLACK, Ray{ZERO3, ZERO3}, false}

func DidScatter(attenuation Vec3, scattered Ray) Scatter {
	return Scatter{attenuation, scattered, true}
}

type Material interface {
	scatter(r Ray, hit Hit) Scatter
}

type NullMat struct{}

func (_ *NullMat) scatter(r Ray, hit Hit) Scatter {
	return NoScatter
}

var NULL_MAT = NullMat{}
