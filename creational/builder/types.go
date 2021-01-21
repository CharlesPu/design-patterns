package builder

import (
	"fmt"
)

type (
	HouseBuilder interface {
		BuildFloors(n int)
		BuildGarge(v string)
		BuildGarden(s int)
		BuildTrees(n int)
	}

	House struct {
		floors      *Floors
		garge       *Garge
		garden      *Garden
		treesAround int
	}

	Floors struct {
		n        int
		material string
	}
	Garge struct {
		vehicle string
	}
	Garden struct {
		flower string
		area   int
	}
)

func (h *House) String() string {
	return fmt.Sprintf("House{floors: %+v, garge: %+v, garden: %+v, trees: %+v}", h.floors, h.garge, h.garden, h.treesAround)
}
