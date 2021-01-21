package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	d := &director{}

	developerA := NewCountryGarden()
	developerB := NewEverGrande()

	d.BuildVilla(developerA)
	villa := developerA.GetHouse()
	t.Logf("villa built by countryGarden: %+v", villa)

	d.BuildFlat(developerB)
	flat := developerB.GetHouse()
	t.Logf("flat built by everGrande: %+v", flat)
}
