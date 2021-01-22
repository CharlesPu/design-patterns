package prototype

type (
	developer struct {
		name string
	}
)

func NewDeveloper(n string) *developer {
	return &developer{
		name: n,
	}
}

func (d *developer) Build(n int) *Building {
	h := &commercialHouse{
		area:    120,
		roomNum: 4,
	}

	b := &Building{}
	b.houses = append(b.houses, h)

	for i := 1; i < n; i++ {
		b.houses = append(b.houses, h.Clone())
	}

	return b
}
