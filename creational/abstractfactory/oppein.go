package abstractfactory

type (
	oppeinCompany struct {
	}
)

func (oc *oppeinCompany) BuildFloor() Floor {
	return Floor{
		material: "pottery",
	}
}

func (oc *oppeinCompany) BuildDoor() Door {
	return Door{
		kind: "glass door",
	}
}

func (oc *oppeinCompany) BuildWindow() Window {
	return Window{
		shape: "circle",
	}
}
