package prototype

import (
	"fmt"
)

type (
	House interface {
		String() string
		Clone() House
	}

	Building struct {
		houses []House
	}
)

func (b Building) String() string {
	if len(b.houses) == 0 {
		return "nil Building"
	}
	return fmt.Sprintf("Building with %+v houses: %+v", len(b.houses), b.houses[0])
}
