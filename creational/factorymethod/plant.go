package factorymethod

type (
	plant struct {
	}
)

func (p *plant) GetHouseType() string {
	return "plant"
}

type (
	plantHouseBuilder struct {
		generalHouseBuilder
	}
)

func (pb *plantHouseBuilder) BuildHouse() House {
	return &plant{}
}
