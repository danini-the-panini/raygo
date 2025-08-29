package main

type Group struct {
	Objects []Object
}

func NewGroup() Group {
	return Group{[]Object{}}
}

func (self *Group) clear() {
	self.Objects = []Object{}
}

func (self *Group) add(obj Object) {
	self.Objects = append(self.Objects, obj)
}

func (self *Group) hit(r Ray, ray_t Interval) Hit {
	var temp_hit = NoHit
	var hit Hit = NoHit
	var closest = ray_t.Max

	for _, obj := range self.Objects {
		temp_hit = obj.hit(r, Interval{ray_t.Min, closest})
		if temp_hit.DidHit {
			closest = temp_hit.T
			hit = temp_hit
		}
	}

	return hit
}
