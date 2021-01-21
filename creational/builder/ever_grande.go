package builder

type (
	// 恒大
	everGrande struct {
		house *House
	}
)

func NewEverGrande() *everGrande {
	return &everGrande{
		house: &House{},
	}
}

func (eg *everGrande) BuildFloors(n int) {
	eg.house.floors = &Floors{
		n:        n,
		material: "pottery",
	}
}

func (eg *everGrande) BuildGarge(v string) {
	eg.house.garge = &Garge{
		vehicle: v,
	}
}

func (eg *everGrande) BuildGarden(s int) {
	eg.house.garden = &Garden{
		flower: "tulip",
		area:   s,
	}
}

func (eg *everGrande) BuildTrees(n int) {
	eg.house.treesAround = n
}

func (eg *everGrande) GetHouse() *House {
	defer func() {
		eg.house = &House{}
	}()
	return eg.house
}
