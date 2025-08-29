package main

type Group struct {
	Objects []Object
}

func NewGroup() Group {
	return Group{[]Object{}}
}

func (g *Group) clear() {
	g.Objects = []Object{}
}

func (g *Group) add(obj Object) {
	g.Objects = append(g.Objects, obj)
}

func (g *Group) hit(r Ray, ray_t Interval) Hit {
	var temp_hit = NoHit
	var hit Hit = NoHit
	var closest = ray_t.Max

	for _, obj := range g.Objects {
		temp_hit = obj.hit(r, Interval{ray_t.Min, closest})
		if temp_hit.DidHit {
			closest = temp_hit.T
			hit = temp_hit
		}
	}

	return hit
}
