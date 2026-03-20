package structuresandmethods

type Package struct {
	ID        string
	Delivered bool
}

func (p Package) Status() string {

	if p.Delivered == true {
		return "Package " + p.ID + " has been delivered"
	} else {
		return "Package " + p.ID + " is still in transit"
	}
}
