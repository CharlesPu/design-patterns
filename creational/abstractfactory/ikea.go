package abstractfactory

type (
	ikeaCompany struct {
	}
)

func (ic *ikeaCompany) BuildFloor() Floor {
	return Floor{
		material: "wooden",
	}
}

func (ic *ikeaCompany) BuildDoor() Door {
	return Door{
		kind: "security door",
	}
}

func (ic *ikeaCompany) BuildWindow() Window {
	return Window{
		shape: "rectangle",
	}
}
