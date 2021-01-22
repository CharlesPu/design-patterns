package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var coX, coY FurnitureAbstractFactory
	coX = &ikeaCompany{}
	coY = &oppeinCompany{}

	houseA := &House{
		floor:  coX.BuildFloor(),
		door:   coY.BuildDoor(),
		window: coX.BuildWindow(),
	}
	t.Log(houseA.String())

	houseB := &House{
		floor:  coY.BuildFloor(),
		door:   coX.BuildDoor(),
		window: coY.BuildWindow(),
	}
	t.Log(houseB.String())
}
