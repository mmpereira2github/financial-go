package index

import "fmt"

type RefIndex interface {
	Id() int
	Name() string
	String() string
}

type refIndex struct {
	id   int
	name string
}

func NewRefIndex(id int, name string) RefIndex { return refIndex{id, name} }
func (r refIndex) Id() int                     { return r.id }
func (r refIndex) Name() string                { return r.name }
func (r refIndex) String() string              { return fmt.Sprintf("RefIndex(id=%d, name='%s')", r.id, r.name) }
