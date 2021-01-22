package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var house House

	house = new(generalHouseBuilder).BuildHouse()
	t.Logf("house type: %+v", house.GetHouseType())

	house = new(plantHouseBuilder).BuildHouse()
	t.Logf("house type: %+v", house.GetHouseType())

	house = new(villaHouseBuilder).BuildHouse()
	t.Logf("house type: %+v", house.GetHouseType())
}
