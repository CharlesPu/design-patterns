package abstractfactory

import (
	"fmt"
)

type (
	FurnitureAbstractFactory interface {
		BuildFloor() Floor
		BuildDoor() Door
		BuildWindow() Window
	}

	Floor struct {
		material string
	}
	Door struct {
		kind string
	}
	Window struct {
		shape string
	}

	House struct {
		floor  Floor
		door   Door
		window Window
	}
)

func (h *House) String() string {
	return fmt.Sprintf("House{fllor: %+v, door: %+v, window: %+v}", h.floor.material, h.door.kind, h.window.shape)
}
