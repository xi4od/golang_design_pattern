package flyweight

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

type FlyWeightFactory struct {
	Pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{Pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := f.Pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.Pool[name] = weight
	}
	return weight
}
