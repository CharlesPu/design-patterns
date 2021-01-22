package factorymethod

type (
	villa struct {
	}
)

func (v *villa) GetHouseType() string {
	return "villa"
}

type (
	villaHouseBuilder struct {
		generalHouseBuilder
	}
)

func (vb *villaHouseBuilder) BuildHouse() House {
	return &villa{}
}
