package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	d := NewDeveloper("countryGarden")
	building := d.Build(30)
	t.Logf("building: %+v", building)
}
