package domain

type Branch struct {
	Name string
}

func (p Branch) String() string {
	return p.Name
}
