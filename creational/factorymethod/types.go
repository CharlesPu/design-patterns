package factorymethod

type (
	House interface {
		GetHouseType() string
	}

	HouseFactory interface {
		BuildHouse() House
	}
)
