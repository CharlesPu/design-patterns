package prototype

import (
	"fmt"
)

type (
	commercialHouse struct {
		area    int
		roomNum int
	}
)

func (cb *commercialHouse) Clone() House {
	r := *cb
	return &r
}

func (cb *commercialHouse) String() string {
	return fmt.Sprintf("House{area: %+v, roomNum: %+v}", cb.area, cb.roomNum)
}
