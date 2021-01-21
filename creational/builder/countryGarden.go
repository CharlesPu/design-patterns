package builder

type (
	// 碧桂园
	countryGarden struct {
		house *House
	}
)

func NewCountryGarden() *countryGarden {
	return &countryGarden{
		house: &House{},
	}
}

func (cg *countryGarden) BuildFloors(n int) {
	cg.house.floors = &Floors{
		n:        n,
		material: "wooden",
	}
}

func (cg *countryGarden) BuildGarge(v string) {
	cg.house.garge = &Garge{
		vehicle: v,
	}
}

func (cg *countryGarden) BuildGarden(s int) {
	cg.house.garden = &Garden{
		flower: "orchid",
		area:   s,
	}
}

func (cg *countryGarden) BuildTrees(n int) {
	cg.house.treesAround = n
}

func (cg *countryGarden) GetHouse() *House {
	defer func() {
		cg.house = &House{}
	}()
	return cg.house
}
