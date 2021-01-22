package factorymethod

type (
	generalHouse struct {
	}
)

func (gh *generalHouse) GetHouseType() string {
	return "general house"
}

type (
	generalHouseBuilder struct {
	}
)

// override by children
func (gb *generalHouseBuilder) BuildHouse() House {
	return &generalHouse{}
}
