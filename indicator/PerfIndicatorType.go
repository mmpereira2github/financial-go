package indicator

import "fmt"

// PerfIndicatorType means a referenced index in the financial market used to compare performance
type PerfIndicatorType interface {
	ID() int
	Name() string
	String() string
}

type perfIndicatorType struct {
	id   int
	name string
}

// NewPerfIndicatorType creates a new PerfIndicatorType instance given its id and name
func NewPerfIndicatorType(id int, name string) PerfIndicatorType {
	return &perfIndicatorType{id, name}
}
func (r *perfIndicatorType) ID() int      { return r.id }
func (r *perfIndicatorType) Name() string { return r.name }
func (r *perfIndicatorType) String() string {
	return fmt.Sprintf("PerfIndicatorType(id=%d, name='%s')", r.id, r.name)
}
