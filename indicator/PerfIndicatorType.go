package indicator

import "fmt"

// PerfIndicatorCategory means a referenced index in the financial market used to compare performance
type PerfIndicatorCategory interface {
	Name() string
	String() string
}

type perfIndicatorCat struct {
	id   int
	name string
}

// NewPerfIndicatorType creates a new PerfIndicatorType instance given its id and name
func NewPerfIndicatorType(name string) PerfIndicatorCategory {
	return &perfIndicatorCat{-1, name}
}

func (r *perfIndicatorCat) Name() string { return r.name }
func (r *perfIndicatorCat) String() string {
	return fmt.Sprintf("PerfIndicatorType(id=%d, name='%s')", r.id, r.name)
}
