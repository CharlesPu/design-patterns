package builder

type (
	director struct {
	}
)

func (d *director) BuildVilla(b HouseBuilder) {
	b.BuildFloors(3)
	b.BuildGarge("car")
	b.BuildGarden(100)
	b.BuildTrees(3)
}

func (d *director) BuildFlat(b HouseBuilder) {
	b.BuildFloors(30)
	b.BuildGarge("bike")
}
