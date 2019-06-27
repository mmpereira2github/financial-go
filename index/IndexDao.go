package index

var refIndexDao = inMemoryRefIndexDao{}

type RefIndexDao interface {
	GetById(int) (RefIndex, error)
	FindByName(string) RefIndex
	Save(RefIndex) error
}

func GetRefIndexDao() RefIndexDao { return refIndexDao }
